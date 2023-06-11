package services_server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartUserSever(us *mux.Router) {

  userPort := "localhost:8081"

  userSever := &http.Server{
    Addr: userPort,
    Handler: us,
  }
  
  log.Println("user server running on port 8081")
  log.Fatal(userSever.ListenAndServe())
  
}
