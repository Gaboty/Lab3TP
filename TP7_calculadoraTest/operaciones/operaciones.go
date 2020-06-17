package operaciones

import "errors"

func Sumar( primerOperando float64, segundoOperando float64 ) ( float64 ) {
	return primerOperando + segundoOperando
}

func Restar( primerOperando float64, segundoOperando float64 ) ( float64 ) {
	return primerOperando - segundoOperando
}

func Multiplicar( primerOperando float64, segundoOperando float64 ) ( float64 ) {
	return primerOperando * segundoOperando
}

func Dividir( primerOperando float64, segundoOperando float64 ) ( float64, error ) {
	if segundoOperando == 0 {
		return -1, errors.New("no se puede dividir por cero")
	}
	return primerOperando / segundoOperando, nil
}
