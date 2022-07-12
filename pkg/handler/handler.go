package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yusuf/track-space/pkg/config"
)

//Implement the repository pattern to access multiple package all at once


//Repository : this will give me access to the app configuration package
//and the database collections as well
type Repository struct {
	App_config *config.AppConfig
	User_col   *mongo.Collection
	Mail_col   *mongo.Collection
}

var AppRepo *Repository

func NewRepository(app_config *config.AppConfig, user_col *mongo.Collection, mail_col *mongo.Collection) *Repository {
	return &Repository{
		App_config: app_config,
		User_col:   user_col,
		Mail_col:   mail_col,
	}
}
func NewHandler(r *Repository) {
	r = AppRepo
}

//HomePage handler to display the home page of the application
func (rp *Repository) HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// rp.User_col.CountDocuments()
		ctx.HTML(http.StatusOK, "home.page.tmpl", gin.H{})
	}
}

//LoginPage handler for the user to sign up for an account  and login as well
func (rp *Repository) LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}
