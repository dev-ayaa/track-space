package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-chi/chi"

	"github.com/yusuf/track-space/pkg/handler"
	// "github.com/yusuf/track-space/pkg/config"
	// "github.com/yusuf/track-space/pkg/handler"

)

func Routes(routes *gin.Engine) {
	routes.GET("/", handler.AppRepo.HomePage())
	routes.GET("/login", handler.AppRepo.LoginPage())

}
