package main

import (
	"log"
	"net/rpc"

)

// Calculator stands for the RPC client implementation.
type Calculator struct {
	Client *rpc.Client
}

// Request represents the data args for the service.
type Request struct {
	A, B float64
}

// Response represents the data results from the service.
type Response struct {
	Result float64
}


func main() {

	// Connecting to the server
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	c := &Calculator{Client: client}

	result, err := c.addition(5.6, 3.1)
	if err != nil {
		log.Println("Addition error: " + err.Error())
	} else {
		log.Printf("Addition result: %f", result)
	}

	result, err = c.subtraction(10.2, 3)
	if err != nil {
		log.Println("Subtraction error: " + err.Error())
	} else {
		log.Printf("Subtraction result: %f", result)
	}

}

// addition calls the Addition remote method from the calculator service.
func (c *Calculator) addition(a, b float64) (float64, error) {

	args := Request{A: a, B: b}
	var response Response
	err := c.Client.Call("Calculator.Addition", args, &response)
	return response.Result, err
}

// subtraction calls the Subtraction remote method from the calculator service.
func (c *Calculator) subtraction(a, b float64) (float64, error) {

	args := Request{A: a, B: b}
	var response Response
	err := c.Client.Call("Calculator.Subtraction", args, &response)
	return response.Result, err
}
