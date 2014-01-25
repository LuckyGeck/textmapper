package org.textmapper.tool.test.bootstrap.states;

import java.io.IOException;
import java.io.Reader;
import java.text.MessageFormat;

public class StatesLexer {

	public static class LapgSymbol {
		public Object value;
		public int symbol;
		public int state;
		public int line;
		public int offset;
	}

	public interface States {
		public static final int initial = 0;
		public static final int a = 1;
		public static final int b = 2;
		public static final int c = 3;
		public static final int d = 4;
	}

	public interface Lexems {
		public static final int Unavailable_ = -1;
		public static final int eoi = 0;
		public static final int x = 1;
	}

	public interface ErrorReporter {
		void error(String message, int line, int offset);
	}

	public static final int TOKEN_SIZE = 2048;

	private Reader stream;
	final private ErrorReporter reporter;

	final private char[] data = new char[2048];
	private int datalen, l, tokenStart;
	private char chr;

	private int state;

	final private StringBuilder token = new StringBuilder(TOKEN_SIZE);

	private int tokenLine = 1;
	private int currLine = 1;
	private int currOffset = 0;

	public StatesLexer(Reader stream, ErrorReporter reporter) throws IOException {
		this.reporter = reporter;
		reset(stream);
	}

	public void reset(Reader stream) throws IOException {
		this.stream = stream;
		this.state = 0;
		datalen = stream.read(data);
		l = 0;
		tokenStart = -1;
		chr = l < datalen ? data[l++] : 0;
	}

	protected void advance() throws IOException {
		if (chr == 0) return;
		currOffset++;
		if (chr == '\n') {
			currLine++;
		}
		if (l >= datalen) {
			if (tokenStart >= 0) {
				token.append(data, tokenStart, l - tokenStart);
				tokenStart = 0;
			}
			l = 0;
			datalen = stream.read(data);
		}
		chr = l < datalen ? data[l++] : 0;
	}

	public int getState() {
		return state;
	}

	public void setState(int state) {
		this.state = state;
	}

	public int getTokenLine() {
		return tokenLine;
	}

	public int getLine() {
		return currLine;
	}

	public void setLine(int currLine) {
		this.currLine = currLine;
	}

	public int getOffset() {
		return currOffset;
	}

	public void setOffset(int currOffset) {
		this.currOffset = currOffset;
	}

	public String current() {
		return token.toString();
	}

	private static final short tmCharClass[] = {
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 6, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 13, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 3, 4, 5, 1, 12, 1, 1, 7, 1, 1, 10, 1, 8, 1,
		1, 1, 1, 1, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	};

	private static final short tmStateMap[] = {
		0, 1, 2, 3, 4
	};

	private static final short[] tmRuleSymbol = unpack_short(16,
		"\1\1\1\1\1\1\1\1\1\1\1\1\1\1\1\1");

	private static final int tmClassesCount = 14;

	private static final short[] tmGoto = unpack_vc_short(420,
		"\1\ufffe\1\uffff\1\5\16\uffff\1\6\1\7\1\10\1\11\1\12\5\uffff\1\13\2\uffff\1\14\1" +
		"\uffff\1\15\1\16\1\11\1\12\5\uffff\1\13\2\uffff\1\17\1\20\1\uffff\1\21\1\11\1\12" +
		"\5\uffff\1\13\2\uffff\1\22\1\23\1\24\2\uffff\1\12\6\uffff\16\ufffd\16\ufffc\16\ufffb" +
		"\16\ufffa\16\ufff0\10\uffff\1\25\5\uffff\16\uffee\16\ufff9\16\ufff8\16\ufff7\16\ufff6" +
		"\16\ufff5\16\ufff4\16\ufff3\16\ufff2\16\ufff1\7\uffff\1\26\17\uffff\1\27\13\uffff" +
		"\1\30\10\uffff\1\31\25\uffff\1\32\16\uffff\1\33\16\uffff\1\34\16\uffff\1\35\16\uffef");

	private static short[] unpack_vc_short(int size, String... st) {
		short[] res = new short[size];
		int t = 0;
		int count = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; ) {
				count = i > 0 || count == 0 ? s.charAt(i++) : count;
				if (i < slen) {
					short val = (short) s.charAt(i++);
					while (count-- > 0) res[t++] = val;
				}
			}
		}
		assert res.length == t;
		return res;
	}

	private static int mapCharacter(int chr) {
		if (chr >= 0 && chr < 128) {
			return tmCharClass[chr];
		}
		return 1;
	}

	public LapgSymbol next() throws IOException {
		LapgSymbol lapg_n = new LapgSymbol();
		int state;

		do {
			lapg_n.offset = currOffset;
			tokenLine = lapg_n.line = currLine;
			if (token.length() > TOKEN_SIZE) {
				token.setLength(TOKEN_SIZE);
				token.trimToSize();
			}
			token.setLength(0);
			tokenStart = l - 1;

			for (state = tmStateMap[this.state]; state >= 0; ) {
				state = tmGoto[state * tmClassesCount + mapCharacter(chr)];
				if (state == -1 && chr == 0) {
					lapg_n.symbol = 0;
					lapg_n.value = null;
					reporter.error("Unexpected end of input reached", lapg_n.line, lapg_n.offset);
					lapg_n.offset = currOffset;
					tokenStart = -1;
					return lapg_n;
				}
				if (state >= -1 && chr != 0) {
					currOffset++;
					if (chr == '\n') {
						currLine++;
					}
					if (l >= datalen) {
						token.append(data, tokenStart, l - tokenStart);
						tokenStart = l = 0;
						datalen = stream.read(data);
					}
					chr = l < datalen ? data[l++] : 0;
				}
			}

			if (state == -1) {
				if (l - 1 > tokenStart) {
					token.append(data, tokenStart, l - 1 - tokenStart);
				}
				reporter.error(MessageFormat.format("invalid lexeme at line {0}: `{1}`, skipped", currLine, current()), lapg_n.line, lapg_n.offset);
				lapg_n.symbol = -1;
				continue;
			}

			if (state == -2) {
				lapg_n.symbol = 0;
				lapg_n.value = null;
				tokenStart = -1;
				return lapg_n;
			}

			if (l - 1 > tokenStart) {
				token.append(data, tokenStart, l - 1 - tokenStart);
			}

			lapg_n.symbol = tmRuleSymbol[-state - 3];
			lapg_n.value = null;

		} while (lapg_n.symbol == -1 || !createToken(lapg_n, -state - 3));
		tokenStart = -1;
		return lapg_n;
	}

	protected boolean createToken(LapgSymbol lapg_n, int ruleIndex) throws IOException {
		boolean spaceToken = false;
		switch (ruleIndex) {
			case 0: // x: /a/
				state = States.a;
				break;
			case 1: // x: /b/
				state = States.b;
				break;
			case 2: // x: /c/
				state = States.c;
				break;
			case 3: // x: /d/
				state = States.d;
				break;
			case 4: // x: /a/
				state = States.a;
				break;
			case 5: // x: /c/
				state = States.c;
				break;
			case 6: // x: /d/
				state = States.d;
				break;
			case 7: // x: /a/
				state = States.a;
				break;
			case 8: // x: /b/
				state = States.b;
				break;
			case 9: // x: /d/
				state = States.d;
				break;
			case 10: // x: /a/
				state = States.a;
				break;
			case 11: // x: /b/
				state = States.b;
				break;
			case 12: // x: /c/
				state = States.c;
				break;
			case 13: // x: /!/
				switch(state) {
					case States.b:
						state = States.c;
						break;
					case States.c:
						state = States.d;
						break;
					default:
						state = States.b;
						break;
				}
				break;
			case 14: // x: /initialIfD/
				switch(state) {
					case States.d:
						state = States.initial;
						break;
				}
				break;
			case 15: // x: /D/
				state = States.d;
				break;
		}
		return !(spaceToken);
	}

	/* package */ static int[] unpack_int(int size, String... st) {
		int[] res = new int[size];
		boolean second = false;
		char first = 0;
		int t = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; i++) {
				if (second) {
					res[t++] = (s.charAt(i) << 16) + first;
				} else {
					first = s.charAt(i);
				}
				second = !second;
			}
		}
		assert !second;
		assert res.length == t;
		return res;
	}

	/* package */ static short[] unpack_short(int size, String... st) {
		short[] res = new short[size];
		int t = 0;
		for (String s : st) {
			int slen = s.length();
			for (int i = 0; i < slen; i++) {
				res[t++] = (short) s.charAt(i);
			}
		}
		assert res.length == t;
		return res;
	}
}
