package net.sf.lapg.gen.options;

import java.io.CharArrayReader;
import java.io.IOException;
import java.io.Reader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import net.sf.lapg.gen.options.OptdefLexer.ErrorReporter;
import net.sf.lapg.gen.options.OptdefParser.ParseException;

public class OptdefTree<T> {

	private final TextSource source;
	private final T root;
	private final List<ParseProblem> errors;

	public OptdefTree(TextSource source, T root, List<ParseProblem> errors) {
		this.source = source;
		this.root = root;
		this.errors = errors;
	}

	public TextSource getSource() {
		return source;
	}

	public T getRoot() {
		return root;
	}

	public List<ParseProblem> getErrors() {
		return errors;
	}

	public boolean hasErrors() {
		return errors.size() > 0;
	}


	public static OptdefTree<Object> parse(TextSource source) {
		final List<ParseProblem> list = new ArrayList<ParseProblem>();
		ErrorReporter reporter = new ErrorReporter() {
			public void error(int start, int end, int line, String s) {
				list.add(new ParseProblem(KIND_ERROR, start, end, s, null));
			}
		};

		try {
			OptdefLexer lexer = new OptdefLexer(source.getStream(), reporter);
			lexer.setLine(source.getInitialLine());

			OptdefParser parser = new OptdefParser(reporter);
			Object result = parser.parse(lexer);

			return new OptdefTree<Object>(source, result, list);
		} catch(ParseException ex) {
			/* not parsed */
		} catch(IOException ex) {
			list.add(new ParseProblem(KIND_FATAL, 0, 0, "I/O problem: " + ex.getMessage(), ex));
		}
		return new OptdefTree<Object>(source, null, list);
	}


	public static final int KIND_FATAL = 0;
	public static final int KIND_ERROR = 1;
	public static final int KIND_WARN = 2;

	public static class ParseProblem extends Exception {
		int kind;
		int start;
		int end;

		private static final long serialVersionUID = 1L;

		public ParseProblem(int kind, int start, int end, String message, Throwable cause) {
			super(message, cause);
			this.kind = kind;
			this.start = start;
			this.end = end;
		}

		public int getKind() {
			return kind;
		}

		public int getStart() {
			return start;
		}

		public int getEnd() {
			return end;
		}
	}

	public static class TextSource {

		private final String file;
		private final int initialLine;
		private final char[] contents;
		private int[] lineoffset;

		public TextSource(String file, char[] contents, int initialLine) {
			this.file = file;
			this.initialLine = initialLine;
			this.contents = contents;
		}

		public String getFile() {
			return file;
		}

		public int getInitialLine() {
			return initialLine;
		}

		public Reader getStream() {
			return new CharArrayReader(contents);
		}

		public String getLocation(int offset) {
			return file + "," + (initialLine + lineForOffset(offset));
		}

		public String getText(int start, int end) {
			if (start < 0 || start > contents.length || end > contents.length || start > end) {
				return "";
			}
			return new String(contents, start, end - start);
		}

		public int lineForOffset(int offset) {
			if (lineoffset == null) {
				lineoffset = getLineOffsets(contents);
			}
			int line = Arrays.binarySearch(lineoffset, offset);
			return line > 0 ? line : -line - 1;
		}
		
		public char[] getContents() {
			return contents;
		}
	}

	private static int[] getLineOffsets(char[] contents) {
		int size = 1;
		for (int i = 0; i < contents.length; i++) {
			if (contents[i] == '\n') {
				size++;
			} else if(contents[i] == '\r') {
				if(i+1 < contents.length && contents[i+1] == '\n') {
					i++;
				}
				size++;
			}
		}
		int[] result = new int[size];
		result[0] = 0;
		int e = 1;
		for (int i = 0; i < contents.length; i++) {
			if (contents[i] == '\n') {
				result[e++] = i + 1;
			} else if(contents[i] == '\r') {
				if(i+1 < contents.length && contents[i+1] == '\n') {
					i++;
				}
				result[e++] = i + 1;
			}
		}
		if(e != size) {
			throw new IllegalStateException();
		}
		return result;
	}
}