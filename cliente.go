package main

//paquete net para crear el servidor tcp
import (
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
	c.Write([]byte(msg))
	c.Close()
}
func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
