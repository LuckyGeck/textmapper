package net.sf.lapg.lalr;

import java.util.Iterator;

public abstract class EngineBase {

	protected EngineBase(Grammar g, IError err, int debuglev) {
		this.err = err;
		this.debuglev = debuglev;
		this.errors = 0;

		this.nsyms = g.syms.size();
		this.rules = g.rules.size();
		this.nprio = g.prio.size();
		this.nterms = g.nterms;
		this.input = g.input;
		this.eoi = g.eoi;
		this.errorn = g.errorn;
		this.situations = g.situations;

		this.sym = new Symbol[g.syms.size()];
		for (Iterator<Symbol> it = g.syms.iterator(); it.hasNext();) {
			Symbol s = it.next();
			sym[s.index] = s;
		}

		this.priorul = new int[nprio];

		int i = 0;
		for (Iterator<Integer> it = g.prio.iterator(); it.hasNext();) {
			this.priorul[i++] = it.next();
		}

		this.rleft = new int[rules];
		this.rprio = new int[rules];
		this.rindex = new int[rules];
		this.rlines = new int[rules];
		this.raction = new String[rules];
		this.rright = new int[situations];

		i = 0;
		int curr_rindex = 0;
		for (Iterator<Rule> it = g.rules.iterator(); it.hasNext(); i++) {
			Rule r = it.next();
			this.rleft[i] = r.left;
			this.rprio[i] = r.prio;
			this.rindex[i] = curr_rindex;
			this.rlines[i] = r.line;
			this.raction[i] = r.action;
			for (int e = 0; e < r.right.length; e++)
				this.rright[curr_rindex++] = r.right[e];
		}

		assert situations == curr_rindex;
	}

	// log

	protected final int debuglev, errors;
	protected final IError err;

	// grammar information

	protected final int nsyms, nterms, input, eoi, errorn;
	protected final int rules, situations, nprio;
	protected final Symbol[] sym;
	protected final int[] priorul;
	protected final String[] raction;
	protected final int[] rleft, rindex, rright, rprio, rlines;

	// info

	protected void print_situation(int errl, int situation) {
		int rulenum, i;

		for (i = situation; rright[i] >= 0; i++)
			;
		rulenum = -rright[i] - 1;

		// left part of the rule
		err.error(errl, "  " + sym[rleft[rulenum]].name + " ::=");

		for (i = rindex[rulenum]; rright[i] >= 0; i++) {
			if (i == situation)
				err.error(errl, " _");
			err.error(errl, " " + sym[rright[i]].name);
		}
		if (i == situation)
			err.error(errl, " _");
		err.error(errl, "\n");
	}

	protected void print_rule(int errl, int rule) {
		int rr = rindex[rule];

		err.error(errl, "  " + sym[rleft[rule]].name + " ::=");

		for (; rright[rr] >= 0; rr++)
			err.error(errl, " " + sym[rright[rr]].name);

		err.error(errl, "\n");
	}
	
	public abstract void clear();
}
