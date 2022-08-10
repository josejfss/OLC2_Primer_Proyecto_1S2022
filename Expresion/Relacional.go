package Expresion

import (
	"OLC2/TS"
	"OLC2/interfaces"
)

type Relacional struct {
	operador     TS.OperadorAritmetico
	operacionIzq interfaces.Expresion
	operacionDer interfaces.Expresion
	fila         int
	columna      int
	tipo         TS.TIPO
}

func NuevaRelacional(operador TS.OperadorAritmetico, operacionIzq interfaces.Expresion, operacionDer interfaces.Expresion, fila int, columna int) Relacional {

	arit := Relacional{operador, operacionIzq, operacionDer, fila, columna, TS.BOOLEANO}
	return arit
}

func (p Relacional) Interpretar(tree TS.Arbol, table TS.TablaSimbolo) interfaces.Symbol {
	var der interfaces.Symbol
	var izq interfaces.Symbol
	var retorno interfaces.Symbol

	izq = p.operacionIzq.Interpretar(tree, table)
	if izq.Tipo == TS.ERROR {
		return izq
	}
	if p.operacionDer != nil {
		der = p.operacionDer.Interpretar(tree, table)
		if der.Tipo == TS.ERROR {
			return der
		}
	}


	return retorno
}