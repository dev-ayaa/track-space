package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/yusuf/track-space/pkg/config"

)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	app.AppInProduction = false
	app.UseTempCache = false

	app_router := gin.New()

	app_router.Use(gin.Logger(), gin.Recovery())
	//this is to verify a list of the html/tmpl
	//and render the correct template with it templates data
	app_router.SetFuncMap(template.FuncMap{})

	app_router.LoadHTMLGlob("templates/*.html")

	Routes(app_router)

	err := app_router.Run(portNumber)
	if err != nil {
		log.Fatal(err)
	}

}
