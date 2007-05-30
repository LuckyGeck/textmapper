package net.sf.lapg.templates.ast;

import java.util.ArrayList;

import net.sf.lapg.templates.api.IEvaluationEnvironment;

public class CompoundNode extends Node {

	protected ArrayList<Node> instructions;

	public ArrayList<Node> getInstructions() {
		return instructions;
	}

	public void setInstructions(ArrayList<Node> instructions) {
		this.instructions = instructions;
	}

	protected void emit(StringBuffer sb, Object context, IEvaluationEnvironment env) {
		if (instructions != null) {
			for (Node n : instructions) {
				n.emit(sb, context, env);
			}
		}
	}
}