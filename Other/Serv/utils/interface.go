package utils

type Operations interface {

	registrarUsuario(*Request, *Response1, *Response2)
	guardarMensaje(*Request, *Response1, *Response2)

}
