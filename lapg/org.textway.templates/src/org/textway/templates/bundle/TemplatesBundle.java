/**
 * Copyright 2002-2010 Evgeny Gryaznov
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
package org.textway.templates.bundle;

import org.textway.templates.api.SourceElement;
import org.textway.templates.api.TemplatesStatus;
import org.textway.templates.ast.TemplatesTree;
import org.textway.templates.ast.TemplatesTree.TemplatesProblem;
import org.textway.templates.ast.TemplatesTree.TextSource;
import org.textway.templates.storage.Resource;

import java.util.List;

public class TemplatesBundle {

	private final Resource resource;
	private final IBundleEntity[] entities;

	public TemplatesBundle(Resource resource, IBundleEntity[] entities) {
		this.resource = resource;
		this.entities = entities;
	}

	public String getName() {
		return resource.getUri().getPath();
	}

	public IBundleEntity[] getEntities() {
		return entities;
	}

	public static TemplatesBundle parse(final Resource resource, String templatePackage,
										final TemplatesStatus status) {

		TextSource source = new TextSource(resource.getUri().getPath(), resource.getContents().toCharArray(), resource.getInitialLine());
		final TemplatesTree<List<IBundleEntity>> tree = TemplatesTree.parseInput(source, templatePackage);
		for (final TemplatesProblem problem : tree.getErrors()) {
			status.report(TemplatesStatus.KIND_ERROR, problem.getMessage(), new SourceElement() {
				public String getResourceName() {
					return resource.getUri().getPath();
				}

				public int getOffset() {
					return problem.getOffset();
				}

				public int getEndOffset() {
					return problem.getEndOffset();
				}

				public int getLine() {
					return tree.getSource().lineForOffset(problem.getOffset());
				}
			});
		}
		IBundleEntity[] entities = tree.getRoot() != null ? tree.getRoot().toArray(new IBundleEntity[tree.getRoot().size()]) : new IBundleEntity[0];
		return new TemplatesBundle(resource, entities);
	}
}
