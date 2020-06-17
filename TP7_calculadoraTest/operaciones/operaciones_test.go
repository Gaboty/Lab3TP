package operaciones

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumarOk( t *testing.T ) {
	// GIVEN
	a := 5.0
	b := 4.0
	// WHEN
	resultado := Sumar( a, b )
	// THEN
	assert.Equal(t, 9.0, resultado, fmt.Sprintf("Fallo suma!!"))
}

func TestRestarOk( t *testing.T ) {
	// GIVEN
	a := 5.0
	b := 4.0
	// WHEN
	resultado := Restar( a, b )
	// THEN
	assert.Equal(t, 1.0, resultado, fmt.Sprintf("Fallo resta!!"))
}

func TestMultiplicarOk ( t *testing.T ) {
	// GIVEN
	a := 3.5
	b := 2.0
	// WHEN
	resultado := Multiplicar( a, b )
	// THEN
	assert.Equal(t, 7.0, resultado, fmt.Sprintf("Fallo Multiplicacion!!"))
}

func TestDividirOk ( t *testing.T ) {
	// GIVEN
	a := 10.0
	b := 2.5
	// WHEN
	resultado, err := Dividir( a, b )
	// THEN
	assert.Equal(t, 4.0, resultado, fmt.Sprintf("Fallo Division!!"))
	assert.Nil(t, err, fmt.Sprintf("Falla Division - error no nulo!!"))
}

func TestDividirPorCero( t *testing.T ) {
	// GIVEN
	a := 6.0
	b := 0.0
	// WHEN
	resultado, err := Dividir( a, b )
	// THEN
	assert.Equal(t, -1.0, resultado, fmt.Sprintf("Wrong div!!"))
	assert.NotNil(t, err, fmt.Sprintf("Wrong div - error nil!!"))
	assert.Equal(t, "no se puede dividir por cero", err.Error(), fmt.Sprintf("Wrong error message!!"))
}
