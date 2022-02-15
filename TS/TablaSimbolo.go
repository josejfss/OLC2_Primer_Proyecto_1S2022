package TS

import (
	"fmt"
	"reflect"
	"strings"
)

type TablaSimbolo struct {
	tabla    map[string]Simbolo
	anterior *TablaSimbolo
}

func NuevaTabla(anterior TablaSimbolo) TablaSimbolo {
	var tabla map[string]Simbolo
	newTable := TablaSimbolo{tabla, &anterior}
	return newTable
}

func (p *TablaSimbolo) SetTabla(simbolo Simbolo) interface{} { //agrega una variable

	encontrado := false

	for key, element := range p.tabla {
		if key == simbolo.identificador {
			encontrado = true
			fmt.Println(element.valor)
		}

	}
	if encontrado {
		exe := NuevaExcepcion("Semantico", "La variable: "+simbolo.identificador+" ya existe", simbolo.fila, simbolo.columna)
		return exe
	} else {
		p.tabla[strings.ToLower(simbolo.identificador)] = simbolo
		return nil
	}
}

func (p TablaSimbolo) GetTabla(id string) interface{} {
	tablaActual := p

	for reflect.DeepEqual(tablaActual, nil) {
		encontrado := false
		for key, element := range tablaActual.tabla {
			if key == strings.ToLower(id) {
				encontrado = true
				fmt.Println(element.valor)
			}
		}

		if encontrado {
			return tablaActual.tabla[strings.ToLower(id)] //retorna simbolo
		} else {
			tablaActual = *tablaActual.anterior
		}
	}
	return nil
}

func (p TablaSimbolo) ActualizarTabla(simbolo Simbolo) interface{} {
	tablaActual := p

	for reflect.DeepEqual(tablaActual, nil) {
		encontrado := false
		for key, element := range tablaActual.tabla {
			if key == strings.ToLower(simbolo.identificador) {
				encontrado = true
				fmt.Println(element.valor)
			}
		}

		if encontrado {
			if tablaActual.tabla[strings.ToLower(simbolo.identificador)].tipo == simbolo.tipo {
				tablaActual.tabla[strings.ToLower(simbolo.identificador)] = simbolo
				return nil
			} else {
				exe := NuevaExcepcion("Semantico", "Tipo de dato Diferente en Asignacion", simbolo.fila, simbolo.columna)
				return exe
			}
		} else {
			tablaActual = *tablaActual.anterior
		}
	}
	exe := NuevaExcepcion("Semantico", "Variable No encontrada en Asignacion", simbolo.fila, simbolo.columna)
	return exe
}
