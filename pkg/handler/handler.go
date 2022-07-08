package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

// "fmt"
// "net/http"

// // fmt.Println("Welcome to Track Space
// wr.Write([]byte(fmt.Sprintf("Hello World, Welcome to Track Space")))
func HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.page.tmpl",gin.H{
			
		})
	}
}


func LoginPage() gin.HandlerFunc{
	return func(ctx *gin.Context) {}

}