package TS

type Excepcion struct {
	tipo  string
	descripcion string
	fila int
	columna int
}

func NuevaExcepcion(tipo string, descripcion string, fila int, columna int) Excepcion{
	exce := Excepcion{tipo, descripcion, fila, columna}
	return exce
}
	