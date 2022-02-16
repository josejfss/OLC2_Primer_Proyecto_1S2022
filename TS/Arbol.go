package TS

import (
	"fmt"

	"github.com/colegno/arraylist"
)

type Arbol struct {
	instrucciones *arraylist.List
	funciones     []interface{}
	excepciones   arraylist.List
	consola       string
	dot           string
	contador      int
	TSglobal      TablaSimbolo
}

func NuevoArbol() Arbol {
	var funciones []interface{}
	var excepciones arraylist.List
	var TSglobal TablaSimbolo
	var instrucciones *arraylist.List
	newTree := Arbol{instrucciones, funciones, excepciones, "", "", 0, TSglobal}
	return newTree
}

func (p *Arbol) SetInstrucciones(instrucciones *arraylist.List) {
	p.instrucciones = instrucciones
}

func (p Arbol) GetInstrucciones() *arraylist.List {
	return p.instrucciones
}

func (p *Arbol) SetExcepciones(excepciones arraylist.List) {
	p.excepciones = excepciones
}

func (p Arbol) GetExcepciones() interface{} {
	return p.excepciones
}

func (p Arbol) UpdateExcepciones(excepcion Excepcion){
	p.excepciones.Add(excepcion)
}

func (p *Arbol) SetConsola(consola string) {
	p.consola = consola
}

func (p Arbol) GetConsola() interface{} {
	return p.consola
}

func (p *Arbol) UpdateConsola(val interface{}) {
	valor := fmt.Sprintf(" %v \n", val)
	p.consola += valor
}

func (p *Arbol) SetTSglobal(TSglobal TablaSimbolo) {
	p.TSglobal = TSglobal
}

func (p Arbol) GetTSglobal() interface{} {
	return p.TSglobal
}

func (p Arbol) GetFunciones() interface{} {
	return p.funciones
}

/*
func (p Arbol) getFuncion(nombre string) interface{} {
	for _, elem := range p.funciones {
		if nombre == elem.nombre {
			return elem
		}
	}
	return nil
} */

func (p *Arbol) AddFuncion(funcion interface{}) {
	p.funciones = append(p.funciones, funcion)
}

func (p *Arbol) GetDot(instrucciones *arraylist.List) {
	p.instrucciones = instrucciones
}
