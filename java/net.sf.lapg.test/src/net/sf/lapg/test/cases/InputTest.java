package net.sf.lapg.test.cases;

import java.io.InputStream;
import java.util.HashMap;

import junit.framework.Assert;
import junit.framework.TestCase;
import net.sf.lapg.LexerTables;
import net.sf.lapg.ParserTables;
import net.sf.lapg.api.Grammar;
import net.sf.lapg.api.Lexem;
import net.sf.lapg.api.Rule;
import net.sf.lapg.api.Symbol;
import net.sf.lapg.gen.OutputUtils;
import net.sf.lapg.input.SyntaxUtil;
import net.sf.lapg.lalr.Builder;
import net.sf.lapg.lex.LexicalBuilder;
import net.sf.lapg.test.ErrorReporter;

public class InputTest extends TestCase {

	private static final String TESTCONTAINER = "net/sf/lapg/test/cases/input";
	private static final String RESULTCONTAINER = "net/sf/lapg/test/cases/expected";

	private InputStream openStream(String name, String root) {
		InputStream is = getClass().getClassLoader().getResourceAsStream(root + "/" + name);
		Assert.assertNotNull(is);
		return is;
	}

	private static void normalizeSpaces(StringBuffer sb) {

		// remove \t
		int i = 0;
		while( (i = sb.indexOf("\t",i)) >= 0 ) {
			sb.replace(i, i+1, " ");
		}
		// several spaces => one space
		i = 0;
		while( (i = sb.indexOf("  ",i)) >= 0 ) {
			int e = i+1;
			while(sb.charAt(e) == ' ') {
				e++;
			}
			sb.replace(i, e, " ");
			i++;
		}
	}

	private void checkGenTables( Grammar g, String outputId ) {
		LexerTables lt = LexicalBuilder.compile(g.getLexems(), new ErrorReporter(), 0);
		ParserTables pt = Builder.compile(g, new ErrorReporter(), 0);

		StringBuffer sb = new StringBuffer();

		OutputUtils.printTables(sb, lt);
		OutputUtils.printTables(sb, pt);

		StringBuffer expected = new StringBuffer(SyntaxUtil.getFileContents(openStream(outputId, RESULTCONTAINER)));
		normalizeSpaces(sb);
		normalizeSpaces(expected);

		Assert.assertEquals(expected.toString().trim(), sb.toString().trim());
	}

	public void testCheckSimple1() {
		Grammar g = SyntaxUtil.parseSyntax("syntax1", openStream("syntax1", TESTCONTAINER), new ErrorReporter(), new HashMap<String,String>());
		Assert.assertNotNull(g);
		Assert.assertEquals(0, g.getEoi());

		Symbol[] syms = g.getSymbols();
		Assert.assertEquals(7, syms.length);
		Assert.assertEquals("eoi", syms[0].getName());
		Assert.assertEquals("identifier", syms[1].getName());
		Assert.assertEquals("Licon", syms[2].getName());
		Assert.assertEquals("_skip", syms[3].getName());  // TODO do not collect skip symbols
		Assert.assertEquals("input", syms[4].getName());
		Assert.assertEquals("list", syms[5].getName());
		Assert.assertEquals("list_item", syms[6].getName());

		Rule[] rules = g.getRules();
		Assert.assertEquals(5, rules.length);
		Assert.assertEquals("input", rules[0].getLeft().getName());
		Assert.assertEquals("list", rules[0].getRight()[0].getName());
		Assert.assertEquals(1, rules[0].getRight().length);

		Lexem[] lexems = g.getLexems();
		Assert.assertEquals(3, lexems.length);
		Assert.assertEquals("@?[a-zA-Z_][A-Za-z_0-9]*", lexems[0].getRegexp());
		Assert.assertEquals("([1-9][0-9]*|0[0-7]*|0[xX][0-9a-fA-F]+)([uU](l|L|ll|LL)?|(l|L|ll|LL)[uU]?)?", lexems[1].getRegexp());
		Assert.assertEquals("[\\t\\r\\n ]+", lexems[2].getRegexp());
		Assert.assertEquals(" continue; ", lexems[2].getAction());

		checkGenTables(g, "syntax1.tbl");
	}

	public void testCheckSimple2() {
		Grammar g = SyntaxUtil.parseSyntax("syntax2", openStream("syntax2", TESTCONTAINER), new ErrorReporter(), new HashMap<String,String>());
		Assert.assertNotNull(g);
		Assert.assertEquals(0, g.getEoi());

		Symbol[] syms = g.getSymbols();
		Assert.assertEquals(9, syms.length);
		Assert.assertEquals("eoi", syms[0].getName());
		Assert.assertEquals("a", syms[1].getName());
		Assert.assertEquals("b", syms[2].getName());
		Assert.assertEquals("'('", syms[3].getName());
		Assert.assertEquals("')'", syms[4].getName());
		Assert.assertEquals("input", syms[5].getName());
		Assert.assertEquals("list", syms[6].getName());
		Assert.assertEquals("item", syms[7].getName());
		Assert.assertEquals("listopt", syms[8].getName());
		Assert.assertEquals(8, g.getRules().length);
		Assert.assertEquals("  ${for a in b}..!..$$  ", g.getRules()[7].getAction());

		checkGenTables(g, "syntax2.tbl");
	}
}
