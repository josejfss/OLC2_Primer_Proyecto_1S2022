// Generated from c:\Users\GMG\Documents\GitHub\OLC2_Primer_Proyecto_1S2022\ChemsLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ChemsLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		RSENTENCIA=1, RCONSOLA=2, RPUBLICO=3, RMAIN=4, RINTEGER=5, RSTRING=6, 
		RREAL=7, RBOOLEAN=8, RIF=9, RENTONCES=10, RBOOLEANO=11, ENTERO=12, DECIMAL=13, 
		ID=14, CADENA=15, COMMENT=16, LINE_COMMENT=17, COMA=18, PUNTO=19, PUNTOCOMA=20, 
		OR=21, AND=22, NOT=23, MAYORIGUAL=24, MENORIGUAL=25, MAYORQUE=26, MENORQUE=27, 
		POR=28, DIV=29, MOD=30, SUMA=31, RESTA=32, PARA=33, PARC=34, LLAVEA=35, 
		LLAVEC=36, CORA=37, CORC=38, WHITESPACE=39;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"RSENTENCIA", "RCONSOLA", "RPUBLICO", "RMAIN", "RINTEGER", "RSTRING", 
			"RREAL", "RBOOLEAN", "RIF", "RENTONCES", "RBOOLEANO", "ENTERO", "DECIMAL", 
			"ID", "CADENA", "COMMENT", "LINE_COMMENT", "COMA", "PUNTO", "PUNTOCOMA", 
			"OR", "AND", "NOT", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", 
			"POR", "DIV", "MOD", "SUMA", "RESTA", "PARA", "PARC", "LLAVEA", "LLAVEC", 
			"CORA", "CORC", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'sentencia'", "'consola'", "'publico'", "'main'", "'integer'", 
			"'string'", "'real'", "'boolean'", "'if'", "'entonces'", null, null, 
			null, null, null, null, null, "','", "'.'", "';'", "'||'", "'&&'", "'!'", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RSENTENCIA", "RCONSOLA", "RPUBLICO", "RMAIN", "RINTEGER", "RSTRING", 
			"RREAL", "RBOOLEAN", "RIF", "RENTONCES", "RBOOLEANO", "ENTERO", "DECIMAL", 
			"ID", "CADENA", "COMMENT", "LINE_COMMENT", "COMA", "PUNTO", "PUNTOCOMA", 
			"OR", "AND", "NOT", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", 
			"POR", "DIV", "MOD", "SUMA", "RESTA", "PARA", "PARC", "LLAVEA", "LLAVEC", 
			"CORA", "CORC", "WHITESPACE"
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


	public ChemsLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ChemsLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2)\u0116\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00a4\n\f\3\r\6\r\u00a7\n\r\r"+
		"\r\16\r\u00a8\3\16\6\16\u00ac\n\16\r\16\16\16\u00ad\3\16\3\16\6\16\u00b2"+
		"\n\16\r\16\16\16\u00b3\3\17\3\17\7\17\u00b8\n\17\f\17\16\17\u00bb\13\17"+
		"\3\20\3\20\7\20\u00bf\n\20\f\20\16\20\u00c2\13\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\7\21\u00ca\n\21\f\21\16\21\u00cd\13\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\7\22\u00d8\n\22\f\22\16\22\u00db\13\22\3\22"+
		"\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\30"+
		"\3\30\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36"+
		"\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3"+
		"(\6(\u010e\n(\r(\16(\u010f\3(\3(\3)\3)\3)\3\u00cb\2*\3\3\5\4\7\5\t\6\13"+
		"\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'"+
		"\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'"+
		"M(O)Q\2\3\2\t\3\2\62;\4\2C\\c|\6\2\62;C\\aac|\3\2$$\4\2\f\f\17\17\6\2"+
		"\13\f\17\17\"\"^^\t\2\"#%%--/\60<<BB]_\2\u011d\2\3\3\2\2\2\2\5\3\2\2\2"+
		"\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3"+
		"\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2"+
		"\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2"+
		"\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2"+
		"\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2"+
		"\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2"+
		"\2M\3\2\2\2\2O\3\2\2\2\3S\3\2\2\2\5]\3\2\2\2\7e\3\2\2\2\tm\3\2\2\2\13"+
		"r\3\2\2\2\rz\3\2\2\2\17\u0081\3\2\2\2\21\u0086\3\2\2\2\23\u008e\3\2\2"+
		"\2\25\u0091\3\2\2\2\27\u00a3\3\2\2\2\31\u00a6\3\2\2\2\33\u00ab\3\2\2\2"+
		"\35\u00b5\3\2\2\2\37\u00bc\3\2\2\2!\u00c5\3\2\2\2#\u00d3\3\2\2\2%\u00de"+
		"\3\2\2\2\'\u00e0\3\2\2\2)\u00e2\3\2\2\2+\u00e4\3\2\2\2-\u00e7\3\2\2\2"+
		"/\u00ea\3\2\2\2\61\u00ec\3\2\2\2\63\u00ef\3\2\2\2\65\u00f2\3\2\2\2\67"+
		"\u00f4\3\2\2\29\u00f6\3\2\2\2;\u00f8\3\2\2\2=\u00fa\3\2\2\2?\u00fc\3\2"+
		"\2\2A\u00fe\3\2\2\2C\u0100\3\2\2\2E\u0102\3\2\2\2G\u0104\3\2\2\2I\u0106"+
		"\3\2\2\2K\u0108\3\2\2\2M\u010a\3\2\2\2O\u010d\3\2\2\2Q\u0113\3\2\2\2S"+
		"T\7u\2\2TU\7g\2\2UV\7p\2\2VW\7v\2\2WX\7g\2\2XY\7p\2\2YZ\7e\2\2Z[\7k\2"+
		"\2[\\\7c\2\2\\\4\3\2\2\2]^\7e\2\2^_\7q\2\2_`\7p\2\2`a\7u\2\2ab\7q\2\2"+
		"bc\7n\2\2cd\7c\2\2d\6\3\2\2\2ef\7r\2\2fg\7w\2\2gh\7d\2\2hi\7n\2\2ij\7"+
		"k\2\2jk\7e\2\2kl\7q\2\2l\b\3\2\2\2mn\7o\2\2no\7c\2\2op\7k\2\2pq\7p\2\2"+
		"q\n\3\2\2\2rs\7k\2\2st\7p\2\2tu\7v\2\2uv\7g\2\2vw\7i\2\2wx\7g\2\2xy\7"+
		"t\2\2y\f\3\2\2\2z{\7u\2\2{|\7v\2\2|}\7t\2\2}~\7k\2\2~\177\7p\2\2\177\u0080"+
		"\7i\2\2\u0080\16\3\2\2\2\u0081\u0082\7t\2\2\u0082\u0083\7g\2\2\u0083\u0084"+
		"\7c\2\2\u0084\u0085\7n\2\2\u0085\20\3\2\2\2\u0086\u0087\7d\2\2\u0087\u0088"+
		"\7q\2\2\u0088\u0089\7q\2\2\u0089\u008a\7n\2\2\u008a\u008b\7g\2\2\u008b"+
		"\u008c\7c\2\2\u008c\u008d\7p\2\2\u008d\22\3\2\2\2\u008e\u008f\7k\2\2\u008f"+
		"\u0090\7h\2\2\u0090\24\3\2\2\2\u0091\u0092\7g\2\2\u0092\u0093\7p\2\2\u0093"+
		"\u0094\7v\2\2\u0094\u0095\7q\2\2\u0095\u0096\7p\2\2\u0096\u0097\7e\2\2"+
		"\u0097\u0098\7g\2\2\u0098\u0099\7u\2\2\u0099\26\3\2\2\2\u009a\u009b\7"+
		"v\2\2\u009b\u009c\7t\2\2\u009c\u009d\7w\2\2\u009d\u00a4\7g\2\2\u009e\u009f"+
		"\7h\2\2\u009f\u00a0\7c\2\2\u00a0\u00a1\7n\2\2\u00a1\u00a2\7u\2\2\u00a2"+
		"\u00a4\7g\2\2\u00a3\u009a\3\2\2\2\u00a3\u009e\3\2\2\2\u00a4\30\3\2\2\2"+
		"\u00a5\u00a7\t\2\2\2\u00a6\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00a6"+
		"\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\32\3\2\2\2\u00aa\u00ac\t\2\2\2\u00ab"+
		"\u00aa\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ab\3\2\2\2\u00ad\u00ae\3\2"+
		"\2\2\u00ae\u00af\3\2\2\2\u00af\u00b1\7\60\2\2\u00b0\u00b2\t\2\2\2\u00b1"+
		"\u00b0\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b3\u00b4\3\2"+
		"\2\2\u00b4\34\3\2\2\2\u00b5\u00b9\t\3\2\2\u00b6\u00b8\t\4\2\2\u00b7\u00b6"+
		"\3\2\2\2\u00b8\u00bb\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba"+
		"\36\3\2\2\2\u00bb\u00b9\3\2\2\2\u00bc\u00c0\7$\2\2\u00bd\u00bf\n\5\2\2"+
		"\u00be\u00bd\3\2\2\2\u00bf\u00c2\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0\u00c1"+
		"\3\2\2\2\u00c1\u00c3\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c3\u00c4\7$\2\2\u00c4"+
		" \3\2\2\2\u00c5\u00c6\7*\2\2\u00c6\u00c7\7,\2\2\u00c7\u00cb\3\2\2\2\u00c8"+
		"\u00ca\13\2\2\2\u00c9\u00c8\3\2\2\2\u00ca\u00cd\3\2\2\2\u00cb\u00cc\3"+
		"\2\2\2\u00cb\u00c9\3\2\2\2\u00cc\u00ce\3\2\2\2\u00cd\u00cb\3\2\2\2\u00ce"+
		"\u00cf\7,\2\2\u00cf\u00d0\7+\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d2\b\21"+
		"\2\2\u00d2\"\3\2\2\2\u00d3\u00d4\7\61\2\2\u00d4\u00d5\7\61\2\2\u00d5\u00d9"+
		"\3\2\2\2\u00d6\u00d8\n\6\2\2\u00d7\u00d6\3\2\2\2\u00d8\u00db\3\2\2\2\u00d9"+
		"\u00d7\3\2\2\2\u00d9\u00da\3\2\2\2\u00da\u00dc\3\2\2\2\u00db\u00d9\3\2"+
		"\2\2\u00dc\u00dd\b\22\2\2\u00dd$\3\2\2\2\u00de\u00df\7.\2\2\u00df&\3\2"+
		"\2\2\u00e0\u00e1\7\60\2\2\u00e1(\3\2\2\2\u00e2\u00e3\7=\2\2\u00e3*\3\2"+
		"\2\2\u00e4\u00e5\7~\2\2\u00e5\u00e6\7~\2\2\u00e6,\3\2\2\2\u00e7\u00e8"+
		"\7(\2\2\u00e8\u00e9\7(\2\2\u00e9.\3\2\2\2\u00ea\u00eb\7#\2\2\u00eb\60"+
		"\3\2\2\2\u00ec\u00ed\7@\2\2\u00ed\u00ee\7?\2\2\u00ee\62\3\2\2\2\u00ef"+
		"\u00f0\7>\2\2\u00f0\u00f1\7?\2\2\u00f1\64\3\2\2\2\u00f2\u00f3\7@\2\2\u00f3"+
		"\66\3\2\2\2\u00f4\u00f5\7>\2\2\u00f58\3\2\2\2\u00f6\u00f7\7,\2\2\u00f7"+
		":\3\2\2\2\u00f8\u00f9\7\61\2\2\u00f9<\3\2\2\2\u00fa\u00fb\7\'\2\2\u00fb"+
		">\3\2\2\2\u00fc\u00fd\7-\2\2\u00fd@\3\2\2\2\u00fe\u00ff\7/\2\2\u00ffB"+
		"\3\2\2\2\u0100\u0101\7*\2\2\u0101D\3\2\2\2\u0102\u0103\7+\2\2\u0103F\3"+
		"\2\2\2\u0104\u0105\7}\2\2\u0105H\3\2\2\2\u0106\u0107\7\177\2\2\u0107J"+
		"\3\2\2\2\u0108\u0109\7]\2\2\u0109L\3\2\2\2\u010a\u010b\7_\2\2\u010bN\3"+
		"\2\2\2\u010c\u010e\t\7\2\2\u010d\u010c\3\2\2\2\u010e\u010f\3\2\2\2\u010f"+
		"\u010d\3\2\2\2\u010f\u0110\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u0112\b("+
		"\2\2\u0112P\3\2\2\2\u0113\u0114\7^\2\2\u0114\u0115\t\b\2\2\u0115R\3\2"+
		"\2\2\f\2\u00a3\u00a8\u00ad\u00b3\u00b9\u00c0\u00cb\u00d9\u010f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}