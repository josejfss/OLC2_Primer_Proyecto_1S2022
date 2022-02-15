package TS

type Simbolo struct {
	identificador string
	tipo          TIPO
	fila          int
	columna       int
	valor         interface{}
}

func NewSimbol(identificador string, tipo TIPO, fila int, columna int, valor interface{}) Simbolo {
	simbo := Simbolo{identificador, tipo, fila, columna, valor}
	return simbo
}

func (p Simbolo) GetID() string {
	return p.identificador
}

func (p *Simbolo) SetID(identificador string) {
	p.identificador = identificador
}

func (p Simbolo) GetTipo() TIPO {
	return p.tipo
}

func (p *Simbolo) SetTipo(tipo TIPO) {
	p.tipo = tipo
}

func (p Simbolo) GetValor() interface{} {
	return p.valor
}

func (p *Simbolo) SetValor(valor interface{}) {
	p.valor = valor
}

func (p Simbolo) GetColumna() int {
	return p.columna
}

func (p *Simbolo) SetColumna(columna int) {
	p.columna = columna
}
func (p Simbolo) GetFila() int {
	return p.fila
}

func (p *Simbolo) SetFila(fila int) {
	p.fila = fila
}
