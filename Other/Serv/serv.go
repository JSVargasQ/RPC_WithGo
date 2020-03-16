package main

import (
	"log"
	//"net"
	//"net/http"
	"net/rpc"

	"./utils"
)

func main() {

	infChat := new(utils.Chat)
	error := rpc.Register(infChat)

	if error != nil {
		log.Fatal("Error de servicio de registro:", error)
	}

	rpc.HandleHTTP()
}