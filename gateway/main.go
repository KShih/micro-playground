package main

import (
	"log"
	// "micro-playground/gateway/plugins/auth"
	"micro-playground/gateway/plugins/sso"

	"github.com/micro/micro/v2/client/api"

	"github.com/micro/micro/v2/cmd"
)

func main() {
	err := api.Register(sso.NewPlugin())
	if err != nil {
		log.Fatal("auth register")
	}
	cmd.Init()
}
