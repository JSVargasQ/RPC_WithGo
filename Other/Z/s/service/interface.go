package service

// Operations stands for all the remote fuctions definition in the service.
type Operations interface {
	Addition(*Request, *Response) error
	Subtraction(*Request, *Response) error
}
