package utils

type Operations interface {

	registrarUsuario(*Request, *Response)
	guardarMensaje(*Request, *Response)

}
