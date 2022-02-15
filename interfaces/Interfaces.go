package interfaces

import "OLC2/TS"

type Symbol struct {
	Tipo    TS.TIPO
	Valor   interface{}
	Fila    int
	Columna int
}

var Ast TS.Arbol = TS.NuevoArbol()

type Expresion interface {
	Interpretar(tree TS.Arbol, table TS.TablaSimbolo) Symbol
}

type Instruction interface {
	Ejecutar() interface{}
}

type Instruccion interface {
	Interpretar(tree TS.Arbol, table TS.TablaSimbolo) interface{}
}
