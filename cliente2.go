package main

//paquete net para crear el servidor tcp
import (
	"encoding/gob"
	"fmt"
	"net"
)

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hola mundo"
	fmt.Println(msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
