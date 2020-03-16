package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"./service"	
)

func main() {

	s := new(service.Calculator)
	err := rpc.Register(s)
	if err != nil {
		log.Fatal("Service registration error:", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("Server started!")
	log.Println(l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
