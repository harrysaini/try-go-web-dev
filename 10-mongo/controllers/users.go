package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/10-mongo/models"
	"github.com/harrysaini/try-go-web-dev/10-mongo/models/requests"
	"github.com/harrysaini/try-go-web-dev/10-mongo/utils"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController handle user routes
type UserController struct {
	usersCollection *mongo.Collection
}

// NewUserController create new intance
func NewUserController(db *mongo.Database) *UserController {
	usersCollection := db.Collection("users")

	return &UserController{
		usersCollection,
	}
}

// Index route controller
func (uc *UserController) Index(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	io.WriteString(w, "Hello world")
}

// NewUser create new user
func (uc *UserController) NewUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := req.Context()
	userData := requests.CreateUser{}

	err := json.NewDecoder(req.Body).Decode(&userData)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:   userData.Name,
		Gender: userData.Gender,
		Age:    userData.Age,
		ID:     primitive.NewObjectID(),
	}

	insertResults, err := uc.usersCollection.InsertOne(ctx, user)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	log.Printf("Inserted row with id %v", insertResults.InsertedID)

	utils.SendResponse(w, user)

}

// GetUsers fetches list of all users
func (uc *UserController) GetUsers(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := req.Context()
	filter := bson.D{{}}

	count, err := uc.usersCollection.CountDocuments(ctx, filter)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// bhai jyada optimisation bhi danger hai, len 0 is very important here
	users := make([]models.User, 0, count)

	cursor, err := uc.usersCollection.Find(ctx, filter)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	for cursor.Next(ctx) {
		var user models.User

		err := cursor.Decode(&user)
		if err != nil {
			utils.SendErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	cursor.Close(ctx)

	utils.SendResponse(w, users)

}

// FindUser find user by id
func (uc *UserController) FindUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := req.Context()
	id := params.ByName("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	var user models.User

	filter := bson.M{"_id": objectID}

	err = uc.usersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, user)

}

// UpdateUser update user
func (uc *UserController) UpdateUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := req.Context()
	id := params.ByName("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	userData := make(map[string]interface{})

	err = json.NewDecoder(req.Body).Decode(&userData)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": userData,
	}

	updateResult, err := uc.usersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	var user models.User
	err = uc.usersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, user)

}

// DeleteUser
func (uc *UserController) DeleteUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := req.Context()
	id := params.ByName("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	filter := bson.M{"_id": objectID}

	deleteRes, err := uc.usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if deleteRes.DeletedCount == 0 {
		err := errors.New("No user found")
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.SendResponse(w, nil)

}
