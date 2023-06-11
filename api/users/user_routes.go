package users

import "github.com/gorilla/mux"

func InitializeRoutes(r *mux.Router){
  r.Handle("", GetUsers()).Methods("GET")
  r.Handle("", CreateUser()).Methods("POST")
  r.Handle("/{id}", DeleteUser()).Methods("DELETE")
}

