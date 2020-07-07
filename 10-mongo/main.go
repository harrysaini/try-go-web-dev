package main

import (
	"context"
	"log"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/10-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := httprouter.New()
	userController := controllers.NewUserController(getDatabaseConnection())

	router.GET("/", userController.Index)
	router.POST("/users", userController.NewUser)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.FindUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	http.ListenAndServe(":8090", router)
}

func getDatabaseConnection() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("test_db")

	return db
}
