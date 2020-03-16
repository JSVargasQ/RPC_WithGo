package utils

type Request struct {
	datos []string
}

type Response1 struct {
	resultado1 []string
}

type Response2 struct {
	resultado2 [][]string
}

type Chat int

var usuarios []string

var mensajes [][]string 


func (c *Chat) registrarUsuario(rq *Request, rp1 *Response1, rp2 *Response2) error {

	usuarios = append(usuarios, rq.datos[0])
	rp1.resultado1 = usuarios
	return nil
}

func (c *Chat) guardarMensaje(rq *Request, rp1 *Response1, rp2 *Response2) error {

	fila := []string{rq.datos[0], rq.datos[1], rq.datos[2], rq.datos[3]}

	mensajes = append( mensajes, fila )

	rp2.resultado2 = mensajes
	return nil
}


