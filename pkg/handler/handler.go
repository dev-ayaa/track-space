package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomePage handler to display the home page of the application
func HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.page.tmpl", gin.H{})
	}
}

//LoginPage handler for the user to sign up for an account  and login as well
func LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}
