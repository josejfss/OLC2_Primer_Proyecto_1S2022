package Instrucciones

import (
	"OLC2/TS"
	"OLC2/interfaces"
)

type Imprimir struct {
	expresion interfaces.Expresion
	fila      int
	columna   int
}

func NuevoImprimir(valor interfaces.Expresion, fila int, columna int) Imprimir {
	imp := Imprimir{valor, fila, columna}
	return imp
}

func (p Imprimir) Interpretar(tree TS.Arbol, table TS.TablaSimbolo) interface{} {

	value := p.expresion.Interpretar(interfaces.Ast, table)

	if value.Tipo == TS.ERROR {
		return value.Valor
	}

	interfaces.Ast.UpdateConsola(value.Valor)
	return nil
}
