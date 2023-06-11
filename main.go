package main

import (
	"github.com/msssrp/Go-Sandbox/config"
	"github.com/msssrp/Go-Sandbox/services_server"
)

func main() {

  go services_server.StartUserSever()
  

  config.MainServer()

 }
