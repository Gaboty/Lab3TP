package TP4_calculadora

import (
	"fmt"
	"github.com/Gaboty/Lab3TP/TP4_calculadora/operaciones"
)

var desviacion float32 = 2.0

func Calculadora() {

	var num1 float32 = 0.0
	var num2 float32 = 1.0
	println("Ingrese el valor para el primer operando")
	fmt.Scan(&num1)
	println("Ingrese el valor para el segundo operando")
	fmt.Scan(&num2)

	resultado := Resultados{
		resultadoSuma: operaciones.Sumar(num1, num2),
		resultadoResta: operaciones.Restar(num1, num2),
		resultadoMultiplicacion: operaciones.Multiplicar(num1, num2),
		resultadoDivision: operaciones.Dividir(num1, num2),
	}
	fmt.Printf("Suma: ")
	fmt.Print( resultado.resultadoSuma )
	fmt.Printf("\n")
	fmt.Printf("Resta: ")
	fmt.Print( resultado.resultadoResta )
	fmt.Printf("\n")
	fmt.Printf("Multiplicacion: ")
	fmt.Print( resultado.resultadoMultiplicacion )
	fmt.Printf("\n")
	fmt.Printf("Dividir: ")
	fmt.Print( resultado.resultadoDivision )
	fmt.Printf("\n")
	fmt.Printf("Estructura bugeada: ")
	fmt.Println( fmt.Sprintf( "%v", CambiarResultados( resultado ) ) )

}

type Resultados struct {

	resultadoSuma float32
	resultadoResta float32
	resultadoMultiplicacion float32
	resultadoDivision float32

}

func CambiarResultados( resultado Resultados ) Resultados {
	return Resultados{
		resultadoSuma: resultado.resultadoSuma + desviacion,
		resultadoResta: resultado.resultadoResta - desviacion,
		resultadoMultiplicacion: resultado.resultadoMultiplicacion * desviacion,
		resultadoDivision: resultado.resultadoDivision / desviacion,
	}
}
