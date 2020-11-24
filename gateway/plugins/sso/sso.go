package sso

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

const API_URL string = "https://api-sso-ensaas.sa.wise-paas.com/v4.0"

type Token struct {
	tok string
}

var token Token

// TODO: 408 error
// Reason: unknown
// Solution: restart the target service
func NewPlugin() plugin.Plugin {
	return plugin.NewPlugin(
		plugin.WithName("sso"),
		plugin.WithFlag(
			&cli.StringFlag{
				Name:  "username",
				Usage: "user name",
				Value: os.Getenv("SSO_USER"),
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "user password",
				Value: os.Getenv("SSO_PASS"),
			}),

		// Get token by username/password
		plugin.WithInit(func(ctx *cli.Context) error {
			username := ctx.String("username")
			password := ctx.String("password")

			if username == "" {
				log.Fatal("os-env: SSO_USER is not specify.")
				return nil
			}

			if password == "" {
				log.Fatal("os-env: SSO_PASS is not specify.")
				return nil
			}

			reqBody, err := json.Marshal(map[string]string{
				"username": username,
				"password": password,
			})

			resp, err := http.Post(API_URL+"/auth/native", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode == 200 {
				defer resp.Body.Close()
				var result map[string]interface{}

				json.NewDecoder(resp.Body).Decode(&result)

				fmt.Println("getToken result: ")
				fmt.Println(result)

				token.tok = result["accessToken"].(string)
			}

			return nil
		}),

		// validate sso token
		plugin.WithHandler(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				reqToken := r.Header.Get("Authorization") // Pattern:
				splitToken := strings.Split(reqToken, "Bearer")
				if len(splitToken) != 2 {
					log.Fatal("Bearer token not in proper format.")
				}
				reqToken = strings.TrimSpace(splitToken[1])
				reqBody, err := json.Marshal(map[string]string{
					"token": reqToken,
					// "token": token.tok,
				})

				resp, err := http.Post(API_URL+"/tokenvalidation", "application/json", bytes.NewBuffer(reqBody))
				if err != nil {
					fmt.Println(err)
				}

				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Validate response body: " + string(body))

				if err != nil || resp.StatusCode != 200 {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				h.ServeHTTP(w, r)
			})
		}),
	)
}
