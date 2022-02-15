// Generated from c:\Users\GMG\Documents\GitHub\OLC2_Primer_Proyecto_1S2022\Chems.g4 by ANTLR 4.8

    import "OLC2/interfaces"
    import "OLC2/Expresion"
    import "OLC2/TS"
    import "OLC2/Instrucciones"
    import arrayList "github.com/colegno/arraylist"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Chems extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		RSENTENCIA=1, RCONSOLA=2, RPUBLICO=3, RMAIN=4, RINTEGER=5, RSTRING=6, 
		RREAL=7, RBOOLEAN=8, RIF=9, RENTONCES=10, ENTERO=11, DECIMAL=12, ID=13, 
		CADENA=14, COMMENT=15, LINE_COMMENT=16, COMA=17, PUNTO=18, PUNTOCOMA=19, 
		OR=20, AND=21, NOT=22, MAYORIGUAL=23, MENORIGUAL=24, MAYORQUE=25, MENORQUE=26, 
		POR=27, DIV=28, SUMA=29, RESTA=30, PARA=31, PARC=32, LLAVEA=33, LLAVEC=34, 
		CORA=35, CORC=36, WHITESPACE=37;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_expression = 3, 
		RULE_expr_arit = 4, RULE_primitivo = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "expression", "expr_arit", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'sentencia'", "'consola'", "'publico'", "'main'", "'integer'", 
			"'string'", "'real'", "'boolean'", "'if'", "'entonces'", null, null, 
			null, null, null, null, "','", "'.'", "';'", "'||'", "'&&'", "'!'", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RSENTENCIA", "RCONSOLA", "RPUBLICO", "RMAIN", "RINTEGER", "RSTRING", 
			"RREAL", "RBOOLEAN", "RIF", "RENTONCES", "ENTERO", "DECIMAL", "ID", "CADENA", 
			"COMMENT", "LINE_COMMENT", "COMA", "PUNTO", "PUNTOCOMA", "OR", "AND", 
			"NOT", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", "POR", "DIV", 
			"SUMA", "RESTA", "PARA", "PARC", "LLAVEA", "LLAVEC", "CORA", "CORC", 
			"WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Chems.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Chems(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			((StartContext)_localctx).instrucciones = instrucciones();
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List l;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RSENTENCIA) {
				{
				{
				setState(15);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(20);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			      listInt := localctx.(*InstruccionesContext).GetE()
			      		for _, e := range listInt {
			            _localctx.l.Add(e.GetInstr())
			          }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruccion instr;
		public Token RSENTENCIA;
		public ExpressionContext expression;
		public TerminalNode RSENTENCIA() { return getToken(Chems.RSENTENCIA, 0); }
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode RCONSOLA() { return getToken(Chems.RCONSOLA, 0); }
		public TerminalNode PARA() { return getToken(Chems.PARA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARC() { return getToken(Chems.PARC, 0); }
		public TerminalNode PUNTOCOMA() { return getToken(Chems.PUNTOCOMA, 0); }
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			((InstruccionContext)_localctx).RSENTENCIA = match(RSENTENCIA);
			setState(24);
			match(PUNTO);
			setState(25);
			match(RCONSOLA);
			setState(26);
			match(PARA);
			setState(27);
			((InstruccionContext)_localctx).expression = expression();
			setState(28);
			match(PARC);
			setState(29);
			match(PUNTOCOMA);
			_localctx.instr = Instrucciones.NuevoImprimir(((InstruccionContext)_localctx).expression.p, (((InstruccionContext)_localctx).RSENTENCIA!=null?((InstruccionContext)_localctx).RSENTENCIA.getLine():0), (((InstruccionContext)_localctx).RSENTENCIA!=null?((InstruccionContext)_localctx).RSENTENCIA.getCharPositionInLine():0) )
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext expr_arit;
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			((ExpressionContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((ExpressionContext)_localctx).expr_arit.p
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_aritContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext opIz;
		public PrimitivoContext primitivo;
		public ExpressionContext expression;
		public Token op;
		public Expr_aritContext opDe;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public TerminalNode PARA() { return getToken(Chems.PARA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARC() { return getToken(Chems.PARC, 0); }
		public List<Expr_aritContext> expr_arit() {
			return getRuleContexts(Expr_aritContext.class);
		}
		public Expr_aritContext expr_arit(int i) {
			return getRuleContext(Expr_aritContext.class,i);
		}
		public TerminalNode POR() { return getToken(Chems.POR, 0); }
		public TerminalNode DIV() { return getToken(Chems.DIV, 0); }
		public TerminalNode SUMA() { return getToken(Chems.SUMA, 0); }
		public TerminalNode RESTA() { return getToken(Chems.RESTA, 0); }
		public Expr_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_arit; }
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
			case DECIMAL:
			case CADENA:
				{
				setState(36);
				((Expr_aritContext)_localctx).primitivo = primitivo();
				_localctx.p = ((Expr_aritContext)_localctx).primitivo.p
				}
				break;
			case PARA:
				{
				setState(39);
				match(PARA);
				setState(40);
				((Expr_aritContext)_localctx).expression = expression();
				setState(41);
				match(PARC);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(58);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(56);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(46);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(47);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==POR || _la==DIV) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(48);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);

						                if (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null) == "*"{
						                  _localctx.p = Expresion.NuevaAritmetica(TS.POR, ((Expr_aritContext)_localctx).opIz.p, ((Expr_aritContext)_localctx).opDe.p, (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getLine():0), (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getCharPositionInLine():0))
						                } else if (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null) == "/"{
						                  _localctx.p = Expresion.NuevaAritmetica(TS.DIV, ((Expr_aritContext)_localctx).opIz.p, ((Expr_aritContext)_localctx).opDe.p, (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getLine():0), (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getCharPositionInLine():0))
						                }
						                
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(51);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(52);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==SUMA || _la==RESTA) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(53);
						((Expr_aritContext)_localctx).opDe = expr_arit(4);

						                if (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null) == "+"{
						                  _localctx.p = Expresion.NuevaAritmetica(TS.MAS, ((Expr_aritContext)_localctx).opIz.p, ((Expr_aritContext)_localctx).opDe.p, (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getLine():0), (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getCharPositionInLine():0))
						                } else if (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null) == "-"{
						                  _localctx.p = Expresion.NuevaAritmetica(TS.MENOS, ((Expr_aritContext)_localctx).opIz.p, ((Expr_aritContext)_localctx).opDe.p, (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getLine():0), (((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getCharPositionInLine():0))
						                }
						              
						}
						break;
					}
					} 
				}
				setState(60);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Token ENTERO;
		public Token CADENA;
		public Token DECIMAL;
		public TerminalNode ENTERO() { return getToken(Chems.ENTERO, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode DECIMAL() { return getToken(Chems.DECIMAL, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_primitivo);
		try {
			setState(67);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(61);
				((PrimitivoContext)_localctx).ENTERO = match(ENTERO);

				            	num,err := strconv.Atoi((((PrimitivoContext)_localctx).ENTERO!=null?((PrimitivoContext)_localctx).ENTERO.getText():null))
				                if err!= nil{
				                    fmt.Println(err)
				                }
				            _localctx.p = Expresion.NuevoPrimitivo(num, TS.ENTERO, (((PrimitivoContext)_localctx).ENTERO!=null?((PrimitivoContext)_localctx).ENTERO.getLine():0), (((PrimitivoContext)_localctx).ENTERO!=null?((PrimitivoContext)_localctx).ENTERO.getCharPositionInLine():0))
				       
				}
				break;
			case CADENA:
				enterOuterAlt(_localctx, 2);
				{
				setState(63);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);
				 
				      str:= (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1:len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null))-1]
				      _localctx.p = Expresion.NuevoPrimitivo(str,TS.CADENA, (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getLine():0), (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getCharPositionInLine():0))
				      
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(65);
				((PrimitivoContext)_localctx).DECIMAL = match(DECIMAL);

				      s, err := strconv.ParseFloat((((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getText():null), 64); 
				          if err == nil {
				          fmt.Println(err)
				        }
				      _localctx.p = Expresion.NuevoPrimitivo(s, TS.DECIMAL, (((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getLine():0), (((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getCharPositionInLine():0))
				    
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\'H\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\3\2\3\3\7\3\23\n\3\f\3\16\3"+
		"\26\13\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6/\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\7\6;\n\6\f\6\16\6>\13\6\3\7\3\7\3\7\3\7\3\7\3\7\5\7F\n\7\3"+
		"\7\2\3\n\b\2\4\6\b\n\f\2\4\3\2\35\36\3\2\37 \2G\2\16\3\2\2\2\4\24\3\2"+
		"\2\2\6\31\3\2\2\2\b\"\3\2\2\2\n.\3\2\2\2\fE\3\2\2\2\16\17\5\4\3\2\17\20"+
		"\b\2\1\2\20\3\3\2\2\2\21\23\5\6\4\2\22\21\3\2\2\2\23\26\3\2\2\2\24\22"+
		"\3\2\2\2\24\25\3\2\2\2\25\27\3\2\2\2\26\24\3\2\2\2\27\30\b\3\1\2\30\5"+
		"\3\2\2\2\31\32\7\3\2\2\32\33\7\24\2\2\33\34\7\4\2\2\34\35\7!\2\2\35\36"+
		"\5\b\5\2\36\37\7\"\2\2\37 \7\25\2\2 !\b\4\1\2!\7\3\2\2\2\"#\5\n\6\2#$"+
		"\b\5\1\2$\t\3\2\2\2%&\b\6\1\2&\'\5\f\7\2\'(\b\6\1\2(/\3\2\2\2)*\7!\2\2"+
		"*+\5\b\5\2+,\7\"\2\2,-\b\6\1\2-/\3\2\2\2.%\3\2\2\2.)\3\2\2\2/<\3\2\2\2"+
		"\60\61\f\6\2\2\61\62\t\2\2\2\62\63\5\n\6\7\63\64\b\6\1\2\64;\3\2\2\2\65"+
		"\66\f\5\2\2\66\67\t\3\2\2\678\5\n\6\689\b\6\1\29;\3\2\2\2:\60\3\2\2\2"+
		":\65\3\2\2\2;>\3\2\2\2<:\3\2\2\2<=\3\2\2\2=\13\3\2\2\2><\3\2\2\2?@\7\r"+
		"\2\2@F\b\7\1\2AB\7\20\2\2BF\b\7\1\2CD\7\16\2\2DF\b\7\1\2E?\3\2\2\2EA\3"+
		"\2\2\2EC\3\2\2\2F\r\3\2\2\2\7\24.:<E";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}