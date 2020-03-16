package service

type Operations interface {

	nuevoUsuario(*Request, *Response) error

}
