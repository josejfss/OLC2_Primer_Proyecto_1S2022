// Generated from c:\Users\GMG\Documents\GitHub\OLC2_Primer_Proyecto_1S2022\ChemsLexer.g4 by ANTLR 4.8
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"RSENTENCIA", "RCONSOLA", "RPUBLICO", "RMAIN", "RINTEGER", "RSTRING", 
			"RREAL", "RBOOLEAN", "RIF", "RENTONCES", "ENTERO", "DECIMAL", "ID", "CADENA", 
			"COMMENT", "LINE_COMMENT", "COMA", "PUNTO", "PUNTOCOMA", "OR", "AND", 
			"NOT", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", "POR", "DIV", 
			"SUMA", "RESTA", "PARA", "PARC", "LLAVEA", "LLAVEC", "CORA", "CORC", 
			"WHITESPACE", "ESC_SEQ"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\'\u0105\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\6\f\u0098"+
		"\n\f\r\f\16\f\u0099\3\r\6\r\u009d\n\r\r\r\16\r\u009e\3\r\3\r\6\r\u00a3"+
		"\n\r\r\r\16\r\u00a4\3\16\3\16\7\16\u00a9\n\16\f\16\16\16\u00ac\13\16\3"+
		"\17\3\17\7\17\u00b0\n\17\f\17\16\17\u00b3\13\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\7\20\u00bb\n\20\f\20\16\20\u00be\13\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\21\3\21\3\21\3\21\7\21\u00c9\n\21\f\21\16\21\u00cc\13\21\3\21\3"+
		"\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3"+
		"\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3"+
		"\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\6&\u00fd"+
		"\n&\r&\16&\u00fe\3&\3&\3\'\3\'\3\'\3\u00bc\2(\3\3\5\4\7\5\t\6\13\7\r\b"+
		"\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26"+
		"+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M\2\3\2"+
		"\t\3\2\62;\4\2C\\c|\6\2\62;C\\aac|\3\2$$\4\2\f\f\17\17\6\2\13\f\17\17"+
		"\"\"^^\t\2\"#%%--/\60<<BB]_\2\u010b\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2"+
		"\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23"+
		"\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2"+
		"\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2"+
		"\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3"+
		"\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2"+
		"\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\3O\3\2\2\2"+
		"\5Y\3\2\2\2\7a\3\2\2\2\ti\3\2\2\2\13n\3\2\2\2\rv\3\2\2\2\17}\3\2\2\2\21"+
		"\u0082\3\2\2\2\23\u008a\3\2\2\2\25\u008d\3\2\2\2\27\u0097\3\2\2\2\31\u009c"+
		"\3\2\2\2\33\u00a6\3\2\2\2\35\u00ad\3\2\2\2\37\u00b6\3\2\2\2!\u00c4\3\2"+
		"\2\2#\u00cf\3\2\2\2%\u00d1\3\2\2\2\'\u00d3\3\2\2\2)\u00d5\3\2\2\2+\u00d8"+
		"\3\2\2\2-\u00db\3\2\2\2/\u00dd\3\2\2\2\61\u00e0\3\2\2\2\63\u00e3\3\2\2"+
		"\2\65\u00e5\3\2\2\2\67\u00e7\3\2\2\29\u00e9\3\2\2\2;\u00eb\3\2\2\2=\u00ed"+
		"\3\2\2\2?\u00ef\3\2\2\2A\u00f1\3\2\2\2C\u00f3\3\2\2\2E\u00f5\3\2\2\2G"+
		"\u00f7\3\2\2\2I\u00f9\3\2\2\2K\u00fc\3\2\2\2M\u0102\3\2\2\2OP\7u\2\2P"+
		"Q\7g\2\2QR\7p\2\2RS\7v\2\2ST\7g\2\2TU\7p\2\2UV\7e\2\2VW\7k\2\2WX\7c\2"+
		"\2X\4\3\2\2\2YZ\7e\2\2Z[\7q\2\2[\\\7p\2\2\\]\7u\2\2]^\7q\2\2^_\7n\2\2"+
		"_`\7c\2\2`\6\3\2\2\2ab\7r\2\2bc\7w\2\2cd\7d\2\2de\7n\2\2ef\7k\2\2fg\7"+
		"e\2\2gh\7q\2\2h\b\3\2\2\2ij\7o\2\2jk\7c\2\2kl\7k\2\2lm\7p\2\2m\n\3\2\2"+
		"\2no\7k\2\2op\7p\2\2pq\7v\2\2qr\7g\2\2rs\7i\2\2st\7g\2\2tu\7t\2\2u\f\3"+
		"\2\2\2vw\7u\2\2wx\7v\2\2xy\7t\2\2yz\7k\2\2z{\7p\2\2{|\7i\2\2|\16\3\2\2"+
		"\2}~\7t\2\2~\177\7g\2\2\177\u0080\7c\2\2\u0080\u0081\7n\2\2\u0081\20\3"+
		"\2\2\2\u0082\u0083\7d\2\2\u0083\u0084\7q\2\2\u0084\u0085\7q\2\2\u0085"+
		"\u0086\7n\2\2\u0086\u0087\7g\2\2\u0087\u0088\7c\2\2\u0088\u0089\7p\2\2"+
		"\u0089\22\3\2\2\2\u008a\u008b\7k\2\2\u008b\u008c\7h\2\2\u008c\24\3\2\2"+
		"\2\u008d\u008e\7g\2\2\u008e\u008f\7p\2\2\u008f\u0090\7v\2\2\u0090\u0091"+
		"\7q\2\2\u0091\u0092\7p\2\2\u0092\u0093\7e\2\2\u0093\u0094\7g\2\2\u0094"+
		"\u0095\7u\2\2\u0095\26\3\2\2\2\u0096\u0098\t\2\2\2\u0097\u0096\3\2\2\2"+
		"\u0098\u0099\3\2\2\2\u0099\u0097\3\2\2\2\u0099\u009a\3\2\2\2\u009a\30"+
		"\3\2\2\2\u009b\u009d\t\2\2\2\u009c\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e"+
		"\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a2\7\60"+
		"\2\2\u00a1\u00a3\t\2\2\2\u00a2\u00a1\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4"+
		"\u00a2\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\32\3\2\2\2\u00a6\u00aa\t\3\2"+
		"\2\u00a7\u00a9\t\4\2\2\u00a8\u00a7\3\2\2\2\u00a9\u00ac\3\2\2\2\u00aa\u00a8"+
		"\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\34\3\2\2\2\u00ac\u00aa\3\2\2\2\u00ad"+
		"\u00b1\7$\2\2\u00ae\u00b0\n\5\2\2\u00af\u00ae\3\2\2\2\u00b0\u00b3\3\2"+
		"\2\2\u00b1\u00af\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2\u00b4\3\2\2\2\u00b3"+
		"\u00b1\3\2\2\2\u00b4\u00b5\7$\2\2\u00b5\36\3\2\2\2\u00b6\u00b7\7*\2\2"+
		"\u00b7\u00b8\7,\2\2\u00b8\u00bc\3\2\2\2\u00b9\u00bb\13\2\2\2\u00ba\u00b9"+
		"\3\2\2\2\u00bb\u00be\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bc\u00ba\3\2\2\2\u00bd"+
		"\u00bf\3\2\2\2\u00be\u00bc\3\2\2\2\u00bf\u00c0\7,\2\2\u00c0\u00c1\7+\2"+
		"\2\u00c1\u00c2\3\2\2\2\u00c2\u00c3\b\20\2\2\u00c3 \3\2\2\2\u00c4\u00c5"+
		"\7\61\2\2\u00c5\u00c6\7\61\2\2\u00c6\u00ca\3\2\2\2\u00c7\u00c9\n\6\2\2"+
		"\u00c8\u00c7\3\2\2\2\u00c9\u00cc\3\2\2\2\u00ca\u00c8\3\2\2\2\u00ca\u00cb"+
		"\3\2\2\2\u00cb\u00cd\3\2\2\2\u00cc\u00ca\3\2\2\2\u00cd\u00ce\b\21\2\2"+
		"\u00ce\"\3\2\2\2\u00cf\u00d0\7.\2\2\u00d0$\3\2\2\2\u00d1\u00d2\7\60\2"+
		"\2\u00d2&\3\2\2\2\u00d3\u00d4\7=\2\2\u00d4(\3\2\2\2\u00d5\u00d6\7~\2\2"+
		"\u00d6\u00d7\7~\2\2\u00d7*\3\2\2\2\u00d8\u00d9\7(\2\2\u00d9\u00da\7(\2"+
		"\2\u00da,\3\2\2\2\u00db\u00dc\7#\2\2\u00dc.\3\2\2\2\u00dd\u00de\7@\2\2"+
		"\u00de\u00df\7?\2\2\u00df\60\3\2\2\2\u00e0\u00e1\7>\2\2\u00e1\u00e2\7"+
		"?\2\2\u00e2\62\3\2\2\2\u00e3\u00e4\7@\2\2\u00e4\64\3\2\2\2\u00e5\u00e6"+
		"\7>\2\2\u00e6\66\3\2\2\2\u00e7\u00e8\7,\2\2\u00e88\3\2\2\2\u00e9\u00ea"+
		"\7\61\2\2\u00ea:\3\2\2\2\u00eb\u00ec\7-\2\2\u00ec<\3\2\2\2\u00ed\u00ee"+
		"\7/\2\2\u00ee>\3\2\2\2\u00ef\u00f0\7*\2\2\u00f0@\3\2\2\2\u00f1\u00f2\7"+
		"+\2\2\u00f2B\3\2\2\2\u00f3\u00f4\7}\2\2\u00f4D\3\2\2\2\u00f5\u00f6\7\177"+
		"\2\2\u00f6F\3\2\2\2\u00f7\u00f8\7]\2\2\u00f8H\3\2\2\2\u00f9\u00fa\7_\2"+
		"\2\u00faJ\3\2\2\2\u00fb\u00fd\t\7\2\2\u00fc\u00fb\3\2\2\2\u00fd\u00fe"+
		"\3\2\2\2\u00fe\u00fc\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ff\u0100\3\2\2\2\u0100"+
		"\u0101\b&\2\2\u0101L\3\2\2\2\u0102\u0103\7^\2\2\u0103\u0104\t\b\2\2\u0104"+
		"N\3\2\2\2\13\2\u0099\u009e\u00a4\u00aa\u00b1\u00bc\u00ca\u00fe\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}