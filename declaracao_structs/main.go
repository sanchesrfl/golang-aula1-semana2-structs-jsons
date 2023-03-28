package main

import "fmt"

//structs

type Aluna struct {
	Nome      string
	Sobrenome string
	Matricula int16
	Cidade    string
}

func main() {

	aluna1 := Aluna{"aluna1", "sobrenome1", 12312, "Florianopolis"}
	aluna2 := Aluna{"aluna2", "sobrenome2", 12562, "São José"}

	fmt.Println(aluna1.Nome)
	fmt.Println(aluna1.Sobrenome)
	fmt.Println(aluna1.Matricula)
	fmt.Println(aluna1.Cidade)

	fmt.Println(aluna2.Nome)

}
