package config

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/msssrp/Go-Sandbox/api/users"
	"github.com/msssrp/Go-Sandbox/services_server"
)

func MainServer() {
	mainServer := mux.NewRouter()


  /*incase of using middleware for handler somethin
  r.Use(api.Authorization)*/


  userRouter := UserRouter(mainServer)
  go services_server.StartUserSever(userRouter)

  srv := &http.Server{
    Handler : mainServer,
    Addr: "localhost:8080",
    WriteTimeout: 5*time.Second,
    ReadTimeout: 5*time.Second,
  }


  log.Printf("main server Running on port 8000")
  log.Fatal(srv.ListenAndServe())

}

func UserRouter(ms *mux.Router) *mux.Router{
  userRouter := ms.PathPrefix("/users").Subrouter()
  users.InitializeRoutes(userRouter)
  return userRouter
}
