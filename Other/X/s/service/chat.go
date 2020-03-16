package service

import (
	"log"
)

type Chat int


func (c *Chat) nuevoUsuario(rq *Request, rp *Response) error {

	log.Println("Se registro.")
	rp.resultado = rq.datos[0] + " ha entrado en el chat."
	return nil

}


