package auth

import (
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

// 認證相關參數

// Claims是一些實體（通常指的用戶）的狀態和額外的數據
type Claims struct {
	// 在jwt默認Claims基礎上增加用戶ID信息
	UserId string `json:"userId"`
	jwt.StandardClaims
}

// 自己封裝的Plugin工廠方法，可參考官方插件增加一些options參數來增加插件的靈活配置
func NewPlugin() plugin.Plugin {
	var pubKey *rsa.PublicKey
	return plugin.NewPlugin(
		// 插件名
		plugin.WithName("auth"),
		// token解碼需要用到公鑰，這裡順百年演示了如何配置命令行參數
		plugin.WithFlag(
			&cli.StringFlag{
				Name:  "auth_key",
				Usage: "auth key file",
				Value: "./conf/public.key",
			}),
		// 配置插件初始化操作，cli.Context中包含了項目啟動參數
		plugin.WithInit(func(ctx *cli.Context) error {
			pubKeyFile := ctx.String("auth_key")
			pubKey = test.LoadRSAPublicKeyFromDisk(pubKeyFile)
			return nil
		}),
		// 配置處理函數，注意: 與wrapper不同，他的參數是http包的ResponseWriter和Request
		plugin.WithHandler(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var claims Claims
				token, err := request.ParseFromRequest(
					r,
					request.AuthorizationHeaderExtractor,
					func(*jwt.Token) (interface{}, error) {
						return pubKey, nil
					},
					request.WithClaims(&claims),
				)

				if err != nil {
					log.Print("token invalid: ", err.Error())
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				// token.Valid是否成功，取決於jwt中Claims interface定義的Valid() error方法
				// TODO: 這裏直接使用了默認Claims實現jwt.StandardClaims提供的方法，實際可能需要重寫
				if token == nil || !token.Valid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				// TODO: 雖然是有效的token，但並不意味著此用戶有權訪問所有接口

				// 從Claims種解析userID並加入Header
				r.Header.Set("userId", claims.UserId)

				// 通過了上述驗證後，必須執行下面這一步，保證其他插件和業務程式碼的執行
				h.ServeHTTP(w, r)
			})
		}),
	)
}
