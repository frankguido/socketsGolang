package main

//paquete net para crear el servidor tcp
import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

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
	var persona Persona
	err := gob.NewDecoder(c).Decode(&persona)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", persona) //comvierte de bits a texto

	}

}
func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)

}
