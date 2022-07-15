package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

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
	//MailCol   *mongo.Collection
}

var AppRepo *Repository

func NewRepository(appConfig *config.AppConfig, userCol *mongo.Collection) *Repository {
	return &Repository{
		AppConfig: appConfig,
		UserCol:   userCol,
		//MailCol:   mailCol,
	}
}
func NewHandler(r *Repository) {
	r = AppRepo
}

//HomePage handler to display the home page of the application
func (rp *Repository) HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {

		//var user model.User
		//if err := c.BindJSON(&user); err != nil {
		//	c.Header("Content-Type", "application/json")
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot verify user model"})
		//	return
		//}

		ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancelCtx()
		var documents = []interface{}{
			bson.D{
				{"first_name", "Yusuf"},
				{"last_name", "Akinleye"},
				{"email", "ayaaakinleye@gmail,com"},
				{"stack", bson.A{"Go lang", "Python", "MongoDB"}},
				{"project_details", bson.A{"track-space", bson.A{"gin-framework", "mongoDB Atlas", "Built-ins Packages"}, time.Now(), time.Now(), time.Second, "Work-Management System"}},
			},
		}

		result, err := rp.UserCol.InsertOne(ctx, documents)
		if err != nil {
			log.Println("Cannot insert document in the database")
			return
		}
		log.Println(result.InsertedID)
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

//LoginPage handler for the user to sign up for an account  and login as well
func (rp *Repository) LoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.page.tmpl", gin.H{})
	}
}
