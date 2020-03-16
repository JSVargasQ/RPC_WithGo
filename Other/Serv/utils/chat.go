package utils


var usuarios []string

var mensajes [][]string 

type Chat int

func registrarUsuario(pUserName string) {

	usuarios = append(usuarios, pUserName)

}

func guardarMensaje(pMensaje string, pUsuario string, pDate string, pIp string) {

	fila := []string{pMensaje, pUsuario, pDate, pIp}

	mensajes = append( mensajes, fila )
}


