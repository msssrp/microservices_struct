package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
  ID        primitive.ObjectID   `json:"id" bson:"_id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

