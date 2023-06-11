package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUsersRouter() http.Handler {
  
  userRouter := mux.NewRouter()

  InitializeRoutes(userRouter)

  return userRouter
}
