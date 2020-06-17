package TP7_calculadoraTest

import (
	"fmt"
	"github.com/Gaboty/Lab3TP/TP7_calculadoraTest/operaciones"
)

type Resultados struct {

	resultadoSuma float64
	resultadoResta float64
	resultadoMultiplicacion float64
	resultadoDivision float64

}

func CalculadoraTest() {

	result, error := operaciones.Dividir( 4.0, 3.0)
	if error != nil {
		print(error)
		return
	}
	resultado := Resultados{
		resultadoSuma: operaciones.Sumar(2.0,3.0),
		resultadoResta: operaciones.Restar(2.0, 4.0),
		resultadoMultiplicacion: operaciones.Multiplicar( 5.0, 10.0),
		resultadoDivision: result,
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

func CambiarResultados( resultado Resultados ) Resultados {
	return Resultados{
		resultadoSuma: resultado.resultadoSuma + 2,
		resultadoResta: resultado.resultadoResta - 2,
		resultadoMultiplicacion: resultado.resultadoMultiplicacion * 2,
		resultadoDivision: resultado.resultadoDivision / 2,
	}
}
