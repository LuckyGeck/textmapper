/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.lapg.builder;

import org.junit.Test;
import org.textmapper.lapg.LapgCore;
import org.textmapper.lapg.api.*;
import org.textmapper.lapg.api.builder.GrammarBuilder;
import org.textmapper.lapg.api.builder.RuleBuilder;

import java.util.Collections;

import static org.junit.Assert.*;

/**
 * evgeny, 6/29/12
 */
public class GrammarFacadeTest {

	@Test
	public void testSimpleGrammar() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();

		// id: /[a-z][a-z0-9]+/
		// input ::= id ;

		Symbol id = builder.addSymbol(Symbol.KIND_TERM, "id", null, null);
		builder.addLexem(Lexem.KIND_NONE, id, LapgCore.parse("id", "[a-z][a-z0-9]+"), 1, 0, null, null);

		Symbol input = builder.addSymbol(Symbol.KIND_NONTERM, "input", null, null);
		builder.addInput(input, true, null);
		RuleBuilder rule = builder.rule(null, input, null);
		rule.addPart(null, id, null, null);
		rule.create();

		Grammar grammar = builder.create();

		Symbol[] symbols = grammar.getSymbols();
		assertEquals(3, symbols.length);
		assertTrue(grammar.getEoi() == symbols[0]);
		assertEquals("id", symbols[1].getName());
		assertEquals(Symbol.KIND_TERM, symbols[1].getKind());
		assertEquals("input", symbols[2].getName());
		assertEquals(Symbol.KIND_NONTERM, symbols[2].getKind());
		assertEquals("non-terminal", symbols[2].kindAsString());
		for (int i = 0; i < symbols.length; i++) {
			assertEquals(i, symbols[i].getIndex());
		}

		Rule[] rules = grammar.getRules();
		assertEquals(1, rules.length);
		assertEquals(id.getIndex(), grammar.getRules()[0].getPriority());
		assertTrue(symbols[2] == rules[0].getLeft());
		assertEquals(1, rules[0].getRight().length);
		assertTrue(symbols[1] == rules[0].getRight()[0].getTarget());
		assertNull(rules[0].getRight()[0].getAlias());
		assertNull(rules[0].getRight()[0].getNegativeLA());

		// lexer
		Lexem[] lexems = grammar.getLexems();
		assertEquals(1, lexems.length);
		assertFalse(lexems[0].isExcluded());
		assertTrue(symbols[1] == lexems[0].getSymbol());
		assertNull(lexems[0].getClassLexem());
		assertEquals("none", lexems[0].getKindAsText());
		assertEquals("[a-z][a-z0-9]+", lexems[0].getRegexp().toString());
		assertEquals(1, lexems[0].getGroups());

		// input
		InputRef[] inputRefs = grammar.getInput();
		assertEquals(1, inputRefs.length);
		assertNull(((DerivedSourceElement)inputRefs[0]).getOrigin());
		assertEquals(symbols[2], inputRefs[0].getTarget());
		assertEquals(true, inputRefs[0].hasEoi());

		// grammar
		assertNull(grammar.getError());
		assertEquals(2, grammar.getTerminals());
		assertEquals(3, grammar.getGrammarSymbols());

		// empty lists
		assertEquals(0, grammar.getPatterns().length);
		assertEquals(0, grammar.getPriorities().length);
	}

	@Test
	public void testLexerOnlyGrammar() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();

		// id: /[a-z]+/  (class)
		// kw:  /keyword/
		// spc: /[\t ]+/

		Symbol id = builder.addSymbol(Symbol.KIND_TERM, "id", "string", null);
		Lexem idLexem = builder.addLexem(Lexem.KIND_CLASS, id, LapgCore.parse("id", "[a-z]+"), 1, 0, null, null);
		Symbol kw = builder.addSoftSymbol("kw", id, null);
		builder.addLexem(Lexem.KIND_SOFT, kw, LapgCore.parse("kw", "keyword"), 1, 0, idLexem, null);
		Symbol spc = builder.addSymbol(Symbol.KIND_TERM, "spc", null, null);
		builder.addLexem(Lexem.KIND_SPACE, spc, LapgCore.parse("spc", "[\t ]+"), 1, 0, null, null);
		Grammar grammar = builder.create();

		Symbol[] symbols = grammar.getSymbols();
		assertEquals(4, symbols.length);
		assertTrue(grammar.getEoi() == symbols[0]);

		// id
		assertEquals("id", symbols[1].getName());
		assertEquals(Symbol.KIND_TERM, symbols[1].getKind());
		assertEquals("terminal", symbols[1].kindAsString());
		assertFalse(symbols[1].isSoft());
		assertNull(symbols[1].getSoftClass());

		// kw
		assertEquals("kw", symbols[2].getName());
		assertEquals(Symbol.KIND_SOFTTERM, symbols[2].getKind());
		assertEquals("soft-terminal", symbols[2].kindAsString());
		assertTrue(symbols[2].isSoft());
		assertTrue(symbols[1] == symbols[2].getSoftClass());

		// spc
		assertEquals("spc", symbols[3].getName());
		assertEquals(Symbol.KIND_TERM, symbols[3].getKind());
		assertEquals("terminal", symbols[3].kindAsString());
		for (int i = 0; i < symbols.length; i++) {
			assertEquals(i, symbols[i].getIndex());
			assertNull(((DerivedSourceElement)symbols[i]).getOrigin());
		}

		Lexem[] lexems = grammar.getLexems();
		for (int i = 0; i < lexems.length; i++) {
			assertEquals(i, lexems[i].getIndex());
			assertNull(((DerivedSourceElement)lexems[i]).getOrigin());
			assertEquals(0, lexems[i].getPriority());
		}
		assertEquals(3, lexems.length);
		// id
		assertTrue(symbols[1] == lexems[0].getSymbol());
		assertNull(lexems[0].getClassLexem());
		assertEquals("[a-z]+", lexems[0].getRegexp().toString());
		assertEquals(1, lexems[0].getGroups());
		assertEquals(Lexem.KIND_CLASS, lexems[0].getKind());
		assertEquals("class", lexems[0].getKindAsText());
		// kw
		assertTrue(symbols[2] == lexems[1].getSymbol());
		assertTrue(lexems[1].getClassLexem() == lexems[0]);
		assertEquals("keyword", lexems[1].getRegexp().toString());
		assertTrue(lexems[1].getRegexp().isConstant());
		assertEquals("keyword", lexems[1].getRegexp().getConstantValue());
		assertEquals(1, lexems[1].getGroups());
		assertEquals(Lexem.KIND_SOFT, lexems[1].getKind());
		assertEquals("soft", lexems[1].getKindAsText());
		assertTrue(lexems[1].isExcluded());
		// spc
		assertTrue(symbols[3] == lexems[2].getSymbol());
		assertEquals(Lexem.KIND_SPACE, lexems[2].getKind());
		assertEquals("space", lexems[2].getKindAsText());
		assertEquals(1, lexems[2].getGroups());

		// empty lists
		assertNull(grammar.getRules());
		assertEquals(0, grammar.getPatterns().length);
		assertNull(grammar.getPriorities());
		assertNull(grammar.getInput());
		assertNull(grammar.getError());

		//
		assertEquals(4, grammar.getTerminals());
		assertEquals(4, grammar.getGrammarSymbols());
	}

	@Test
	public void testEoi() throws Exception {
		Symbol eoi = GrammarFacade.createBuilder().getEoi();
		assertEquals(0, eoi.getIndex());
		assertEquals("eoi", eoi.getName());
		assertNull(eoi.getType());
		assertEquals("terminal", eoi.kindAsString());
		assertEquals(Symbol.KIND_TERM, eoi.getKind());
	}

	@Test
	public void testPrio() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();

		// id: /[a-z][a-z0-9]+/
		// input ::= id ;

		Symbol id = builder.addSymbol(Symbol.KIND_TERM, "id", null, null);
		builder.addLexem(Lexem.KIND_NONE, id, LapgCore.parse("id", "[a-z][a-z0-9]+"), 1, 0, null, null);

		Symbol input = builder.addSymbol(Symbol.KIND_NONTERM, "input", null, null);
		builder.addInput(input, true, null);
		RuleBuilder rule = builder.rule(null, input, null);
		rule.addPart(null, id, null, null);
		rule.setPriority(id);
		rule.create();

		assertNotNull(builder.addPrio(Prio.RIGHT, Collections.singleton(id), null));

		Grammar grammar = builder.create();

		assertEquals(1, grammar.getRules().length);
		assertEquals(id.getIndex(), grammar.getRules()[0].getPriority());
		assertEquals(1, grammar.getPriorities().length);
		assertEquals(Prio.RIGHT, grammar.getPriorities()[0].getPrio());
		assertEquals(1, grammar.getPriorities()[0].getSymbols().length);
		assertTrue(id == grammar.getPriorities()[0].getSymbols()[0]);
		assertNull(((DerivedSourceElement)grammar.getPriorities()[0]).getOrigin());
	}

	@Test
	public void testNamedPatterns() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();

		// pattern = /[a-z]+/
		// id:  /{pattern}/

		builder.addPattern("pattern", LapgCore.parse("pattern", "[a-z]+"), null);
		Symbol id = builder.addSymbol(Symbol.KIND_TERM, "id", "string", null);
		builder.addLexem(Lexem.KIND_NONE, id, LapgCore.parse("id", "{pattern}"), 1, 0, null, null);
		Grammar grammar = builder.create();

		NamedPattern[] patterns = grammar.getPatterns();
		assertEquals(1, patterns.length);
		assertEquals("pattern", patterns[0].getName());
		assertEquals("[a-z]+", patterns[0].getRegexp().toString());
		assertNull(((DerivedSourceElement)patterns[0]).getOrigin());
	}
}