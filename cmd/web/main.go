package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/data"

)

var app config.AppConfig

func main() {

	app.AppInProduction = false
	app.UseTempCache = false

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file available")
		return
	}

	mongodb_uri := os.Getenv("MONGODB_URI")
	if mongodb_uri == "" {
		log.Println("mongodb cluster uri not found : ")
	}


	//Setting up the database connection for mongoDB
	err = data.DatabaseConnection(mongodb_uri)
	if err != nil{
		log.Println("error connecting to database")
		log.Panic(err)
	}


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
