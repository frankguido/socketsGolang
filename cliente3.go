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

func cliente(persona Persona) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
func main() {
	persona := Persona{
		Nombre: "Fran Guido",
		Email: []string{
			"fgfg@ulasalle.edu.pe",
			"ggggg@ulasalle.edu.pe",
		},
	}
	go cliente(persona)

	var input string
	fmt.Scanln(&input)
}
