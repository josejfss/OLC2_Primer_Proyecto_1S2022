// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/interfaces"
import "OLC2/Expresion"
import "OLC2/TS"
import "OLC2/Instrucciones"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 74, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 7, 3, 19, 10, 3, 12, 3, 14, 3, 22, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 47, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 59,
	10, 6, 12, 6, 14, 6, 62, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 72, 10, 7, 3, 7, 2, 3, 10, 8, 2, 4, 6, 8, 10, 12, 2, 4,
	3, 2, 30, 32, 3, 2, 33, 34, 2, 74, 2, 14, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2,
	6, 25, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 71, 3, 2,
	2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 8, 2, 1, 2, 16, 3, 3, 2, 2, 2, 17, 19,
	5, 6, 4, 2, 18, 17, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2,
	20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 8,
	3, 1, 2, 24, 5, 3, 2, 2, 2, 25, 26, 7, 3, 2, 2, 26, 27, 7, 21, 2, 2, 27,
	28, 7, 4, 2, 2, 28, 29, 7, 35, 2, 2, 29, 30, 5, 8, 5, 2, 30, 31, 7, 36,
	2, 2, 31, 32, 7, 22, 2, 2, 32, 33, 8, 4, 1, 2, 33, 7, 3, 2, 2, 2, 34, 35,
	5, 10, 6, 2, 35, 36, 8, 5, 1, 2, 36, 9, 3, 2, 2, 2, 37, 38, 8, 6, 1, 2,
	38, 39, 5, 12, 7, 2, 39, 40, 8, 6, 1, 2, 40, 47, 3, 2, 2, 2, 41, 42, 7,
	35, 2, 2, 42, 43, 5, 8, 5, 2, 43, 44, 7, 36, 2, 2, 44, 45, 8, 6, 1, 2,
	45, 47, 3, 2, 2, 2, 46, 37, 3, 2, 2, 2, 46, 41, 3, 2, 2, 2, 47, 60, 3,
	2, 2, 2, 48, 49, 12, 6, 2, 2, 49, 50, 9, 2, 2, 2, 50, 51, 5, 10, 6, 7,
	51, 52, 8, 6, 1, 2, 52, 59, 3, 2, 2, 2, 53, 54, 12, 5, 2, 2, 54, 55, 9,
	3, 2, 2, 55, 56, 5, 10, 6, 6, 56, 57, 8, 6, 1, 2, 57, 59, 3, 2, 2, 2, 58,
	48, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 11, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64,
	7, 13, 2, 2, 64, 72, 8, 7, 1, 2, 65, 66, 7, 16, 2, 2, 66, 72, 8, 7, 1,
	2, 67, 68, 7, 14, 2, 2, 68, 72, 8, 7, 1, 2, 69, 70, 7, 19, 2, 2, 70, 72,
	8, 7, 1, 2, 71, 63, 3, 2, 2, 2, 71, 65, 3, 2, 2, 2, 71, 67, 3, 2, 2, 2,
	71, 69, 3, 2, 2, 2, 72, 13, 3, 2, 2, 2, 7, 20, 46, 58, 60, 71,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'sentencia'", "'consola'", "'publico'", "'main'", "'integer'", "'string'",
	"'real'", "'boolean'", "'if'", "'entonces'", "", "", "", "", "", "", "",
	"','", "'.'", "';'", "'||'", "'&&'", "'!'", "'>='", "'<='", "'>'", "'<'",
	"'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "RSENTENCIA", "RCONSOLA", "RPUBLICO", "RMAIN", "RINTEGER", "RSTRING",
	"RREAL", "RBOOLEAN", "RIF", "RENTONCES", "ENTERO", "DECIMAL", "ID", "CADENA",
	"COMMENT", "LINE_COMMENT", "BOOLEANO", "COMA", "PUNTO", "PUNTOCOMA", "OR",
	"AND", "NOT", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", "POR",
	"DIV", "MOD", "SUMA", "RESTA", "PARA", "PARC", "LLAVEA", "LLAVEC", "CORA",
	"CORC", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "expression", "expr_arit", "primitivo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF          = antlr.TokenEOF
	ChemsRSENTENCIA   = 1
	ChemsRCONSOLA     = 2
	ChemsRPUBLICO     = 3
	ChemsRMAIN        = 4
	ChemsRINTEGER     = 5
	ChemsRSTRING      = 6
	ChemsRREAL        = 7
	ChemsRBOOLEAN     = 8
	ChemsRIF          = 9
	ChemsRENTONCES    = 10
	ChemsENTERO       = 11
	ChemsDECIMAL      = 12
	ChemsID           = 13
	ChemsCADENA       = 14
	ChemsCOMMENT      = 15
	ChemsLINE_COMMENT = 16
	ChemsBOOLEANO     = 17
	ChemsCOMA         = 18
	ChemsPUNTO        = 19
	ChemsPUNTOCOMA    = 20
	ChemsOR           = 21
	ChemsAND          = 22
	ChemsNOT          = 23
	ChemsMAYORIGUAL   = 24
	ChemsMENORIGUAL   = 25
	ChemsMAYORQUE     = 26
	ChemsMENORQUE     = 27
	ChemsPOR          = 28
	ChemsDIV          = 29
	ChemsMOD          = 30
	ChemsSUMA         = 31
	ChemsRESTA        = 32
	ChemsPARA         = 33
	ChemsPARC         = 34
	ChemsLLAVEA       = 35
	ChemsLLAVEC       = 36
	ChemsCORA         = 37
	ChemsCORC         = 38
	ChemsWHITESPACE   = 39
)

// Chems rules.
const (
	ChemsRULE_start         = 0
	ChemsRULE_instrucciones = 1
	ChemsRULE_instruccion   = 2
	ChemsRULE_expression    = 3
	ChemsRULE_expr_arit     = 4
	ChemsRULE_primitivo     = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChemsRSENTENCIA {
		{
			p.SetState(15)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RSENTENCIA returns the _RSENTENCIA token.
	Get_RSENTENCIA() antlr.Token

	// Set_RSENTENCIA sets the _RSENTENCIA token.
	Set_RSENTENCIA(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruccion
	_RSENTENCIA antlr.Token
	_expression IExpressionContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_RSENTENCIA() antlr.Token { return s._RSENTENCIA }

func (s *InstruccionContext) Set_RSENTENCIA(v antlr.Token) { s._RSENTENCIA = v }

func (s *InstruccionContext) Get_expression() IExpressionContext { return s._expression }

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruccion { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruccion) { s.instr = v }

func (s *InstruccionContext) RSENTENCIA() antlr.TerminalNode {
	return s.GetToken(ChemsRSENTENCIA, 0)
}

func (s *InstruccionContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsPUNTO, 0)
}

func (s *InstruccionContext) RCONSOLA() antlr.TerminalNode {
	return s.GetToken(ChemsRCONSOLA, 0)
}

func (s *InstruccionContext) PARA() antlr.TerminalNode {
	return s.GetToken(ChemsPARA, 0)
}

func (s *InstruccionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstruccionContext) PARC() antlr.TerminalNode {
	return s.GetToken(ChemsPARC, 0)
}

func (s *InstruccionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPUNTOCOMA, 0)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)

		var _m = p.Match(ChemsRSENTENCIA)

		localctx.(*InstruccionContext)._RSENTENCIA = _m
	}
	{
		p.SetState(24)
		p.Match(ChemsPUNTO)
	}
	{
		p.SetState(25)
		p.Match(ChemsRCONSOLA)
	}
	{
		p.SetState(26)
		p.Match(ChemsPARA)
	}
	{
		p.SetState(27)

		var _x = p.Expression()

		localctx.(*InstruccionContext)._expression = _x
	}
	{
		p.SetState(28)
		p.Match(ChemsPARC)
	}
	{
		p.SetState(29)
		p.Match(ChemsPUNTOCOMA)
	}
	localctx.(*InstruccionContext).instr = Instrucciones.NuevoImprimir(localctx.(*InstruccionContext).Get_expression().GetP(), (func() int {
		if localctx.(*InstruccionContext).Get_RSENTENCIA() == nil {
			return 0
		} else {
			return localctx.(*InstruccionContext).Get_RSENTENCIA().GetLine()
		}
	}()), (func() int {
		if localctx.(*InstruccionContext).Get_RSENTENCIA() == nil {
			return 0
		} else {
			return localctx.(*InstruccionContext).Get_RSENTENCIA().GetColumn()
		}
	}()))

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expresion
	_expr_arit IExpr_aritContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) GetP() interfaces.Expresion { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Chems) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_aritContext

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_aritContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_aritContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expresion
	opIz        IExpr_aritContext
	_primitivo  IPrimitivoContext
	_expression IExpressionContext
	op          antlr.Token
	opDe        IExpr_aritContext
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpIz() IExpr_aritContext { return s.opIz }

func (s *Expr_aritContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) GetP() interfaces.Expresion { return s.p }

func (s *Expr_aritContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Expr_aritContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expr_aritContext) PARA() antlr.TerminalNode {
	return s.GetToken(ChemsPARA, 0)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) PARC() antlr.TerminalNode {
	return s.GetToken(ChemsPARC, 0)
}

func (s *Expr_aritContext) AllExpr_arit() []IExpr_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem())
	var tst = make([]IExpr_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_aritContext)
		}
	}

	return tst
}

func (s *Expr_aritContext) Expr_arit(i int) IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_aritContext) POR() antlr.TerminalNode {
	return s.GetToken(ChemsPOR, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(ChemsDIV, 0)
}

func (s *Expr_aritContext) MOD() antlr.TerminalNode {
	return s.GetToken(ChemsMOD, 0)
}

func (s *Expr_aritContext) SUMA() antlr.TerminalNode {
	return s.GetToken(ChemsSUMA, 0)
}

func (s *Expr_aritContext) RESTA() antlr.TerminalNode {
	return s.GetToken(ChemsRESTA, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *Chems) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *Chems) expr_arit(_p int) (localctx IExpr_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ChemsRULE_expr_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsENTERO, ChemsDECIMAL, ChemsCADENA, ChemsBOOLEANO:
		{
			p.SetState(36)

			var _x = p.Primitivo()

			localctx.(*Expr_aritContext)._primitivo = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitivo().GetP()

	case ChemsPARA:
		{
			p.SetState(39)
			p.Match(ChemsPARA)
		}
		{
			p.SetState(40)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(41)
			p.Match(ChemsPARC)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(47)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsPOR)|(1<<ChemsDIV)|(1<<ChemsMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opDe = _x
				}

				if (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()) == "*" {
					localctx.(*Expr_aritContext).p = Expresion.NuevaAritmetica(TS.POR, localctx.(*Expr_aritContext).GetOpIz().GetP(), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetLine()
						}
					}()), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetColumn()
						}
					}()))
				} else if (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()) == "/" {
					localctx.(*Expr_aritContext).p = Expresion.NuevaAritmetica(TS.DIV, localctx.(*Expr_aritContext).GetOpIz().GetP(), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetLine()
						}
					}()), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetColumn()
						}
					}()))
				} else if (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()) == "%" {
					localctx.(*Expr_aritContext).p = Expresion.NuevaAritmetica(TS.MOD, localctx.(*Expr_aritContext).GetOpIz().GetP(), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetLine()
						}
					}()), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetColumn()
						}
					}()))
				}

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(52)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsSUMA || _la == ChemsRESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)

					var _x = p.expr_arit(4)

					localctx.(*Expr_aritContext).opDe = _x
				}

				if (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()) == "+" {
					localctx.(*Expr_aritContext).p = Expresion.NuevaAritmetica(TS.MAS, localctx.(*Expr_aritContext).GetOpIz().GetP(), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetLine()
						}
					}()), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetColumn()
						}
					}()))
				} else if (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()) == "-" {
					localctx.(*Expr_aritContext).p = Expresion.NuevaAritmetica(TS.MENOS, localctx.(*Expr_aritContext).GetOpIz().GetP(), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetLine()
						}
					}()), (func() int {
						if localctx.(*Expr_aritContext).GetOp() == nil {
							return 0
						} else {
							return localctx.(*Expr_aritContext).GetOp().GetColumn()
						}
					}()))
				}

			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_BOOLEANO returns the _BOOLEANO token.
	Get_BOOLEANO() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_BOOLEANO sets the _BOOLEANO token.
	Set_BOOLEANO(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         interfaces.Expresion
	_ENTERO   antlr.Token
	_CADENA   antlr.Token
	_DECIMAL  antlr.Token
	_BOOLEANO antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *PrimitivoContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrimitivoContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *PrimitivoContext) Get_BOOLEANO() antlr.Token { return s._BOOLEANO }

func (s *PrimitivoContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *PrimitivoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivoContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *PrimitivoContext) Set_BOOLEANO(v antlr.Token) { s._BOOLEANO = v }

func (s *PrimitivoContext) GetP() interfaces.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *PrimitivoContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(ChemsENTERO, 0)
}

func (s *PrimitivoContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *PrimitivoContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ChemsDECIMAL, 0)
}

func (s *PrimitivoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEANO, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)

			var _m = p.Match(ChemsENTERO)

			localctx.(*PrimitivoContext)._ENTERO = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ENTERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = Expresion.NuevoPrimitivo(num, TS.ENTERO, (func() int {
			if localctx.(*PrimitivoContext).Get_ENTERO() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ENTERO().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_ENTERO() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ENTERO().GetColumn()
			}
		}()))

	case ChemsCADENA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)

			var _m = p.Match(ChemsCADENA)

			localctx.(*PrimitivoContext)._CADENA = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = Expresion.NuevoPrimitivo(str, TS.CADENA, (func() int {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetColumn()
			}
		}()))

	case ChemsDECIMAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)

			var _m = p.Match(ChemsDECIMAL)

			localctx.(*PrimitivoContext)._DECIMAL = _m
		}

		s, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = Expresion.NuevoPrimitivo(s, TS.DECIMAL, (func() int {
			if localctx.(*PrimitivoContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DECIMAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DECIMAL().GetColumn()
			}
		}()))

	case ChemsBOOLEANO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)

			var _m = p.Match(ChemsBOOLEANO)

			localctx.(*PrimitivoContext)._BOOLEANO = _m
		}

		s, err := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEANO() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEANO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = Expresion.NuevoPrimitivo(s, TS.BOOLEANO, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEANO() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEANO().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEANO() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEANO().GetColumn()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
