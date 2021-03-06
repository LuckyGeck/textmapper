/**
 * Copyright 2002-2018 Evgeny Gryaznov
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
import org.textmapper.lapg.api.ast.AstType;
import org.textmapper.lapg.api.builder.GrammarBuilder;

import java.util.Collections;

/**
 * evgeny, 6/29/12
 */
public class GrammarFacadeErrTest {

	@Test(expected = NullPointerException.class)
	public void testErrorNullName() {
		GrammarFacade.createBuilder().addTerminal(null, null, null);
	}

	@Test(expected = IllegalStateException.class)
	public void testErrorDuplicateName() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		builder.addTerminal(LapgCore.name("sym"), null, null);
		builder.addTerminal(LapgCore.name("sym"), null, null);
	}

	@Test(expected = NullPointerException.class)
	public void testErrorNullSymbol() throws Exception {
		GrammarFacade.createBuilder().addLexerRule(
				LexerRule.KIND_NONE, null, LapgCore.parse("id", "a"), Collections.<LexerState>emptyList(), 0, 0, null, null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testWrongState() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal symbol = builder.addTerminal(LapgCore.name("sym"), null, null);
		builder.addLexerRule(LexerRule.KIND_NONE, symbol, LapgCore.parse("id", "a"),
				Collections.singletonList(GrammarFacade.createBuilder().addState(LapgCore.name("initial"), null)), 0, 0, null, null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testNoStates() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal symbol = builder.addTerminal(LapgCore.name("sym"), null, null);
		builder.addLexerRule(LexerRule.KIND_NONE, symbol, LapgCore.parse("id", "a"),
				Collections.<LexerState>emptyList(), 0, 0, null, null);
	}

	@Test(expected = NullPointerException.class)
	public void testNullState() throws Exception {
		GrammarFacade.createBuilder().addState(null, null);
	}

	@Test(expected = IllegalStateException.class)
	public void testDuplicateState() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		builder.addState(LapgCore.name("name"), null);
		builder.addState(LapgCore.name("name"), null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testErrorCrossBuilderSymbols() throws Exception {
		Terminal symbolFromAnotherBuilder = GrammarFacade.createBuilder().addTerminal(
				LapgCore.name("sym"), null, null);
		GrammarFacade.createBuilder().addLexerRule(
				LexerRule.KIND_NONE, symbolFromAnotherBuilder, LapgCore.parse("id", "a"), null, 0, 0, null, null);
	}

	@Test(expected = NullPointerException.class)
	public void testErrorNoRegexp() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal term = builder.addTerminal(LapgCore.name("sym"), null, null);
		builder.addLexerRule(LexerRule.KIND_NONE, term, null, Collections.<LexerState>emptyList(), 0, 0, null, null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testErrorSoftVsNonSoft() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal classterm = builder.addTerminal(LapgCore.name("id"), raw("string"), null);
		Terminal softterm = builder.addTerminal(LapgCore.name("keyword"), null, null);
		builder.makeSoft(softterm, classterm);
		builder.addLexerRule(LexerRule.KIND_NONE, softterm, LapgCore.parse("id", "a"), null, 0, 0, null, null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testErrorSoftVsNonSoft2() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal term = builder.addTerminal(LapgCore.name("s"), raw("s"), null);
		builder.addLexerRule(LexerRule.KIND_SOFT, term, LapgCore.parse("id", "a"), null, 0, 0, null, null);
	}

	@Test(expected = NullPointerException.class)
	public void testErrorSoftNullClass() throws Exception {
		GrammarFacade.createBuilder().makeSoft(null, null);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testErrorClassOfItself() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal symbol = builder.addTerminal(LapgCore.name("sym"), raw("string"), null);
		builder.makeSoft(symbol, symbol);
	}

	@Test(expected = IllegalStateException.class)
	public void testMakeSoftTwice() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal cl = builder.addTerminal(LapgCore.name("sym"), raw("string"), null);
		Terminal symbol = builder.addTerminal(LapgCore.name("ssym"), null, null);
		builder.makeSoft(symbol, cl);
		builder.makeSoft(symbol, cl);
	}

	@Test(expected = IllegalStateException.class)
	public void testMakeSoftBadTypes() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal cl = builder.addTerminal(LapgCore.name("sym"), raw("string"), null);
		Terminal symbol = builder.addTerminal(LapgCore.name("ssym"), raw("int"), null);
		builder.makeSoft(symbol, cl);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testMakeSoftSealed() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal cl = builder.addTerminal(LapgCore.name("sym"), null, null);
		Terminal symbol = builder.addTerminal(LapgCore.name("ssym"), null, null);
		builder.makeSoft(symbol, cl);
		builder.makeSoft(cl, symbol);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testMakeSoftWithSoftClass() {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		Terminal cl = builder.addTerminal(LapgCore.name("sym"), null, null);
		Terminal symbol = builder.addTerminal(LapgCore.name("ssym"), null, null);
		builder.makeSoft(symbol, cl);
		Terminal sym3 = builder.addTerminal(LapgCore.name("sym3"), null, null);
		builder.makeSoft(sym3, symbol);
	}

	@Test(expected = IllegalArgumentException.class)
	public void testErrorBadPriority() throws Exception {
		GrammarFacade.createBuilder().addPrio(100, Collections.<Terminal>emptyList(), null);
	}

	@Test(expected = NullPointerException.class)
	public void testErrorPatternNoName() throws Exception {
		GrammarFacade.createBuilder().addPattern(null, LapgCore.parse("a", "a"), null);
	}

	@Test(expected = NullPointerException.class)
	public void testErrorPatternNoRegex() throws Exception {
		GrammarFacade.createBuilder().addPattern(LapgCore.name("a"), null, null);
	}

	@Test(expected = IllegalStateException.class)
	public void testErrorDuplicatePattern() throws Exception {
		GrammarBuilder builder = GrammarFacade.createBuilder();
		builder.addPattern(LapgCore.name("a"), LapgCore.parse("a", "a"), null);
		builder.addPattern(LapgCore.name("a"), LapgCore.parse("b", "b"), null);
	}

	private static AstType raw(String type) {
		return GrammarFacade.createAstBuilder().rawType(type, null);
	}
}
