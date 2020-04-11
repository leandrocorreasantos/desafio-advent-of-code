package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	andar := 0

	arquivo, err := ioutil.ReadFile("input.txt")

	if err != nil{
		fmt.Println("Erro ao abrir arquivo")
		return
	}

	for i := 0; i < len(arquivo); i++ {
		char := string(arquivo[i])
		if char == "(" {
			andar = andar + 1;
		}else if char == ")" {
			andar = andar - 1;
		}
	}

	fmt.Println("Resultado: ", andar)
}
