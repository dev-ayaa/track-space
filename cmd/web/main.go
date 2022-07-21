package main

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/controller"
	"github.com/yusuf/track-space/pkg/data"
	"github.com/yusuf/track-space/pkg/db"
	"html/template"
	"log"
	"os"
)

var app config.AppConfig

func main() {

	/*Application configuration*/
	app.AppInProduction = false
	app.UseTempCache = false

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file available")
	}

	/* Get the URI string from the ENV*/
	mongodbUri := os.Getenv("MONGODBURI")
	if mongodbUri == "" {
		log.Println("mongodb cluster uri not found : ")
	}

	/*Connecting to the database*/
	var Client = db.DatabaseConnection(mongodbUri)

	user := data.UserData(Client, "user")
	//mail = data.MailData(Client, "mail")

	/*Disconnect the database using defer -> FIRST-IN-LAST-OUT*/
	defer func() {
		if err = Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
			return
		}
	}()

	repo := controller.NewAppHandler(&app, user)

	/* Get the port address string from the ENV file */
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

	keyWare := cookie.NewStore([]byte("appKey"))
	appRouter.Use(sessions.Sessions("session", keyWare))

	//this is to verify a list of the html/tmpl
	//and render the correct template with it templates data
	appRouter.SetFuncMap(template.FuncMap{})

	appRouter.Static("/static", "./static")

	appRouter.LoadHTMLGlob("templates/*.tmpl")

	Routes(appRouter, *repo)

	err = appRouter.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
}
