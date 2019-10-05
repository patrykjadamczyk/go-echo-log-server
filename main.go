package main

import (
	"github.com/patrykjadamczyk/go-echo-log-server/config"
	"github.com/patrykjadamczyk/go-echo-log-server/logserver"
)

func main() {
	AppConfiguration := config.InitConfig()
	go logserver.Main(AppConfiguration)
	select {}
}
