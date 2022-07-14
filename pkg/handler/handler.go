package handler

import (
	// "log"?
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yusuf/track-space/pkg/config"
)

//Implement the repository pattern to access multiple package all at once

//Repository : this will give me access to the app configuration package
//and the database collections as well
type Repository struct {
	AppConfig *config.AppConfig
	UserCol   *mongo.Collection
	MailCol   *mongo.Collection
}

var AppRepo *Repository

func NewRepository(appConfig *config.AppConfig, userCol *mongo.Collection, mailCol *mongo.Collection) *Repository {
	return &Repository{
		AppConfig: appConfig,
		UserCol:   userCol,
		MailCol:   mailCol,
	}
}
func NewHandler(r *Repository) {
	r = AppRepo
}

//HomePage handler to display the home page of the application
func (rp *Repository) HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// rp.User_col.CountDocuments()
		// ctx.Header()
		b := ctx.Request.Header.Get("X-FORWARDED-FOR")

		a := ctx.Request.RemoteAddr
		// a:=ctx.ClientIP()
		ctx.HTML(http.StatusOK, "home.page.tmpl", gin.H{"ip": a, "s": b})
	}
}

//LoginPage handler for the user to sign up for an account  and login as well
func (rp *Repository) LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}
