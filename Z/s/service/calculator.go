package service

import (
	"log"
)

// Calculator represents the service available over the network.
type Calculator int

// Addition stands as an exported method that execute a sum
// operation over the request args.
func (c *Calculator) Addition(rq *Request, rp *Response) error {

	log.Printf("Executing addition with args: %+v", rq)
	rp.Result = rq.A + rq.B
	return nil

}

// Subtraction stands as an exported method that execute a subtraction
// operation over the request args.
func (c *Calculator) Subtraction(rq *Request, rp *Response) error {

	log.Printf("Executing subtraction with args: %+v", rq)
	rp.Result = rq.A - rq.B
	return nil

}
