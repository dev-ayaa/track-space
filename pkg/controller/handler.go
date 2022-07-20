package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusuf/track-space/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

// AppHandler Implement the repository pattern to access multiple package all at once
// this will give me access to the app configuration package
// and the database collections as well
type AppHandler struct {
	AppConfig *config.AppConfig
	UserCol   *mongo.Collection
	// MailCol   *mongo.Collection
}

func NewAppHandler(appConfig *config.AppConfig, userCol *mongo.Collection) *AppHandler {
	return &AppHandler{
		AppConfig: appConfig,
		UserCol:   userCol,
		// MailCol:   mailCol,
	}
}

// HomePage controller to display the home page of the application
func (ah *AppHandler) HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{})
	}
}

func (ah AppHandler) SignUpPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}

func (ah AppHandler) PostSignUpPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_ = c.Request.Form.Get("email")
		_ = c.Request.Form.Get("password")
	}
}

func (ah *AppHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "user-info.page.tmpl", gin.H{})
	}
}

func (ah *AppHandler) PostUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// GetLoginPage LoginPage controller for the user to sign up for an account  and login as well
func (ah *AppHandler) GetLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//still have to design the login page
		c.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}

func (ah *AppHandler) PostLoginPage() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
