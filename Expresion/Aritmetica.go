package Expresion

import (
	"OLC2/TS"
	"OLC2/interfaces"
	"fmt"
	"strconv"
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
			return retorno
		} else if izq.Tipo == TS.ENTERO && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := float64(izq.Valor.(int)) + der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		}else if izq.Tipo == TS.ENTERO && der.Tipo == TS.CADENA {
			retorno.Tipo = TS.CADENA
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := strconv.Itoa(izq.Valor.(int)) + der.Valor.(string)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.CADENA && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.CADENA
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(string) + strconv.Itoa(der.Valor.(int))
			retorno.Valor = resultado
			return retorno
		}else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.CADENA {
			retorno.Tipo = TS.CADENA
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := fmt.Sprintf("%.2f", izq.Valor.(float64)) + der.Valor.(string)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.CADENA && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.CADENA
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(string) + fmt.Sprintf("%.2f",der.Valor.(float64))
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.CADENA && der.Tipo == TS.CADENA {
			retorno.Tipo = TS.CADENA
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(string) + der.Valor.(string)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) + float64(der.Valor.(int))
			retorno.Valor = resultado
			return retorno
		}else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) + der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		}

		//error en suma
		retorno.Tipo = TS.ERROR
		retorno.Valor = TS.NuevaExcepcion("Semantico", "Operacion de suma invalido", p.fila, p.columna)
		return retorno
	//RESTA
	}else if p.operador == TS.MENOS {
		if izq.Tipo == TS.ENTERO && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.ENTERO
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(int) - der.Valor.(int)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.ENTERO && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := float64(izq.Valor.(int)) - der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) - float64(der.Valor.(int))
			retorno.Valor = resultado
			return retorno
		}else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) - der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		}

		//error en suma
		retorno.Tipo = TS.ERROR
		retorno.Valor = TS.NuevaExcepcion("Semantico", "Operacion de resta invalido", p.fila, p.columna)
		return retorno
	//MULTIPLICACION
	}else if p.operador == TS.POR {
		if izq.Tipo == TS.ENTERO && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.ENTERO
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(int) * der.Valor.(int)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.ENTERO && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := float64(izq.Valor.(int)) * der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		} else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.ENTERO {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) * float64(der.Valor.(int))
			retorno.Valor = resultado
			return retorno
		}else if izq.Tipo == TS.DECIMAL && der.Tipo == TS.DECIMAL {
			retorno.Tipo = TS.DECIMAL
			retorno.Columna = p.columna
			retorno.Fila = p.fila
			resultado := izq.Valor.(float64) * der.Valor.(float64)
			retorno.Valor = resultado
			return retorno
		}

		//error en suma
		retorno.Tipo = TS.ERROR
		retorno.Valor = TS.NuevaExcepcion("Semantico", "Operacion de multiplicacion invalido", p.fila, p.columna)
		return retorno
	//MULTIPLICACION
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
