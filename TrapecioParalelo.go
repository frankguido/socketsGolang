package main

//paquetes importados necesarios para la ejecucion paralela
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// tipos de datos para los trapecios
type Trapecio2 struct {
	base float64
	bX   float64
	h    float64
	AREA float64
	tmpo int
	BASE float64
	aX   float64
}

func TrapecioParalelo(para1 float64, para2 float64, para3 float64, para4 float64) *Trapecio2 {
	t := new(Trapecio2)

	t.base = para1
	t.BASE = para2
	t.aX = para3
	t.bX = para4
	t.h = t.bX - t.aX
	t.AREA = ((t.base + t.BASE) * t.h) / 2

	return t
}
func calcular() {

	t := new(Trapecio2)
	t.h = t.bX - t.aX
	t.AREA = ((t.base + t.BASE) * t.h) / 2

}

func (t Trapecio2) GETTIEMPO() int {
	return t.tmpo
}

func (t Trapecio2) GETAREA2() float64 {
	go calcular()
	return t.AREA
}
func funcion(x float64) float64 {

	return (x*x + 1) / 2
}
func main() {
	var a float64 = 5
	var b float64 = 20
	var h float64 = b - a
	var a2 float64
	strr := 0
	var trapecios int = 100
	var channelsTrapecio2 []Trapecio2
	ttotal := 0
	var ATOTAL float64 = 0
	var APARCIAL float64 = 0
	var delta float64 = (h / float64(trapecios))

	ff, err := os.Create("dataParalelo.txt")

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

		channelsTrapecio2 = append(channelsTrapecio2, *TrapecioParalelo(funcion(a2), funcion(b), a2, b))

		channelsTrapecio2[len(channelsTrapecio2)-1].GETAREA2()

	}

	for r := 0; r <= len(channelsTrapecio2); r++ {
		start := time.Now()
		APARCIAL = channelsTrapecio2[r].GETAREA2()
		tiempo := time.Since(start)
		fmt.Print("Area Trapecio ", APARCIAL, " \n ")
		fmt.Print("Tiempo ", tiempo, " \n ")
		strr = int(tiempo)
		_, err2 := ff.WriteString(strconv.Itoa(r) + " " + strconv.Itoa(strr) + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		ATOTAL += APARCIAL
	}
	fmt.Print("---------------\n ")
	fmt.Print(" Area Total ", trapecios, " trapecios ", ATOTAL, "\n")
	fmt.Print(" Tiempo Total ", ttotal, "\n")

}
