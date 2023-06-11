package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msssrp/Go-Sandbox/database"
	"github.com/msssrp/Go-Sandbox/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers() ([]*User , error) {
	db, err := database.ConnectMongo()
  if err != nil{
    log.Println(err)
  }

  coll := db.Client().Database("myDatabase").Collection("Users")
  
  users := []*User{}
  
  cursor ,err := coll.Find(context.TODO(),bson.M{})
  if err != nil{
    return nil , err
  }

  if err := cursor.All(context.TODO(), &users); err != nil{
    return nil,err
  }


  return users , nil
}

func CreateUser(w http.ResponseWriter , r *http.Request) (*mongo.InsertOneResult , error){

  db , err := database.ConnectMongo()
  if err != nil{
    return nil , err
  }

  var newUser models.UserSchema

  if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
    return nil , err
  }
  
  coll := db.Client().Database("myDatabase").Collection("Users")

  result ,err := coll.InsertOne(context.TODO(),&newUser)
  if err != nil{
    return nil , err
  }
  

  return result , nil

}

func DeleteUser(w http.ResponseWriter , r *http.Request) (*mongo.DeleteResult , error){
  
  vars := mux.Vars(r)
  userID := vars["id"]

  objID , err := primitive.ObjectIDFromHex(userID)
  if err != nil{
    return nil , err
  }

  db , err := database.ConnectMongo()
  if err != nil{
    return nil , err
  }

  coll := db.Client().Database("myDatabase").Collection("Users")
  filter := bson.M{"_id": objID}

  result,err := coll.DeleteOne(context.TODO(), filter)

  if err != nil{
    return nil , err
  }


  return result , nil
}
