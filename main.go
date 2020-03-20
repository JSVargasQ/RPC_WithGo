package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"bufio"
	"os"
	"strings"
	"fmt"
	"time"
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

func (a *APP) ObtenerMensajes(empty string, reply *Chat) error {
	*reply = chatRoom
	return nil
}



func (a *APP) RegistrarUsuario(pUserName string, reply *Chat) error {

	chatRoom.Usuarios = append(chatRoom.Usuarios, pUserName)
	log.Println("El usuario " + pUserName + " ha entrado al chat.")

	fila := []string{"El usuario " + pUserName + " ha entrado al chat.", pUserName, "1"}
	chatRoom.Mensajes = append(chatRoom.Mensajes, fila)

	*reply = chatRoom
	return nil
}

func (a *APP) RegistrarMensaje(pMensaje []string, reply *Chat) error {
	
	fila := []string{pMensaje[0], pMensaje[1], "0"}

	chatRoom.Mensajes = append(chatRoom.Mensajes, fila)

	log.Println("El usuario " + pMensaje[1] + " dice: " + pMensaje[0])

	*reply = chatRoom
	return nil
}

func (a *APP) UsuarioExiste(pUserName string, replt *Chat)  error {

	if strings.TrimSpace(pUserName) == "" {

		return errors.New("Nickname no valido.")
	
	}

	for i := range chatRoom.Usuarios {
		if chatRoom.Usuarios[i] == pUserName {
			
			return errors.New("El Nickname ya esta en uso.")
		}
	}
	return nil
}

func (a *APP) UsuarioSalir(pUserName string, replt *Chat)  error {

	for i := range chatRoom.Usuarios {
		if chatRoom.Usuarios[i] == pUserName {
			chatRoom.Usuarios = append(chatRoom.Usuarios[:i], chatRoom.Usuarios[i+1:]...)
			log.Println("El usuaio " + pUserName + " ha abandonado el chat.")
			break
		}
	}
	return nil
}

func input() {

	lector := bufio.NewReader(os.Stdin)
	entrada, error := lector.ReadString('\n')
	entrada = strings.TrimSpace(entrada)

	if error != nil {
		log.Printf("Error: %q\n", error)
	}

	if strings.HasPrefix(entrada, "/apagar") || strings.HasPrefix(entrada, "/APAGAR") {

		fila := []string{"El servidor se ha dado de baja", "", "1"}
		chatRoom.Mensajes = append(chatRoom.Mensajes, fila)

		fila = []string{"Para salir de la conversación ingresa '/salir'", "", "1"}
		chatRoom.Mensajes = append(chatRoom.Mensajes, fila)

		time.Sleep(3 * time.Second)

		fmt.Println("")
		fmt.Println("El servidor se da de baja.")
		
		os.Exit(1)

	}
	
}


func main() {

	api := new(APP)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Error de escucha", err)
	}

	log.Printf("Servidor RPC en el puerto %d", 4040)
	
	fmt.Println("Para apagar el servidor ingresa '/apagar'")
	fmt.Println("")

	go input()
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	

}


