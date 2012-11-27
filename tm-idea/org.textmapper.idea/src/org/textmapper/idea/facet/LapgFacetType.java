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
package org.textmapper.idea.facet;

import com.intellij.facet.Facet;
import com.intellij.facet.FacetType;
import com.intellij.facet.FacetTypeId;
import com.intellij.facet.autodetecting.DetectedFacetPresentation;
import com.intellij.facet.autodetecting.FacetDetector;
import com.intellij.facet.autodetecting.FacetDetectorRegistry;
import com.intellij.openapi.module.Module;
import com.intellij.openapi.module.ModuleType;
import com.intellij.openapi.vfs.VirtualFile;
import com.intellij.openapi.vfs.VirtualFileFilter;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.textmapper.idea.LapgBundle;
import org.textmapper.idea.LapgIcons;
import org.textmapper.idea.lang.syntax.LapgFileType;

import javax.swing.*;
import java.util.Collection;

public class LapgFacetType extends FacetType<LapgFacet, LapgFacetConfiguration> {

	public static final FacetTypeId<LapgFacet> ID = new FacetTypeId<LapgFacet>("lapg");

	public LapgFacetType() {
		super(ID, TmFacetConstants.TM_FACET_ID, TmFacetConstants.TM_FACET_NAME);
	}

	@Override
	public LapgFacetConfiguration createDefaultConfiguration() {
		return new LapgFacetConfiguration();
	}

	@Override
	public LapgFacet createFacet(@NotNull Module module, String name, @NotNull LapgFacetConfiguration configuration, @Nullable Facet underlyingFacet) {
		return new LapgFacet(this, module, name, configuration, underlyingFacet);
	}

	@Override
	public boolean isSuitableModuleType(ModuleType moduleType) {
		return true;
	}

	@Override
	public Icon getIcon() {
		return LapgIcons.LAPG_ICON;
	}

	@Override
	public void registerDetectors(final FacetDetectorRegistry<LapgFacetConfiguration> registry) {
		FacetDetector<VirtualFile, LapgFacetConfiguration> detector = new LapgFacetDetector();
		final boolean[] detected = new boolean[] { false };

		VirtualFileFilter filter = new VirtualFileFilter() {
			public boolean accept(VirtualFile file) {
				if(detected[0]) return true;
				detected[0] = true;
				if(LapgFileType.DEFAULT_EXTENSION.equals(file.getExtension())) {
					registry.customizeDetectedFacetPresentation(new LapgFacetPresentation());
					return true;
				}
				return false;
			}
		};

		registry.registerUniversalDetector(LapgFileType.LAPG_FILE_TYPE, filter, detector);
	}

	private class LapgFacetDetector extends FacetDetector<VirtualFile, LapgFacetConfiguration> {

		private LapgFacetDetector() {
			super("lapg");
		}

		@Override
		public LapgFacetConfiguration detectFacet(VirtualFile source, Collection<LapgFacetConfiguration> existentFacetConfigurations) {
			if (!existentFacetConfigurations.isEmpty()) {
			  return existentFacetConfigurations.iterator().next();
			}
			return createDefaultConfiguration();
		}
	}

	private static class LapgFacetPresentation extends DetectedFacetPresentation {
		@Override
		public String getAutodetectionPopupText(@NotNull Module module, @NotNull FacetType facetType, @NotNull String facetName, @NotNull VirtualFile[] files) {
			return LapgBundle.message("facet.detected");
		}
	}
}
