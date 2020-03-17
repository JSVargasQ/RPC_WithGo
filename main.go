package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Chat struct {
	Usuarios []string
	Mensajes [][]string
}

type APP int


var chatRoom Chat


func (a *APP) ObtenerDatos(empty string, reply *Chat) error {
	*reply = chatRoom
	return nil
}



func (a *APP) RegistrarUsuario(pUserName string, reply *Chat) error {
	chatRoom.Usuarios = append(chatRoom.Usuarios, pUserName)
	log.Println("El usuario " + pUserName + " ha entrado al chat.")
	*reply = chatRoom
	return nil
}

func (a *APP) RegistrarMensaje(pMensaje []string, reply *Chat) error {
	
	fila := []string{pMensaje[0], pMensaje[1]}

	chatRoom.Mensajes = append(chatRoom.Mensajes, fila)

	*reply = chatRoom
	return nil
}

func (a *APP) UsuarioExiste(pUserName string, replt *Chat)  error {

	for i := range chatRoom.Usuarios {
		if chatRoom.Usuarios[i] == pUserName {
			
			return errors.New("Nickname existe")
		}
	}
	return nil
}

func (a *APP) UsuarioSalir(pUserName string, replt *Chat)  error {

	for i := range chatRoom.Usuarios {
		if chatRoom.Usuarios[i] == pUserName {
			chatRoom.Usuarios = append(chatRoom.Usuarios[:i], chatRoom.Usuarios[i+1:]...)
			break
		}
	}
	return nil
}

func main() {

	api := new(APP)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("initial database: ", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("fourth database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)

}


