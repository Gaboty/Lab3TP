package TP6_calculadoraParalela

import (
	"fmt"
	"sync"
	"time"
)

var desviacion float32 = 2.0
var stop sync.WaitGroup


func CalculadoraConHilos() {
	stop.Add(4)

	var ch1 chan float32 = make(chan float32)
	var ch2 chan float32 = make(chan float32)
	var ch3 chan float32 = make(chan float32)
	var ch4 chan float32 = make(chan float32)

	var num1 float32 = 0.0
	var num2 float32 = 1.0
	println("Ingrese el valor para el primer operando")
	fmt.Scan(&num1)
	println("Ingrese el valor para el segundo operando")
	fmt.Scan(&num2)

	go Sumar(num1, num2, ch1)
	go Restar(num1, num2, ch2)
	go Multiplicar(num1, num2, ch3)
	go Dividir(num1, num2, ch4)

	go func() {
		for i := 0; i < 4; i++ {
			select {
			case suma := <-ch1:
				fmt.Printf("Suma: ")
				fmt.Println(suma)
				defer stop.Done()
			case resta := <-ch2:
				fmt.Printf("Resta: ")
				fmt.Println(resta)
				defer stop.Done()
			case multiplicar := <-ch3:
				fmt.Printf("Multiplicacion: ")
				fmt.Println(multiplicar)
				defer stop.Done()
			case dividir := <-ch4:
				fmt.Printf("Division: ")
				fmt.Println(dividir)
				defer stop.Done()
			}
		}
	}()

	stop.Wait()
}

func Sumar( primerOperando float32, segundoOperando float32, ch1 chan float32 ){
	ch1 <- primerOperando + segundoOperando
	time.Sleep(time.Second * 1)
}

func Restar( primerOperando float32, segundoOperando float32, ch2 chan float32 ){
	ch2 <- primerOperando - segundoOperando
	time.Sleep(time.Second * 1)
}

func Multiplicar( primerOperando float32, segundoOperando float32, ch3 chan float32 ){
	ch3 <- primerOperando * segundoOperando
	time.Sleep(time.Second * 1)
}

func Dividir( primerOperando float32, segundoOperando float32, ch4 chan float32 ){
	if segundoOperando == 0 {
		ch4 <- -1.0
	}
	ch4 <- primerOperando / segundoOperando
	time.Sleep(time.Second * 1)
}
