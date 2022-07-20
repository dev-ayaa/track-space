package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-chi/chi"

	"github.com/yusuf/track-space/pkg/controller"
	// "github.com/yusuf/track-space/pkg/config"
	// "github.com/yusuf/track-space/pkg/controller"
)

func Routes(routes *gin.Engine, h controller.AppHandler) {
	routes.GET("/", h.HomePage())

	routes.GET("/sign-up", h.SignUpPage())
	routes.POST("/sign-up", h.PostSignUpPage())

	routes.GET("/user-info", h.GetUserInfo())
	routes.POST("/user-info", h.PostUserInfo())

	routes.GET("/login", h.GetLoginPage())
	routes.POST("/login", h.PostLoginPage())

}
