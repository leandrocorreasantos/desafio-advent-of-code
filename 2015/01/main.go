package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	andar := 0
	posicao := 0
	subsolo := 0

	arquivo, err := ioutil.ReadFile("input.txt")

	if err != nil{
		fmt.Println("Erro ao abrir arquivo")
		return
	}

	for i := 0; i < len(arquivo); i++ {
		posicao = posicao + 1;
		char := string(arquivo[i])
		if char == "(" {
			andar = andar + 1;
		}else if char == ")" {
			andar = andar - 1;
		}

		if subsolo == 0 {
			if andar == -1 {
				subsolo = posicao
			}
		}
	}

	fmt.Println("Resultado parte 1: ", andar)
	fmt.Println("Resultado parte 2: ", subsolo)
}
