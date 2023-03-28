package main

/*
Converter struct para Json
*/

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {

	//cria nova pessoa
	p := Pessoa{"Rafa", 18}

	//conversao para json

	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("pessoa", p)
	fmt.Println("pesso to json", j)

	jString := string(j)
	fmt.Println("json to string", jString)

}
