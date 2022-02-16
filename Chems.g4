parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}

@header {
    import "OLC2/interfaces"
    import "OLC2/Expresion"
    import "OLC2/TS"
    import "OLC2/Instrucciones"
    import arrayList "github.com/colegno/arraylist"
}

start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instruccion*  {
      listInt := localctx.(*InstruccionesContext).GetE()
      		for _, e := range listInt {
            $l.Add(e.GetInstr())
          }
    }
;


instruccion returns [interfaces.Instruccion instr]
  : RSENTENCIA PUNTO RCONSOLA PARA expression PARC PUNTOCOMA {$instr = Instrucciones.NuevoImprimir($expression.p, $RSENTENCIA.line, $RSENTENCIA.pos )}
;

expression returns[interfaces.Expresion p]
    : expr_arit    {$p = $expr_arit.p}
;

expr_arit returns[interfaces.Expresion p]
    : opIz = expr_arit op=( POR| DIV| MOD) opDe = expr_arit {
      if $op.text == "*"{
        $p = Expresion.NuevaAritmetica(TS.POR, $opIz.p, $opDe.p, $op.line, $op.pos)
      } else if $op.text == "/"{
        $p = Expresion.NuevaAritmetica(TS.DIV, $opIz.p, $opDe.p, $op.line, $op.pos)
      } else if $op.text == "%"{
        $p = Expresion.NuevaAritmetica(TS.MOD, $opIz.p, $opDe.p, $op.line, $op.pos)
      }
      }
    | opIz = expr_arit op=(SUMA|RESTA) opDe = expr_arit {
      if $op.text == "+"{
        $p = Expresion.NuevaAritmetica(TS.MAS, $opIz.p, $opDe.p, $op.line, $op.pos)
      } else if $op.text == "-"{
        $p = Expresion.NuevaAritmetica(TS.MENOS, $opIz.p, $opDe.p, $op.line, $op.pos)
      }
    }     
    | primitivo {$p = $primitivo.p} 
    | PARA expression PARC {$p = $expression.p}
;

primitivo returns[interfaces.Expresion p]
    :ENTERO {
            	num,err := strconv.Atoi($ENTERO.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = Expresion.NuevoPrimitivo(num, TS.ENTERO, $ENTERO.line, $ENTERO.pos)
       } 
    | CADENA { 
      str:= $CADENA.text[1:len($CADENA.text)-1]
      $p = Expresion.NuevoPrimitivo(str,TS.CADENA, $CADENA.line, $CADENA.pos)
      }
    | DECIMAL {
      s, err := strconv.ParseFloat($DECIMAL.text, 64); 
          if err == nil {
          fmt.Println(err)
        }
      $p = Expresion.NuevoPrimitivo(s, TS.DECIMAL, $DECIMAL.line, $DECIMAL.pos)
    }
    |BOOLEANO {
      s, err := strconv.ParseBool($BOOLEANO.text); 
          if err == nil {
          fmt.Println(err)
        }
      $p = Expresion.NuevoPrimitivo(s, TS.BOOLEANO, $BOOLEANO.line, $BOOLEANO.pos)
    }
;