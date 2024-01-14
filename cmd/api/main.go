package main

import (
	"github.com/fairusatoir/golang-gin-rest/cmd/api/config"
	"github.com/fairusatoir/golang-gin-rest/internal/app"
	"github.com/fairusatoir/golang-gin-rest/internal/clients"
	"github.com/fairusatoir/golang-gin-rest/pkg/log"
	"github.com/sirupsen/logrus"
)

func init() {
	config.InitConfig()
	clients.InitDatamaster()
}

func main() {
	app, err := app.NewServerAPI()
	if err != nil {
		log.Panic(err.Error(), logrus.Fields{log.Category: log.Server})
	}
	if err := app.Run(); err != nil {
		log.Fatal(err.Error(), logrus.Fields{log.Category: log.Server})
	}
}
