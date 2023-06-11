package database

import (
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://siripoomcontact:hidUbX6HFfoFZWqB@cluster0.hztprdy.mongodb.net/"

func ConnectMongo() (*mongo.Database , error){
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

  client , err := mongo.Connect(context.TODO(), opts)
  if err != nil{
    return nil , err
  }

  err = client.Ping(context.TODO(), nil)
  if err != nil{
    client.Disconnect(context.TODO())
    return nil , err
  }

  return client.Database("admin",nil) , nil
}
