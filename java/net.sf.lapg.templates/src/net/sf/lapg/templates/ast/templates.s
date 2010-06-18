#  syntax: lapg templates source grammar
#
#  Lapg (Lexer and Parser Generator)
#  Copyright 2002-2010 Evgeny Gryaznov
# 
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

.lang        "java" 
.prefix      "Ast"
.package	 "net.sf.lapg.templates.ast"
.breaks		 "on"
.lexemend    "on"
.packLexems  "true"
.positions   "line,offset"
.endpositions "offset"
.gentree	 "on"

# Vocabulary

[0]

any:	/[^$]+/

escdollar:		/$$/
escid(String):	/$[a-zA-Z_][A-Za-z_0-9]*(#[0-9]+)?/	{ $lexem = token.toString().substring(1, token.length()); break; }
escint(Integer):/$[0-9]+/							{ $lexem = Integer.parseInt(token.toString().substring(1, token.length())); break; }

'${':	/${/		{ group = 1; break; }
'$/':   /$\//

[1]

identifier(String):	/[a-zA-Z_][A-Za-z_0-9]*/ -1		{ $lexem = current(); break; }

icon(Integer):	/[0-9]+/							{ $lexem = Integer.parseInt(current()); break; }
ccon(String):	/'([^\n\\']|\\(['"?\\abfnrtv]|x[0-9a-fA-F]+|[0-7]([0-7][0-7]?)?))*'/	{ $lexem = token.toString().substring(1, token.length()-1); break; }

Lcall:		/call/
Lcached:	/cached/
Lcase:		/case/
Lend:		/end/
Lelse:		/else/
Leval:		/eval/
Lfalse:		/false/
Lfor:		/for/
Lfile:		/file/
Lforeach:	/foreach/
Lgrep:		/grep/
Lif:		/if/
Lin:		/in/
Limport:	/import/
Lis:		/is/
Lmap:		/map/
Lnull:		/null/
Lquery:		/query/
Lswitch:	/switch/
Ltemplate:  /template/
Ltrue:		/true/
Lself:		/self/
Lassert:	/assert/

'}':		/}/			{ group = 0; break; }
'-}':		/-}/		{ group = 0; break; }
'+':		/+/
'-':		/-/
'*':		/*/
'/':		/\//
'%':		/%/
'!':		/!/
'|':		/\|/
'[':		/\[/
']':		/\]/
'(':		/\(/
')':		/\)/
'.':		/\./
',':		/,/
'&&':		/&&/
'||':		/\|\|/
'==':		/==/
'=':		/=/
'!=':		/!=/
'->':		/->/
'=>':		/=>/
'<=':		/<=/
'>=':		/>=/
'<':		/</
'>':		/>/
':':		/:/
'?':		/?/

_skip:      /[\t\r\n ]+/    { return false; }

# Grammar

input ::=
	templatesopt 
;

templates ::=
	templates template_declaration_or_space 
	| template_declaration_or_space 
;

template_declaration_or_space ::=
	template_start instructions template_end		{ $template_start.setInstructions($instructions); entities.add($template_start); }
	| template_start template_end					{ entities.add($template_start); }
	| query_def										{ entities.add($query_def); }
	| any
;

query_def (QueryNode) ::=
	'${' cached_flagopt Lquery qualified_id parametersopt '=' expression '}'
                                                    { $$ = new QueryNode($qualified_id, $parametersopt, templatePackage, $expression, $cached_flagopt != null, source, ${query_def.offset}, ${query_def.endoffset}); checkFqn($qualified_id, ${query_def.offset}, ${query_def.endoffset}, ${self[0].line}); }
;

cached_flag (Boolean) ::=
	Lcached											{ $$ = Boolean.TRUE; }
;

template_start (TemplateNode) ::=
	'${' Ltemplate qualified_id parametersopt '[-]}'
                                                    { $$ = new TemplateNode($qualified_id, $parametersopt, templatePackage, source, ${template_start.offset}, ${template_start.endoffset}); checkFqn($qualified_id, ${template_start.offset}, ${template_start.endoffset}, ${self[0].line}); }
; 

parameters (ArrayList) ::=
	'(' identifier_listopt ')' 						{ $$ = $1; }
;

identifier_list (ArrayList) ::=
	  identifier                                    { $$ = new ArrayList(); $identifier_list.add($identifier); }
	| identifier_list ',' identifier                { $identifier_list#0.add($identifier); }
;

template_end ::=
	'${' Lend '}' ;

instructions (ArrayList<Node>) ::=
	instructions instruction						{ $instructions#0.add($instruction); }
	| instruction 									{ $$ = new ArrayList<Node>(); $instructions.add($instruction); }
;

'[-]}' ::=
	'-}'											{ skipSpaces(${self[0].offset}+1); }
	| '}'
;

instruction (Node) ::=
	  control_instruction
	| switch_instruction
	| simple_instruction
	| escid											{ $$ = createEscapedId($escid, ${instruction.offset}, ${instruction.endoffset}); }
	| escint										{ $$ = new IndexNode(null, new LiteralNode($escint, source, ${instruction.offset}, ${instruction.endoffset}), source, ${instruction.offset}, ${instruction.endoffset}); }
	| escdollar										{ $$ = new DollarNode(source, ${instruction.offset}, ${instruction.endoffset}); }
	| any											{ $$ = new TextNode(source, rawText(${instruction.offset}, ${instruction.endoffset}), ${instruction.endoffset}); }
;

simple_instruction (Node) ::=
	'${' sentence '[-]}' 							{ $$ = $1; } 
;

sentence (Node) ::=
	  expression
	| Lcall qualified_id template_argumentsopt template_for_expropt
													{ $$ = new CallTemplateNode($qualified_id, $template_argumentsopt, $template_for_expropt, templatePackage, true, source, ${sentence.offset},${sentence.endoffset}); }
	| Leval conditional_expression comma_expropt	{ $$ = new EvalNode($conditional_expression, $comma_expropt, source, ${sentence.offset},${sentence.endoffset}); }
	| Lassert expression							{ $$ = new AssertNode($expression, source, ${sentence.offset},${sentence.endoffset}); }
;

comma_expr (ExpressionNode) ::=
	',' conditional_expression						{ $$ = $conditional_expression; }
;

qualified_id (String) ::=
	identifier
	| qualified_id '.' identifier					{ $$ = $qualified_id#0 + "." + $identifier; }
;

template_for_expr (ExpressionNode) ::=
	Lfor expression									{ $$ = $1; }
;

template_arguments (ArrayList) ::=
	'(' expression_listopt ')'						{ $$ = $1; } 
;

control_instruction (CompoundNode) ::=
	control_start instructions else_clause 			{ $control_instruction.setInstructions($instructions); applyElse($control_start,$else_clause, ${left().offset}, ${left().endoffset}, ${left().line}); }
;

else_clause (ElseIfNode) ::=
	  '${' Lelse Lif expression '[-]}' instructions else_clause
	  												{ $$ = new ElseIfNode($expression, $instructions, $else_clause#1, source, ${first().offset}, ${instructions.endoffset}); }
	| '${' Lelse '[-]}' instructions control_end
													{ $$ = new ElseIfNode(null, $instructions, null, source, ${first().offset}, ${instructions.endoffset}); }
	| control_end
													{ $$ = null; }
;   

switch_instruction (CompoundNode) ::=
	'${' Lswitch expression '[-]}' anyopt 
           case_list control_end            		{ $$ = new SwitchNode($expression, $case_list, null, source, ${left().offset},${left().endoffset}); checkIsSpace(${anyopt.offset},${anyopt.endoffset}, ${anyopt.line}); }
	| '${' Lswitch expression '[-]}' anyopt 
           case_list '${' Lelse '[-]}' instructions control_end
													{ $$ = new SwitchNode($expression, $case_list, $instructions, source, ${left().offset},${left().endoffset}); checkIsSpace(${anyopt.offset},${anyopt.endoffset}, ${anyopt.line}); }
;

case_list (ArrayList) ::=
	one_case										{ $$ = new ArrayList(); $case_list.add($one_case); }	
	| case_list one_case                            { $case_list#0.add($one_case); }
	| case_list instruction                         { CaseNode.add($case_list#0, $instruction); }
;

one_case (CaseNode) ::=
	'${' Lcase expression '[-]}' 					{ $$ = new CaseNode($expression, source, ${one_case.offset},${one_case.endoffset}); } 
;

control_start (CompoundNode) ::=
	'${' control_sentence '[-]}' 					{ $$ = $1; } ;

control_sentence (CompoundNode) ::=
	  Lforeach identifier Lin expression			{ $$ = new ForeachNode($identifier, $expression, source, ${control_sentence.offset}, ${control_sentence.endoffset}); }
	| Lfor identifier Lin '[' conditional_expression ',' conditional_expression ']'
													{ $$ = new ForeachNode($identifier, $conditional_expression#0, $conditional_expression#1, source, ${control_sentence.offset}, ${control_sentence.endoffset}); }
	| Lif expression								{ $$ = new IfNode($expression, source, ${control_sentence.offset}, ${control_sentence.endoffset}); }
	| Lfile expression								{ $$ = new FileNode($expression, source, ${control_sentence.offset}, ${control_sentence.endoffset}); }
;

control_end ::=
	'${' Lend '[-]}'
	| '$/'
;

primary_expression (ExpressionNode) ::=
  	  identifier									{ $$ = new SelectNode(null, $identifier, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
    | '(' expression ')'							{ $$ = new ParenthesesNode($1, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
	| icon 											{ $$ = new LiteralNode($0, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
	| bcon                                          { $$ = new LiteralNode($0, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
	| ccon 											{ $$ = new LiteralNode($0, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
  	| Lself											{ $$ = new ThisNode(source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
  	| Lnull											{ $$ = new LiteralNode(null, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
    | identifier '(' expression_listopt ')'         { $$ = new MethodCallNode(null, $identifier, $expression_listopt, source, ${primary_expression.offset}, ${primary_expression.endoffset}); }
    | primary_expression '.' identifier				{ $$ = new SelectNode($primary_expression#1, $identifier, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}); }
    | primary_expression '.' identifier '(' expression_listopt ')'   
    												{ $$ = new MethodCallNode($primary_expression#1, $identifier, $expression_listopt, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}); }
    | primary_expression '.' identifier '(' identifier '|' expression ')'
    												{ $$ = createCollectionProcessor($primary_expression#1, $identifier#0, $identifier#1, $expression, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}, ${primary_expression[0].line}); }
    | primary_expression '.' identifier '(' var=identifier '|' key=expression ':' value=expression ')'
    												{ $$ = createMapCollect($primary_expression#1, $identifier, $var, $key, $value, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}, ${primary_expression[0].line}); }
    | primary_expression '->' qualified_id '(' expression_listopt ')'
    												{ $$ = new CallTemplateNode($qualified_id, $expression_listopt, $primary_expression#1, templatePackage, false, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}); }
    | primary_expression '->' '(' expression ')' '(' expression_listopt ')'  
    												{ $$ = new CallTemplateNode($expression,$expression_listopt,$primary_expression#1,templatePackage, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}); }
    | primary_expression '[' expression ']'			{ $$ = new IndexNode($primary_expression#1, $expression, source, ${primary_expression[0].offset}, ${primary_expression[0].endoffset}); }
    | complex_data
;

complex_data (ExpressionNode) ::=
	'[' expression_listopt ']'						{ $$ = new ListNode($expression_listopt, source, ${complex_data.offset}, ${complex_data.endoffset}); }
    | '[' map_entries ']'							{ $$ = new ConcreteMapNode($map_entries, source, ${complex_data.offset}, ${complex_data.endoffset}); }
 ;

map_entries (java.util.@HashMap<String,ExpressionNode>) ::=
	identifier ':' conditional_expression						
													{ $$ = new java.util.@HashMap(); $map_entries.put($identifier, $conditional_expression); }
	| map_entries ',' identifier ':' conditional_expression
													{ $map_entries#0.put($identifier, $conditional_expression); }
;

bcon (Boolean) ::= 
	Ltrue 											{ $$ = Boolean.TRUE; }
	| Lfalse										{ $$ = Boolean.FALSE; }
;

unary_expression (ExpressionNode) ::=
	primary_expression
	| '!' unary_expression							{ $$ = new UnaryExpression(UnaryExpression.NOT, $unary_expression#1, source, ${unary_expression[0].offset}, ${unary_expression[0].endoffset}); }
	| '-' unary_expression							{ $$ = new UnaryExpression(UnaryExpression.MINUS, $unary_expression#1, source, ${unary_expression[0].offset}, ${unary_expression[0].endoffset}); }
;

mult_expression (ExpressionNode) ::=
	unary_expression
	| mult_expression '*' unary_expression			{ $$ = new ArithmeticNode(ArithmeticNode.MULT, $mult_expression#0, $unary_expression, source, ${mult_expression[0].offset}, ${mult_expression[0].endoffset}); }
	| mult_expression '/' unary_expression			{ $$ = new ArithmeticNode(ArithmeticNode.DIV, $mult_expression#0, $unary_expression, source, ${mult_expression[0].offset}, ${mult_expression[0].endoffset}); }
	| mult_expression '%' unary_expression			{ $$ = new ArithmeticNode(ArithmeticNode.REM, $mult_expression#0, $unary_expression, source, ${mult_expression[0].offset}, ${mult_expression[0].endoffset}); }
;

additive_expression (ExpressionNode) ::=
	mult_expression
	| additive_expression '+' mult_expression		{ $$ = new ArithmeticNode(ArithmeticNode.PLUS, $additive_expression#0, $mult_expression, source, ${additive_expression[0].offset}, ${additive_expression[0].endoffset}); }
	| additive_expression '-' mult_expression		{ $$ = new ArithmeticNode(ArithmeticNode.MINUS, $additive_expression#0, $mult_expression, source, ${additive_expression[0].offset}, ${additive_expression[0].endoffset}); }
;


relational_expression (ExpressionNode) ::=
      additive_expression
    | relational_expression '<' additive_expression	{ $$ = new ConditionalNode(ConditionalNode.LT, $relational_expression#0, $additive_expression, source, ${relational_expression[0].offset}, ${relational_expression[0].endoffset}); }
    | relational_expression '>' additive_expression	{ $$ = new ConditionalNode(ConditionalNode.GT, $relational_expression#0, $additive_expression, source, ${relational_expression[0].offset}, ${relational_expression[0].endoffset}); }
    | relational_expression '<=' additive_expression { $$ = new ConditionalNode(ConditionalNode.LE, $relational_expression#0, $additive_expression, source, ${relational_expression[0].offset}, ${relational_expression[0].endoffset}); }
    | relational_expression '>=' additive_expression { $$ = new ConditionalNode(ConditionalNode.GE, $relational_expression#0, $additive_expression, source, ${relational_expression[0].offset}, ${relational_expression[0].endoffset}); }
;

equality_expression (ExpressionNode) ::=
      relational_expression
    | equality_expression '==' relational_expression { $$ = new ConditionalNode(ConditionalNode.EQ, $equality_expression#0, $relational_expression, source, ${equality_expression[0].offset}, ${equality_expression[0].endoffset}); }
    | equality_expression '!=' relational_expression { $$ = new ConditionalNode(ConditionalNode.NE, $equality_expression#0, $relational_expression, source, ${equality_expression[0].offset}, ${equality_expression[0].endoffset}); }
;

conditional_and_expression (ExpressionNode) ::=
      equality_expression
    | conditional_and_expression '&&' equality_expression { $$ = new ConditionalNode(ConditionalNode.AND, $conditional_and_expression#0, $equality_expression, source, ${conditional_and_expression[0].offset}, ${conditional_and_expression[0].endoffset}); }
;

conditional_or_expression (ExpressionNode) ::=
      conditional_and_expression
    | conditional_or_expression '||' conditional_and_expression	{ $$ = new ConditionalNode(ConditionalNode.OR, $conditional_or_expression#0, $conditional_and_expression, source, ${conditional_or_expression[0].offset}, ${conditional_or_expression[0].endoffset}); }
;

conditional_expression (ExpressionNode) ::=
    conditional_or_expression
  | conditional_or_expression '?' conditional_expression ':' conditional_expression
  													{ $$ = new TriplexNode($conditional_or_expression, $conditional_expression#1, $conditional_expression#2, source, ${left().offset}, ${left().endoffset}); }
;

assignment_expression (ExpressionNode) ::=
	conditional_expression
  | identifier '=' conditional_expression			{ $$ = new AssignNode($identifier, $conditional_expression, source, ${left().offset}, ${left().endoffset}); }
;

expression (ExpressionNode) ::=
	assignment_expression
  | expression ',' assignment_expression			{ $$ = new CommaNode($expression#1, $assignment_expression, source, ${left().offset}, ${left().endoffset}); }
;

expression_list (ArrayList) ::=
	conditional_expression							{ $$ = new ArrayList(); $expression_list.add($conditional_expression); }
	| expression_list ',' conditional_expression	{ $expression_list#0.add($conditional_expression); }
;

body (TemplateNode) ::=
	instructions
						{
							$$ = new TemplateNode("inline", null, templatePackage, source, ${body.offset}, ${body.endoffset});
							$body.setInstructions($instructions);
							entities.add($body);
						}
;

%input input body;

##################################################################################
%%
${template java.imports}
import java.util.ArrayList;
import java.util.List;
${end}

${template java_tree.createParser-}
${call base-}
parser.source = source;
${end}

${template java.classcode}
${call base-}
net.sf.lapg.templates.ast.AstTree.@TextSource source;

private ArrayList<net.sf.lapg.templates.api.@IBundleEntity> entities;
private String templatePackage;

private int killEnds = -1;

private int rawText(int start, final int end) {
	char[] buff = source.getContents();
	if( killEnds == start ) {
		while( start < end && (buff[start] == '\t' || buff[start] == ' ') )
			start++;

		if( start < end && buff[start] == '\r' )
			start++;

		if( start < end && buff[start] == '\n' )
			start++;
	}
	return start;
}

private void checkIsSpace(int start, int end, int line) {
	String val = source.getText(rawText(start,end),end).trim();
	if( val.length() > 0 )
		reporter.error(start, end, line, "Unknown text ignored: `"+val+"`");
}

private void applyElse(CompoundNode node, ElseIfNode elseNode, int offset, int endoffset, int line) {
	if (elseNode == null ) {
		return;
	}
	if (node instanceof IfNode) {
		((IfNode)node).applyElse(elseNode);
	} else {
		reporter.error(offset, endoffset, line, "Unknown else node, instructions skipped");
	}
}

private ExpressionNode createMapCollect(ExpressionNode context, String instruction, String varName, ExpressionNode key, ExpressionNode value, net.sf.lapg.templates.ast.AstTree.@TextSource source, int offset, int endoffset, int line) {
	if(!instruction.equals("collect")) {
		reporter.error(offset, endoffset, line, "unknown collection processing instruction: " + instruction);
		return new ErrorNode(source, offset, endoffset);
	}
	return new CollectMapNode(context, varName, key, value, source, offset, endoffset);
}

private ExpressionNode createCollectionProcessor(ExpressionNode context, String instruction, String varName, ExpressionNode foreachExpr, net.sf.lapg.templates.ast.AstTree.@TextSource source, int offset, int endoffset, int line) {
	char first = instruction.charAt(0);
	int kind = 0;
	switch(first) {
	case 'c':
		if(instruction.equals("collect")) {
			kind = CollectionProcessorNode.COLLECT;
		} else if(instruction.equals("collectUnique")) {
			kind = CollectionProcessorNode.COLLECTUNIQUE;
		}
		break;
	case 'r':
		if(instruction.equals("reject")) {
			kind = CollectionProcessorNode.REJECT;
		}
		break;
	case 's':
		if(instruction.equals("select")) {
			kind = CollectionProcessorNode.SELECT;
		} else if(instruction.equals("sort")) {
			kind = CollectionProcessorNode.SORT;
		}
		break;
	case 'f':
		if(instruction.equals("forAll")) {
			kind = CollectionProcessorNode.FORALL;
		}
		break;
	case 'e':
		if(instruction.equals("exists")) {
			kind = CollectionProcessorNode.EXISTS;
		}
		break;
	}
	if(kind == 0) {
		reporter.error(offset, endoffset, line, "unknown collection processing instruction: " + instruction);
		return new ErrorNode(source, offset, endoffset);
	}
	return new CollectionProcessorNode(context, kind, varName, foreachExpr, source, offset, endoffset);
}

private Node createEscapedId(String escid, int offset, int endoffset) {
	int sharp = escid.indexOf('#');
	if( sharp >= 0 ) {
		Integer index = new Integer(escid.substring(sharp+1));
		escid = escid.substring(0, sharp);
		return new IndexNode(new SelectNode(null,escid,source,offset,endoffset), new LiteralNode(index,source,offset,endoffset),source,offset,endoffset);
	
	} else {
		return new SelectNode(null,escid,source,offset,endoffset);
	}
}

private void skipSpaces(int offset) {
	killEnds = offset+1;
}

private void checkFqn(String templateName, int offset, int endoffset, int line) {
	if( templateName.indexOf('.') >= 0 && templatePackage != null) {
		reporter.error(offset, endoffset, line, "template name should be simple identifier");
	}
}

public boolean parse(net.sf.lapg.templates.ast.AstTree.@TextSource source, String templatePackage) {
	this.templatePackage = templatePackage;
	this.entities = new ArrayList<net.sf.lapg.templates.api.@IBundleEntity>();
	this.source = source; 
	try {
		AstLexer lexer = new AstLexer(source.getStream(), reporter);
		parseInput(lexer);
		return true;
	} catch( ParseException ex ) {
		return false;
	} catch( IOException ex ) {
		return false;
	}
}

public boolean parseBody(net.sf.lapg.templates.ast.AstTree.@TextSource source, String templatePackage) {
	this.templatePackage = templatePackage;
	this.entities = new ArrayList<net.sf.lapg.templates.api.@IBundleEntity>();
	this.source = source; 
	try {
		AstLexer lexer = new AstLexer(source.getStream(), reporter);
		parseBody(lexer);
		return true;
	} catch( ParseException ex ) {
		return false;
	} catch( IOException ex ) {
		return false;
	}
}

public net.sf.lapg.templates.api.@IBundleEntity[] getResult() {
	return entities.toArray(new net.sf.lapg.templates.api.@IBundleEntity[entities.size()]);
}
${end}