package main

import (
	"crypto/rsa"
	"log"
	"micro-playground/gateway/plugins/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/test"
)

// 加密token的私鑰
var priKey *rsa.PrivateKey

// 生成用戶ID為123的token
func main() {
	priKey = test.LoadRSAPrivateKeyFromDisk("./conf/private.key")
	token, err := GenerateToken("123")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("token: ", token)
	}
}

// 根據用戶ID產生token
func GenerateToken(userId string) (string, error) {
	// 設置token有效時間
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := auth.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 過期時間
			ExpiresAt: expireTime.Unix(),
			// 指定token發行人
			Issuer: "micro-auth",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// 內部生成簽名字符串，再用於獲取完整、已簽名的token
	token, err := tokenClaims.SignedString(priKey)
	return token, err
}
func ParseToken(token string) (*auth.Claims, error) {

	// 用於解析鑑權的聲明，方法內部主要是具體的解碼和校驗的過程，最終返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return priKey.Public(), nil
	})

	if tokenClaims != nil {
		// 從tokenClaims中獲取到Claims對象，並使用斷言，將該對象轉換為我們自己定義的Claims
		// 要傳入指針，項目中結構體都是用指針傳遞，節省空間。
		if claims, ok := tokenClaims.Claims.(*auth.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
