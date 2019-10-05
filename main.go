package main

import (
	"github.com/patrykjadamczyk/go-echo-log-server/config"
	"github.com/patrykjadamczyk/go-echo-log-server/logdb"
	"github.com/patrykjadamczyk/go-echo-log-server/logserver"
)

func main() {
	AppConfiguration := config.InitConfig()
	AppDatabase := logdb.Create()
	go logserver.Main(AppConfiguration, *AppDatabase)
	select {}
}
