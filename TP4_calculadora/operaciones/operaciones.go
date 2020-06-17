package operaciones

func Sumar( primerOperando float32, segundoOperando float32 ) ( resultado float32 ) {
	return primerOperando + segundoOperando
}

func Restar( primerOperando float32, segundoOperando float32 ) ( resultado float32 ) {
	return primerOperando - segundoOperando
}

func Multiplicar( primerOperando float32, segundoOperando float32 ) ( resultado float32 ) {
	return primerOperando * segundoOperando
}

func Dividir( primerOperando float32, segundoOperando float32 ) ( resultado float32) {
	if segundoOperando == 0 {
		return -1
	}
	return primerOperando / segundoOperando
}
