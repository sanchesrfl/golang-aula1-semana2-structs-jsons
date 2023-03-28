package main

import "fmt"

//criar uma struct que representa um retangulo
//altura e largua como atributos
// ser capaz de calcular sua área e seu perimetro

type Rectangle struct {
	Largura float64
	Altura  float64
}

func (r Rectangle) calculoDaArea() float64 {
	return r.Largura * r.Altura
}

func (r Rectangle) calculoDoPerimetro() float64 {
	return 2*r.Altura + 2*r.Largura
}

func main() {

	r := Rectangle{2.2, 1.5}

	fmt.Println("A área do retangulo é: ", r.calculoDaArea())
	fmt.Println("O perimetro do retangulo é: ", r.calculoDoPerimetro())

}
