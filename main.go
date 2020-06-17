package main

import (
	"fmt"
	"github.com/Gaboty/Lab3TP/TP3_helloWorld"
	"github.com/Gaboty/Lab3TP/TP4_calculadora"
	"github.com/Gaboty/Lab3TP/TP5_consumoApis"
	"github.com/Gaboty/Lab3TP/TP6_calculadoraParalela"
	"github.com/Gaboty/Lab3TP/TP7_calculadoraTest"
	"github.com/Gaboty/Lab3TP/TP8_calculadoraApi"
)

func main(){
	var numeroTP int64
	fmt.Printf("Ingrese el numero del TP que quiere ejecutar: \n")
	fmt.Printf("3): Introduccion a Go \n")
	fmt.Printf("4): Variables y estructuras en Go \n")
	fmt.Printf("5): Consumo de APIs en Go \n")
	fmt.Printf("6): Concurrencia y Paralelirmos en Go \n")
	fmt.Printf("7): Testing \n")
	fmt.Printf("8): Creacion de una API \n\n ")
	fmt.Scan(&numeroTP)
	switch numeroTP {
	case 3:
		TP3_helloWorld.HelloWord()
		break
	case 4:
		TP4_calculadora.Calculadora()
		break
	case 5:
		TP5_consumoApis.ConsumoApi()
		break
	case 6:
		TP6_calculadoraParalela.CalculadoraConHilos()
		break
	case 7:
		TP7_calculadoraTest.CalculadoraTest()
		break
	case 8:
		TP8_calculadoraApi.CalculadoraApi()
		break
	}
}
