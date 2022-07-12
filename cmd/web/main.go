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
		return
	}

	//Setting up the database connection for mongoDB
	//  = data.DatabaseConnection()
	// if err != nil{
	// 	log.Println("error connecting to database")
	// 	log.Panic(err)
	// }
	
	repo := handler.NewRepository(&app,user, mail)
	handler.NewHandler(repo)



	portNumber := os.Getenv("PORTNUMBER")
	if portNumber == ""{
		log.Println("No local server port number created!")
	}

	//
	app_router := gin.New()

	app_router.Use(gin.Logger(), gin.Recovery())

	//this is to verify a list of the html/tmpl
	//and render the correct template with it templates data
	app_router.SetFuncMap(template.FuncMap{})

	app_router.LoadHTMLGlob("templates/*")

	Routes(app_router)

	err = app_router.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}

}
