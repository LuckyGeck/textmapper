/**
 * Copyright (c) 2010-2012 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textmapper.lapg.idea.lang.syntax.psi;

import com.intellij.lang.ASTNode;
import com.intellij.psi.util.PsiTreeUtil;
import org.jetbrains.annotations.NotNull;

import java.util.List;

/**
 * Gryaznov Evgeny, 1/26/11
 */
public class LpsGrammar extends LpsElement {

	public LpsGrammar(@NotNull ASTNode node) {
		super(node);
	}

	public LpsNamedElement[] getNamedElements() {
		return PsiTreeUtil.getChildrenOfType(this, LpsNamedElement.class);
	}

	public LpsNamedElement resolve(String name) {
		if (name.endsWith("opt") && name.length() > 3) {
			name = name.substring(0, name.length() - 3);
		}

		for (LpsNamedElement named : getNamedElements()) {
			if (name.equals(named.getName())) {
				return named;
			}
		}
		return null;
	}

	public List<LpsLexem> getLexems() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, LpsLexem.class);
	}

	public List<LpsNonTerm> getNonTerms() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, LpsNonTerm.class);
	}
}
