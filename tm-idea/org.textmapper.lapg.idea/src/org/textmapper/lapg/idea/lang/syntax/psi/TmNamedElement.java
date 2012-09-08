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
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiNamedElement;
import com.intellij.psi.util.PsiTreeUtil;
import com.intellij.util.IncorrectOperationException;
import org.jetbrains.annotations.NonNls;
import org.jetbrains.annotations.NotNull;

/**
 * Gryaznov Evgeny, 1/26/11
 */
public abstract class TmNamedElement extends TmElement implements PsiNamedElement {

	public TmNamedElement(@NotNull ASTNode node) {
		super(node);
	}

	public TmSymbol getNameSymbol() {
		return PsiTreeUtil.getChildOfType(this, TmSymbol.class);
	}

	public String getName() {
		TmSymbol nameSymbol = getNameSymbol();
		return nameSymbol != null ? nameSymbol.getText() : null;
	}

	public PsiElement setName(@NonNls String name) throws IncorrectOperationException {
		TmSymbol nameSymbol = getNameSymbol();
		if (nameSymbol == null) {
			throw new IncorrectOperationException();
		}
		nameSymbol.replace(TmElementsFactory.createSymbol(getProject(), name));
		return this;
	}

}