package main

import (
	"OLC2/TS"
	"OLC2/interfaces"
	"OLC2/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	tablaVacia := TS.TablaSimbolo{}
	interfaces.Ast.SetInstrucciones(result)
	TSGlobal := TS.NuevaTabla(tablaVacia)
	interfaces.Ast.SetTSglobal(TSGlobal)

	for _, s := range interfaces.Ast.GetInstrucciones().ToArray() {
		s.(interfaces.Instruccion).Interpretar(interfaces.Ast, TSGlobal)
	}

	fmt.Println(interfaces.Ast.GetConsola())

}

func main() {

	// Setup the input
	//is := antlr.NewInputStream("\"El resultado es: \" + (10+5+5+5+5+5)")

	is, _ := antlr.NewFileStream("entrada.txt")
	// Create the Lexer
	lexer := parser.NewChemsLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

}
