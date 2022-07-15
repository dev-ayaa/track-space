package main

import (
	"context"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/data"
	"github.com/yusuf/track-space/pkg/dbdriver"
	"github.com/yusuf/track-space/pkg/handler"
)

var app config.AppConfig

func main() {

	app.AppInProduction = false
	app.UseTempCache = false

	var Client = dbdriver.DatabaseConnection()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file available")
	}

	user := data.UserData(Client, "user")
	//mail = data.MailData(Client, "mail")

	defer func() {
		if err = Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
			return
		}
	}()

	repo := handler.NewAppHandler(&app, user)

	portNumber := os.Getenv("PORTNUMBER")
	if portNumber == "" {
		log.Println("No local server port number created!")
	}

	//
	appRouter := gin.New()
	err = appRouter.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Println(err)
		log.Println("cannot access untrusted server proxy")

	}

	appRouter.Use(gin.Logger(), gin.Recovery())

	//this is to verify a list of the html/tmpl
	//and render the correct template with it templates data
	appRouter.SetFuncMap(template.FuncMap{})

	appRouter.LoadHTMLGlob("templates/*")

	Routes(appRouter, *repo)

	err = appRouter.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
}
