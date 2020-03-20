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

var ultimoMensaje int

func verificarMensajes(pClient *rpc.Client, reply Chat, pNick string) {

	for {

		pClient.Call("APP.ObtenerMensajes", "" , &reply)

		msg := reply.Mensajes

		for i := ultimoMensaje; i <= len(msg)-1; i++ {

			if i == ultimoMensaje {
				continue
			} else if msg[i][1] == pNick {
				log.Println("Tu: " +  msg[i][0])
			} else {
				log.Println(msg[i][1] +  " dice: " +  msg[i][0])
			}

		}

		ultimoMensaje = len(msg)-1

	}

}

func mainLoop( pLector *bufio.Reader, pClient *rpc.Client, pNick string, pReply Chat ) {
	
	reply := pReply
	nick := pNick
	lector :=  pLector

	ultimoMensaje = 0;
	go verificarMensajes(pClient, reply, nick )

	for {

		entrada, error := lector.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		if error != nil {
			log.Printf("Error: %q\n", error)
		}

		if strings.HasPrefix(entrada, "salir") {

			pClient.Call("APP.UsuarioSalir", nick , &reply)
			break
		} else {

			pClient.Call("APP.RegistrarMensaje", []string {entrada, nick} , &reply)
		}
	}	
}

func main() {

	var reply Chat


	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	//log.Println(reflect.TypeOf(client))

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

		if reply.Usuarios[i] == nickname {
			log.Println(reply.Usuarios[i] + " (Yo)") 
		} else {
			log.Println(reply.Usuarios[i])
		} 

		log.Println("")
	}

	

	mainLoop(lector, client, nickname, reply)


}
