package controller

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/key"
	"github.com/yusuf/track-space/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var Validate = validator.New()

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
		c.HTML(http.StatusOK, "signup.page.tmpl", gin.H{})
	}
}

func (ah AppHandler) PostSignUpPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := Validate.Struct(&user); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			return
		}

		//Parse the form and check for an err
		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user.ID = primitive.NewObjectID()
		user.Email = c.Request.Form.Get("email")
		user.Password = key.HashPassword(c.Request.Form.Get("password"))

		scs := sessions.Default(c)
		scs.Set("email", user.Email)
		scs.Set("password", user.Password)
		err = scs.Save()
		if err != nil {
			log.Println(err)
		}

		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCtx()

		signUpDoc := bson.D{
			{"email", user.Email},
			{"password", user.Password},
		}
		_, err = ah.UserCol.InsertOne(ctx, signUpDoc)
		if err != nil {
			log.Println("cannot insert user sign up details in the database")
			_ = c.AbortWithError(http.StatusInternalServerError, gin.Error{Err: err})
			return
		}

		c.Redirect(http.StatusSeeOther, "/user-info")
	}
}

func (ah *AppHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		scs := sessions.Default(c)

		c.JSON(http.StatusOK, gin.H{"email": scs.Get("email"), "password": scs.Get("password")})
		//c.HTML(http.StatusOK, "user-info.page.tmpl", gin.H{})
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
