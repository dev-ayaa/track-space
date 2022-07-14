package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/data"
	"github.com/yusuf/track-space/pkg/dbdriver"
	"github.com/yusuf/track-space/pkg/handler"
	// "github.com/yusuf/track-space/pkg/ipaddress"
)

var app config.AppConfig
var user *mongo.Collection = data.UserData(dbdriver.Client, "user")
var mail *mongo.Collection = data.MailData(dbdriver.Client, "mail")

func main() {

	app.AppInProduction = false
	app.UseTempCache = false

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file available")
	}
	// fmt.Println(ipaddress.UserIpAddress())
	repo := handler.NewRepository(&app, user, mail)
	handler.NewHandler(repo)

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

	Routes(appRouter)

	err = appRouter.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}
}
