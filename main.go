package main

import (
	"fmt"
	"github.com/ZW-Rays/golang-login/config"
	"github.com/ZW-Rays/golang-login/handler"
	"log"
	"net/http"
)

func main() {
	//initiate Config
	config, err := config.InitiateConfig()
	if err != nil {
		_ = fmt.Errorf("ERROR IN CONFIG : %s ", err.Error())
	}

	http.HandleFunc("/login", handler.MakeHandler(handler.LoginHandler))

	log.Println(fmt.Sprintf("Server started at %v:%v", config.Server.Host, config.Server.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), nil))
}
