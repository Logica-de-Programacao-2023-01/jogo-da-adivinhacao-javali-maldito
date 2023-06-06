package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cont := 0
	tentativas := []int{}
	var num, resp int
	fmt.Printf("\tJoguinho do mal\n")
	for resp != 2 {
		rand := rand.Intn(101)
		cont = 0
		for num != rand {
			fmt.Print("Digite um numero:")
			fmt.Scan(&num)
			if num > rand {
				fmt.Println("O numero é menor, digite novamente")
			} else if num < rand {
				fmt.Println("O numero é maior, digite novamente")
			} else {
				fmt.Println("PARABENS, vode acertou!!")
			}
			cont++
		}
		fmt.Println("Voce utilizou ", cont, " tentativas")
		tentativas = append(tentativas, cont)
		fmt.Print("deseja repetir(1.Sim/2.Nao)")
		fmt.Scan(&resp)
	}
	for i := 0; i < len(tentativas); i++ {
		fmt.Println("A quantidade de tentativas na rodada ", i+1, " é:", tentativas[i])
	}
}
