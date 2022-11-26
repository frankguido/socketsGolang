package main

//paquete net para crear el servidor tcp
import (
	"encoding/gob"
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
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", msg) //comvierte de bits a texto

	}

}
func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)

}
