package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
  FirstName string `bson:"firstName,omitempty"`
  LastName string `bson:"lastName,omitempty"`
  Email string `bson:"email,omitempty"`
  Password string `bson:"password,omitempty"`
}
