package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Chat struct {
	Usuarios []string
	Mensajes [][]string
}

var nickname string

func mainLoop( pLector *bufio.Reader ) {
	
	lector :=  pLector

	for {

		entrada, error := lector.ReadString('\n')

		if error != nil {
			log.Printf("Error: %q\n", error)
		}

		log.Println(entrada)
		
		if strings.HasPrefix(entrada, "salir") {

			//client.Call("APP.UsuarioSalir", nickname , &reply)
			break
		}
	}	
}

func main() {

	var reply Chat

	client, err := rpc.DialHTTP("tcp", "localhost:4040")



	if err != nil {
		log.Fatal("Error de conexion: ", err)
	}

	lector := bufio.NewReader(os.Stdin)

	log.Println("Ingrese su NickName:")
	nickname, err :=  lector.ReadString('\n')
	nickname = strings.TrimSpace(nickname)

	if err != nil {
		log.Fatal("Error con el username: ", err)
	}

	err = client.Call("APP.UsuarioExiste", nickname , &reply )

	if err != nil {
		log.Println("ERROR")
		log.Fatal("Nickname en uso. Intente con otro")
	} 

	// client.Call("APP.ObtenerDatos", "" , &reply)

	client.Call("APP.RegistrarUsuario", nickname , &reply )
	log.Println("Bienvenido " + nickname + "\n")


	log.Printf("Usuarios ON:")
	for i := range reply.Usuarios {
		log.Println(reply.Usuarios[i])
	}

	mainLoop(lector)


}
