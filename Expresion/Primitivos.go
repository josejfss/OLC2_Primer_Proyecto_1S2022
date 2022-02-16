package Expresion

import (
	"OLC2/TS"
	"OLC2/interfaces"
)

type Primitivo struct {
	valor   interface{}
	tipo    TS.TIPO
	fila    int
	columna int
}

func (p Primitivo) Interpretar(tree TS.Arbol, table TS.TablaSimbolo) interfaces.Symbol {
	return interfaces.Symbol{
		Tipo:    p.tipo,
		Valor:   p.valor,
		Fila:    p.fila,
		Columna: p.columna,
	}
}

func NuevoPrimitivo(val interface{}, tipo TS.TIPO, fila int, columna int) Primitivo {
	primi := Primitivo{val, tipo, fila, columna}
	return primi
}