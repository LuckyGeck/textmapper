/**
 * Copyright 2002-2013 Evgeny Gryaznov
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

import org.textmapper.lapg.api.Nonterminal;
import org.textmapper.lapg.api.SourceElement;
import org.textmapper.lapg.api.rule.RhsPart;

/**
 * evgeny, 10/27/12
 */
public class LiNonterminal extends LiSymbol implements Nonterminal {

	private RhsPart definition;

	public LiNonterminal(String name, String type, SourceElement origin) {
		super(name, type, origin);
	}

	void setDefinition(RhsPart part) {
		if (definition != null) {
			throw new IllegalStateException("non-terminal is sealed");
		}
		this.definition = part;
	}

	void addRule(RhsPart part) {
		if (definition == null) {
			definition = new LiRootRhsChoice(this);
		} else if (!(definition instanceof LiRootRhsChoice)) {
			throw new IllegalStateException("non-terminal is sealed");
		}
		((LiRootRhsChoice) definition).addRule(part);
	}

	@Override
	public RhsPart getDefinition() {
		return definition;
	}
}