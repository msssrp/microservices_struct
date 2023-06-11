package services_server

import (
	"log"
	"net/http"

	"github.com/msssrp/Go-Sandbox/api/users"
)

func StartUserSever() {
  userPort := "localhost:8081"
	userRoter := users.SetupUsersRouter()

  userSever := &http.Server{
    Addr: userPort,
    Handler: userRoter,
  }
  
  log.Println("user server running on port 8081")
  log.Fatal(userSever.ListenAndServe())
  
}
