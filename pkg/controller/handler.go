package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yusuf/track-space/pkg/config"
	"github.com/yusuf/track-space/pkg/key"
	"github.com/yusuf/track-space/pkg/model"
)

var Validate = validator.New()

// AppHandler Implement the repository pattern to access multiple package all at once
// this will give me access to the app configuration package
// and the database collections as well
type AppHandler struct {
	AppConfig *config.AppConfig
	UserCol   *mongo.Collection
}

func NewAppHandler(appConfig *config.AppConfig, userCol *mongo.Collection) *AppHandler {
	return &AppHandler{
		AppConfig: appConfig,
		UserCol:   userCol,
	}
}

// HomePage controller to display the home page of the application
func (ah *AppHandler) HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home-page.html", gin.H{})
	}
}

func (ah AppHandler) SignUpPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup-page.html", gin.H{})
	}
}

func (ah AppHandler) PostSignUpPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := Validate.Struct(&user); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			return
		}

		// Parse the form and check for an err
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
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
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
		// c.JSON(http.StatusOK, gin.H{"email": scs.Get("email"), "password": scs.Get("password")})
		c.HTML(http.StatusOK, "user-info.html", gin.H{"email": scs.Get("email"), "password": scs.Get("password")})
	}
}

func (ah *AppHandler) PostUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

	
		scs := sessions.Default(c)
		var user model.User

		validateErr := Validate.Struct(&user)
		if validateErr != nil {
			c.AbortWithError(http.StatusBadRequest, gin.Error{Err: validateErr})
			return
		}

		err := c.Request.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}

		userIPAddress := c.Request.RemoteAddr
		user.FirstName = c.Request.Form.Get("first-name")
		user.LastName = c.Request.Form.Get("last-name")
		user.YrsOfExp = c.Request.Form.Get("yrsofexp")
		user.Country = c.Request.Form.Get("nation")
		user.PhoneNumber = c.Request.Form.Get("phone")

		user.Stack = append(user.Stack, c.Request.Form.Get("stack-name"))

		user.IPAddress = userIPAddress
		user.Address = c.Request.Form.Get("address")

		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCtx()

		filter := bson.D{
			{Key: "email", Value: scs.Get("email")},
		}
		update := bson.D{{Key: "$set",Value: bson.D{
			{Key: "first_name", Value: user.FirstName},
			{Key: "last_name", Value: user.LastName},
			{Key: "address", Value: user.Address},
			{Key: "yrs_of_exp", Value: user.YrsOfExp},
			{Key: "country", Value: user.Country},
			{Key: "stack", Value: user.Stack},
			{Key: "ip_address", Value: user.IPAddress}}}}
		var mRes bson.M
		err = ah.UserCol.FindOneAndUpdate(ctx, filter, update).Decode(&mRes)
		if err == mongo.ErrNilDocument{
			return
		}
		log.Println(mRes)
		c.JSONP(http.StatusOK, gin.H{"result": mRes})
		// resultCursor := ah.UserCol.FindOne(ctx)
	}
}

// GetLoginPage LoginPage controller for the user to sign up for an account  and login as well
func (ah *AppHandler) GetLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// still have to design the login page

		c.HTML(http.StatusOK, "login-page.html", gin.H{})
	}
}

func (ah *AppHandler) PostLoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
