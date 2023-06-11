package config

import (
	"net/http"
	"time"
  "log"
	"github.com/gorilla/mux"
)

func MainServer() {
	r := mux.NewRouter()


  /*incase of using middleware for handler somethin
  r.Use(api.Authorization)*/


  srv := &http.Server{
    Handler : r,
    Addr: "localhost:8080",
    WriteTimeout: 5*time.Second,
    ReadTimeout: 5*time.Second,
  }

  
  log.Printf("main server Running on port 8000")
  log.Fatal(srv.ListenAndServe())

}
