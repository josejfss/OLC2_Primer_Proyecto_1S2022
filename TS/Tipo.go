package TS

type TIPO int
type OperadorAritmetico int
type OperadorRelacional int
type OperadorLogico int

const (
	ENTERO TIPO = iota
	DECIMAL
	BOOLEANO
	CHARACTER
	CADENA
	NULO
	ARREGLO
)

const (
	MAS OperadorAritmetico = iota
	MENOS
	POR
	DIV
	POT
	MOD
	UMENOS
)

const (
	MENORQUE OperadorRelacional = iota
	MAYORQUE
	MENORIGUAL
	MAYORIGUAL
	IGUALIGUAL
	DIFERENTE
)

const (
	NOT OperadorLogico = iota
	AND
	OR
)
