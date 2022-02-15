lexer grammar ChemsLexer;

// Tokens

//RESERVADAS
RSENTENCIA:     'sentencia';
RCONSOLA:       'consola';
RPUBLICO:       'publico';
RMAIN:          'main';
RINTEGER:       'integer';
RSTRING:        'string';
RREAL:          'real';
RBOOLEAN:       'boolean';
RIF:            'if';
RENTONCES:      'entonces';


ENTERO:         [0-9]+;
DECIMAL:        [0-9]+'.'[0-9]+;  
ID:             [a-zA-Z][a-zA-Z_0-9]*;      
CADENA:         '"'~["]*'"';
COMMENT:        '(*' .*? '*)' -> skip;
LINE_COMMENT:   '//' ~[\r\n]* -> skip;

COMA:               ',';
PUNTO:              '.';
PUNTOCOMA:          ';';
OR:                 '||';
AND:                '&&';
NOT:                '!';
MAYORIGUAL:         '>=';
MENORIGUAL:         '<=';
MAYORQUE:           '>';
MENORQUE:           '<';
POR:                '*';
DIV:                '/';
SUMA:               '+';
RESTA:              '-';
PARA:               '(';
PARC:               ')';
LLAVEA:             '{';
LLAVEC:             '}';
CORA:               '[';
CORC:               ']';


WHITESPACE: [ \\\r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;


