package users

import "github.com/gorilla/mux"

func InitializeRoutes(r *mux.Router){
  r.Handle("/users", GetUsers()).Methods("GET")
  r.Handle("/users", CreateUser()).Methods("POST")
  r.Handle("/users/{id}", DeleteUser()).Methods("DELETE")
}

