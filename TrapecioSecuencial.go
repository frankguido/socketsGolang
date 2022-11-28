package main

//librerias time ttiempo, log etc
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// tipos de datos para calcular el area del trapecio
type Trapecio struct {
	base float64
	BASE float64
	AREA float64
	aX   float64
	bX   float64
	h    float64
}

// funcion que halla el area de un solo trapecio
func TrapecioSec(para1 float64, para2 float64, para3 float64, para4 float64) *Trapecio {
	t := new(Trapecio)
	t.base = para1
	t.BASE = para2
	t.aX = para3
	t.bX = para4
	t.h = t.bX - t.aX
	t.AREA = ((t.base + t.BASE) * t.h) / 2
	return t
}

// funcion que retorna el area de un trapecio
func (t Trapecio) GETAREA() float64 {
	return t.AREA
}

// llenado de datos para el calculo de el area de n trapecios de forma parcial y total

func main() {
	var a float64 = 5
	var b float64 = 20
	var h float64 = b - a
	var a2 float64
	var funcion = func(x float64) float64 {
		return (x*x + 1) / 2
	}
	var trapecios int = 100
	var channelsTrapecio []Trapecio
	ttotal := 0
	var ATOTAL float64 = 0
	var APARCIAL float64 = 0
	var delta float64 = (h / float64(trapecios))
	ttrac := 0
	ff, err := os.Create("datasecuencial.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer ff.Close()

	for f := 1; f <= trapecios; f++ {
		fmt.Print("----------\n")
		fmt.Print("a: ", a, "\n")
		a2 = a
		a = a + delta
		b = a
		fmt.Print("b: ", b, "\n")

		start := time.Now()
		channelsTrapecio = append(channelsTrapecio, *TrapecioSec(funcion(a2), funcion(b), a2, b))

		APARCIAL = channelsTrapecio[len(channelsTrapecio)-1].GETAREA()
		ATOTAL += APARCIAL
		fmt.Print("Area Trapecio Secuencial ", f, " : ", APARCIAL, " \n ")
		tiempo := time.Since(start)
		fmt.Printf("El tiempo es :", tiempo)
		ttrac = int(tiempo)
		ttotal += int(tiempo)
		_, err2 := ff.WriteString(strconv.Itoa(f) + " " + strconv.Itoa(ttrac) + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}

	}
	fmt.Print("---------------\n ")
	fmt.Print(" Area Total ", trapecios, " trapecios secuencial", ATOTAL, "\n")
	fmt.Print(" Tiempo Total ", ttotal, "\n")

}
