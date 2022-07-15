package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yusuf/track-space/pkg/config"
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

// HomePage handler to display the home page of the application
func (rp *AppHandler) HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancelCtx()
		documents := []interface{}{
			bson.D{
				{Key: "first_name", Value: "Yusuf"},
				{Key: "last_name", Value: "Akinleye"},
				{Key: "email", Value: "ayaaakinleye@gmail,com"},
				{Key: "stack", Value: bson.A{"Go lang", "Python", "MongoDB"}},
				{Key: "project_details", Value: bson.A{"track-space", bson.A{"gin-framework", "mongoDB Atlas", "Built-ins Packages"}, time.Now(), time.Now(), time.Second, "Work-Management System"}},
			},
			bson.D{
				{Key: "first_name", Value: "Yusuf"},
				{"last_name", "Akinleye"},
				{"email", "ayaaakinleye@gmail,com"},
				{"stack", bson.A{"Go lang", "Python", "MongoDB"}},
				{"project_details", bson.A{"track-space", bson.A{"gin-framework", "mongoDB Atlas", "Built-ins Packages"}, time.Now(), time.Now(), time.Second, "Work-Management System"}},
			},
		}

		result, err := rp.UserCol.InsertMany(ctx, documents)
		if err != nil {
			log.Println(err)
			log.Println("Cannot insert document in the database")
			return
		}
		log.Println(result.InsertedIDs)
		//var result bson.M
		//
		//err := rp.UserCol.FindOne(ctx, bson.D{{"first_name", "Yusuf"}}).Decode(&result)
		//if err != nil {
		//	log.Println(mongo.ErrNoDocuments)
		//	return
		//}
		////a := c.Request.RemoteAddr
		c.IndentedJSON(http.StatusOK, result)
		a := c.Request.RemoteAddr
		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{"ip": a})
	}
}

// LoginPage handler for the user to sign up for an account  and login as well
func (rp *AppHandler) LoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}
