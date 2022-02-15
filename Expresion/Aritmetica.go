package Expresion

import (
	"OLC2/TS"
	"OLC2/interfaces"
)

type Aritmetica struct {
	operador     TS.OperadorAritmetico
	operacionIzq interfaces.Expresion
	operacionDer interfaces.Expresion
	fila         int
	columna      int
	tipo         interface{}
}

func NuevaAritmetica(operador TS.OperadorAritmetico, operacionIzq interfaces.Expresion, operacionDer interfaces.Expresion, fila int, columna int) Aritmetica {

	arit := Aritmetica{operador, operacionIzq, operacionDer, fila, columna, nil}
	return arit
}

func (p Aritmetica) Interpretar(tree TS.Arbol, table TS.TablaSimbolo) interfaces.Symbol {
	var der interfaces.Symbol
	var izq interfaces.Symbol
	var retorno interfaces.Symbol

	izq = p.operacionIzq.Interpretar(tree, table)
	if p.operacionDer != nil {
		der = p.operacionDer.Interpretar(tree, table)
	}

	if p.operador == TS.MAS {
		if izq.Tipo == TS.ENTERO && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.ENTERO
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(int) + der.Valor.(int)
			retorno.Valor = resultado
		}
	}

	return retorno

}

func (p Aritmetica) ObtenerValor(tipo TS.TIPO, valor interface{}) interface{} {
	if tipo == TS.ENTERO {
		return valor.(int)
	} else if tipo == TS.DECIMAL {
		return valor.(float64)
	} else if tipo == TS.BOOLEANO {
		return valor.(bool)
	} else {
		return valor.(string)
	}
}
