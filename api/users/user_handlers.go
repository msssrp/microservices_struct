package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/msssrp/Go-Sandbox/services/user"
)

func GetUsers() http.HandlerFunc{
  return func(w http.ResponseWriter , r *http.Request){
    users , err := user.GetAllUsers()
  
    if err != nil{
      w.Header().Set("Content-Type","application/json")
      w.WriteHeader(http.StatusInternalServerError)
      errmsg := []byte(err.Error())
      w.Write(errmsg)
    }
  
    respone , err := json.Marshal(users)
  
    if err != nil{
      log.Fatal(err)
    }
  

    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(respone) 
  }
}

func CreateUser() http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type","application/json")

    result , err := user.CreateUser(w, r)
    if err != nil{
      w.WriteHeader(http.StatusInternalServerError)
      errmsg := []byte(err.Error())
      w.Write(errmsg)
    }
    
    respone , err := json.Marshal(result)

    if err != nil{
      http.Error(w, "faild to parse to json", http.StatusInternalServerError)
    }

    w.WriteHeader(http.StatusOK)
    w.Write(respone)
    
  }
}

func DeleteUser() http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {

    
    result , err := user.DeleteUser(w, r)
    if err != nil{
      errmsg := []byte(err.Error())
      w.WriteHeader(http.StatusInternalServerError)
      w.Write(errmsg)
      return
    }

    if result.DeletedCount == 0{
      http.Error(w, "user not found", http.StatusNotFound)
      return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("user Deleted"))    
  }
}
