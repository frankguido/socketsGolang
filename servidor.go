package main

//paquete net para crear el servidor tcp
import (
	"fmt"
	"net"
)

// servidor para abrir un puerto
func servidor() {
	s, err := net.Listen("tcp", ":9999") //protocolo y puerto va a utilizar
	if err != nil {
		fmt.Println(err)
		return
	}
	//for para mantenerse en escucha
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}
}

// se recibben en bits y luewgo se convierten aun mensaje que se pueda leer
func handleCliente(c net.Conn) {
	b := make([]byte, 100)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", b[:bs])  //comvierte de bits a texto
		fmt.Println("Bytes leidos: ", bs) //imprimew los bytres leidos
	}

}
func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)

}
