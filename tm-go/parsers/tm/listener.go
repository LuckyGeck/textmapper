// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	Identifier NodeType = iota + 1
	IntegerLiteral
	StringLiteral
	BooleanLiteral
	Pattern
	Name // (identifier)+
	Command
	SyntaxProblem
	Input    // header imports=(import_)* options=(option)* lexer=(lexer_part)+ parser=grammar_parts?
	Header   // name target=name?
	Import   // alias=identifier? path=string_literal
	KeyValue // key=identifier value=expression
	Symref   // name=identifier args=symref_args?
	RawType
	NamedPattern         // name=identifier pattern
	StartConditionsScope // start_conditions (lexer_part)+
	StartConditions      // (stateref)*
	Lexeme               // start_conditions? name=identifier rawType? pattern? priority=integer_literal? attrs=lexeme_attrs? command?
	LexemeAttrs          // lexeme_attribute
	LexemeAttribute
	DirectiveBrackets // opening=symref closing=symref
	InclusiveStates   // states=(lexer_state)+
	ExclusiveStates   // states=(lexer_state)+
	StateSelector     // states=(stateref)+
	Stateref          // name=identifier
	LexerState        // name=identifier
	GrammarParts      // grammar_parts? grammar_part
	Nonterm           // annotations? name=identifier params=nonterm_params? type=nonterm_type? (rule0)+
	SubType           // reference=symref
	InterfaceType
	VoidType
	Assoc
	ParamModifier
	TemplateParam     // modifier=param_modifier? param_type name=identifier param_value?
	DirectivePrio     // assoc symbols=references
	DirectiveInput    // inputRefs=(inputref)+
	DirectiveAssert   // rhsSet
	DirectiveSet      // name=identifier rhsSet
	Inputref          // reference=symref
	References        // references? symref
	Rule              // predicate? (rhsPart)* ruleAction? rhsSuffix?
	Predicate         // predicate_expression
	RhsSuffix         // symref
	RuleAction        // action=identifier kind=identifier?
	RhsAnnotated      // annotations inner=rhsPart
	RhsAssignment     // id=identifier inner=rhsPart
	RhsPlusAssignment // id=identifier inner=rhsPart
	RhsOptional       // inner=rhsPart
	RhsCast           // inner=rhsPart target=symref
	ListSeparator     // separator_=references
	RhsSymbol         // reference=symref
	RhsNested         // (rule0)+
	RhsPlusList       // ruleParts=(rhsPart)+ listSeparator
	RhsStarList       // ruleParts=(rhsPart)+ listSeparator
	RhsQuantifier     // inner=rhsPart
	RhsIgnored        // (rule0)+
	RhsPrimary        // rhsSet
	RhsSet            // expr=setExpression
	SetSymbol         // operator=identifier? symbol=symref
	SetCompound       // inner=setExpression
	SetComplement     // inner=setExpression
	SetOr             // left=setExpression right=setExpression
	SetAnd            // left=setExpression right=setExpression
	Annotations       // (annotation)+
	AnnotationImpl    // name=identifier expression?
	NontermParams     // list=(nonterm_param)+
	InlineParameter   // param_type=identifier name=identifier param_value?
	ParamRef          // identifier
	SymrefArgs        // arg_list=(argument)*
	ArgumentImpl      // name=param_ref val=param_value?
	ArgumentTrue      // name=param_ref
	ArgumentFalse     // name=param_ref
	ParamType
	PredicateNot   // param_ref
	PredicateEq    // param_ref literal
	PredicateNotEq // param_ref literal
	PredicateAnd   // left=predicate_expression right=predicate_expression
	PredicateOr    // left=predicate_expression right=predicate_expression
	Array          // (expression)*
	InvalidToken
	MultilineComment
	Comment
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"Identifier",
	"IntegerLiteral",
	"StringLiteral",
	"BooleanLiteral",
	"Pattern",
	"Name",
	"Command",
	"SyntaxProblem",
	"Input",
	"Header",
	"Import",
	"KeyValue",
	"Symref",
	"RawType",
	"NamedPattern",
	"StartConditionsScope",
	"StartConditions",
	"Lexeme",
	"LexemeAttrs",
	"LexemeAttribute",
	"DirectiveBrackets",
	"InclusiveStates",
	"ExclusiveStates",
	"StateSelector",
	"Stateref",
	"LexerState",
	"GrammarParts",
	"Nonterm",
	"SubType",
	"InterfaceType",
	"VoidType",
	"Assoc",
	"ParamModifier",
	"TemplateParam",
	"DirectivePrio",
	"DirectiveInput",
	"DirectiveAssert",
	"DirectiveSet",
	"Inputref",
	"References",
	"Rule",
	"Predicate",
	"RhsSuffix",
	"RuleAction",
	"RhsAnnotated",
	"RhsAssignment",
	"RhsPlusAssignment",
	"RhsOptional",
	"RhsCast",
	"ListSeparator",
	"RhsSymbol",
	"RhsNested",
	"RhsPlusList",
	"RhsStarList",
	"RhsQuantifier",
	"RhsIgnored",
	"RhsPrimary",
	"RhsSet",
	"SetSymbol",
	"SetCompound",
	"SetComplement",
	"SetOr",
	"SetAnd",
	"Annotations",
	"AnnotationImpl",
	"NontermParams",
	"InlineParameter",
	"ParamRef",
	"SymrefArgs",
	"ArgumentImpl",
	"ArgumentTrue",
	"ArgumentFalse",
	"ParamType",
	"PredicateNot",
	"PredicateEq",
	"PredicateNotEq",
	"PredicateAnd",
	"PredicateOr",
	"Array",
	"InvalidToken",
	"MultilineComment",
	"Comment",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var Annotation = []NodeType{
	AnnotationImpl,
	SyntaxProblem,
}

var Argument = []NodeType{
	ArgumentFalse,
	ArgumentImpl,
	ArgumentTrue,
}

var Expression = []NodeType{
	Array,
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
	Symref,
	SyntaxProblem,
}

var GrammarPart = []NodeType{
	DirectiveAssert,
	DirectiveInput,
	DirectivePrio,
	DirectiveSet,
	Nonterm,
	TemplateParam,
}

var LexerPart = []NodeType{
	DirectiveBrackets,
	ExclusiveStates,
	InclusiveStates,
	Lexeme,
	NamedPattern,
	StartConditionsScope,
	StateSelector,
}

var Literal = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
}

var NontermParam = []NodeType{
	InlineParameter,
	ParamRef,
}

var NontermType = []NodeType{
	InterfaceType,
	RawType,
	SubType,
	VoidType,
}

var Option = []NodeType{
	KeyValue,
	SyntaxProblem,
}

var ParamValue = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	ParamRef,
	StringLiteral,
}

var PredicateExpression = []NodeType{
	ParamRef,
	PredicateAnd,
	PredicateEq,
	PredicateNot,
	PredicateNotEq,
	PredicateOr,
}

var RhsPart = []NodeType{
	Command,
	RhsAnnotated,
	RhsAssignment,
	RhsCast,
	RhsIgnored,
	RhsNested,
	RhsOptional,
	RhsPlusAssignment,
	RhsPlusList,
	RhsPrimary,
	RhsQuantifier,
	RhsStarList,
	RhsSymbol,
}

var Rule0 = []NodeType{
	Rule,
	SyntaxProblem,
}

var SetExpression = []NodeType{
	SetAnd,
	SetComplement,
	SetCompound,
	SetOr,
	SetSymbol,
}

var ruleNodeType = [...]NodeType{
	Identifier,           // identifier : ID
	Identifier,           // identifier : 'brackets'
	Identifier,           // identifier : 'inline'
	Identifier,           // identifier : 'prec'
	Identifier,           // identifier : 'shift'
	Identifier,           // identifier : 'returns'
	Identifier,           // identifier : 'input'
	Identifier,           // identifier : 'left'
	Identifier,           // identifier : 'right'
	Identifier,           // identifier : 'nonassoc'
	Identifier,           // identifier : 'generate'
	Identifier,           // identifier : 'assert'
	Identifier,           // identifier : 'empty'
	Identifier,           // identifier : 'nonempty'
	Identifier,           // identifier : 'global'
	Identifier,           // identifier : 'explicit'
	Identifier,           // identifier : 'lookahead'
	Identifier,           // identifier : 'param'
	Identifier,           // identifier : 'flag'
	Identifier,           // identifier : 'no-eoi'
	Identifier,           // identifier : 's'
	Identifier,           // identifier : 'x'
	Identifier,           // identifier : 'soft'
	Identifier,           // identifier : 'class'
	Identifier,           // identifier : 'interface'
	Identifier,           // identifier : 'void'
	Identifier,           // identifier : 'space'
	Identifier,           // identifier : 'layout'
	Identifier,           // identifier : 'language'
	Identifier,           // identifier : 'lalr'
	Identifier,           // identifier : 'lexer'
	Identifier,           // identifier : 'parser'
	Identifier,           // identifier_Kw : ID
	Identifier,           // identifier_Kw : 'brackets'
	Identifier,           // identifier_Kw : 'inline'
	Identifier,           // identifier_Kw : 'prec'
	Identifier,           // identifier_Kw : 'shift'
	Identifier,           // identifier_Kw : 'returns'
	Identifier,           // identifier_Kw : 'input'
	Identifier,           // identifier_Kw : 'left'
	Identifier,           // identifier_Kw : 'right'
	Identifier,           // identifier_Kw : 'nonassoc'
	Identifier,           // identifier_Kw : 'generate'
	Identifier,           // identifier_Kw : 'assert'
	Identifier,           // identifier_Kw : 'empty'
	Identifier,           // identifier_Kw : 'nonempty'
	Identifier,           // identifier_Kw : 'global'
	Identifier,           // identifier_Kw : 'explicit'
	Identifier,           // identifier_Kw : 'lookahead'
	Identifier,           // identifier_Kw : 'param'
	Identifier,           // identifier_Kw : 'flag'
	Identifier,           // identifier_Kw : 'no-eoi'
	Identifier,           // identifier_Kw : 's'
	Identifier,           // identifier_Kw : 'x'
	Identifier,           // identifier_Kw : 'soft'
	Identifier,           // identifier_Kw : 'class'
	Identifier,           // identifier_Kw : 'interface'
	Identifier,           // identifier_Kw : 'void'
	Identifier,           // identifier_Kw : 'space'
	Identifier,           // identifier_Kw : 'layout'
	Identifier,           // identifier_Kw : 'language'
	Identifier,           // identifier_Kw : 'lalr'
	Identifier,           // identifier_Kw : 'lexer'
	Identifier,           // identifier_Kw : 'parser'
	Identifier,           // identifier_Kw : 'true'
	Identifier,           // identifier_Kw : 'false'
	Identifier,           // identifier_Kw : 'separator'
	Identifier,           // identifier_Kw : 'as'
	Identifier,           // identifier_Kw : 'import'
	Identifier,           // identifier_Kw : 'set'
	IntegerLiteral,       // integer_literal : icon
	StringLiteral,        // string_literal : scon
	BooleanLiteral,       // boolean_literal : 'true'
	BooleanLiteral,       // boolean_literal : 'false'
	0,                    // literal : string_literal
	0,                    // literal : integer_literal
	0,                    // literal : boolean_literal
	Pattern,              // pattern : regexp
	0,                    // qualified_name : identifier
	0,                    // qualified_name : qualified_name '.' identifier_Kw
	Name,                 // name : qualified_name
	Command,              // command : code
	SyntaxProblem,        // syntax_problem : error
	0,                    // import__optlist : import__optlist import_
	0,                    // import__optlist :
	Input,                // input : header import__optlist option_optlist lexer_section parser_section
	Input,                // input : header import__optlist option_optlist lexer_section
	0,                    // option_optlist : option_optlist option
	0,                    // option_optlist :
	Header,               // header : 'language' name '(' name ')' ';'
	Header,               // header : 'language' name ';'
	0,                    // lexer_section : '::' 'lexer' lexer_parts
	0,                    // parser_section : '::' 'parser' grammar_parts
	Import,               // import_ : 'import' identifier string_literal ';'
	Import,               // import_ : 'import' string_literal ';'
	KeyValue,             // option : identifier '=' expression
	0,                    // option : syntax_problem
	Symref,               // symref : identifier
	Symref,               // symref_Args : identifier symref_args
	Symref,               // symref_Args : identifier
	RawType,              // rawType : code
	0,                    // lexer_parts : lexer_part
	0,                    // lexer_parts : lexer_parts lexer_part_OrSyntaxError
	0,                    // lexer_part : state_selector
	0,                    // lexer_part : named_pattern
	0,                    // lexer_part : lexeme
	0,                    // lexer_part : lexer_directive
	0,                    // lexer_part : start_conditions_scope
	0,                    // lexer_part_OrSyntaxError : state_selector
	0,                    // lexer_part_OrSyntaxError : named_pattern
	0,                    // lexer_part_OrSyntaxError : lexeme
	0,                    // lexer_part_OrSyntaxError : lexer_directive
	0,                    // lexer_part_OrSyntaxError : start_conditions_scope
	0,                    // lexer_part_OrSyntaxError : syntax_problem
	NamedPattern,         // named_pattern : identifier '=' pattern
	StartConditionsScope, // start_conditions_scope : start_conditions '{' lexer_parts '}'
	StartConditions,      // start_conditions : '<' '*' '>'
	StartConditions,      // start_conditions : '<' stateref_list_Comma_separated '>'
	0,                    // stateref_list_Comma_separated : stateref_list_Comma_separated ',' stateref
	0,                    // stateref_list_Comma_separated : stateref
	Lexeme,               // lexeme : start_conditions identifier rawTypeopt ':' pattern integer_literalopt lexeme_attrsopt commandopt
	Lexeme,               // lexeme : start_conditions identifier rawTypeopt ':'
	Lexeme,               // lexeme : identifier rawTypeopt ':' pattern integer_literalopt lexeme_attrsopt commandopt
	Lexeme,               // lexeme : identifier rawTypeopt ':'
	LexemeAttrs,          // lexeme_attrs : '(' lexeme_attribute ')'
	LexemeAttribute,      // lexeme_attribute : 'soft'
	LexemeAttribute,      // lexeme_attribute : 'class'
	LexemeAttribute,      // lexeme_attribute : 'space'
	LexemeAttribute,      // lexeme_attribute : 'layout'
	DirectiveBrackets,    // lexer_directive : '%' 'brackets' symref symref ';'
	InclusiveStates,      // lexer_directive : '%' 's' lexer_state_list_Comma_separated
	ExclusiveStates,      // lexer_directive : '%' 'x' lexer_state_list_Comma_separated
	0,                    // lexer_state_list_Comma_separated : lexer_state_list_Comma_separated ',' lexer_state
	0,                    // lexer_state_list_Comma_separated : lexer_state
	StateSelector,        // state_selector : '[' stateref_list_Comma_separated ']'
	Stateref,             // stateref : identifier
	LexerState,           // lexer_state : identifier
	GrammarParts,         // grammar_parts : grammar_part
	GrammarParts,         // grammar_parts : grammar_parts grammar_part_OrSyntaxError
	0,                    // grammar_part : nonterm
	0,                    // grammar_part : template_param
	0,                    // grammar_part : directive
	0,                    // grammar_part_OrSyntaxError : nonterm
	0,                    // grammar_part_OrSyntaxError : template_param
	0,                    // grammar_part_OrSyntaxError : directive
	0,                    // grammar_part_OrSyntaxError : syntax_problem
	Nonterm,              // nonterm : annotations identifier nonterm_params nonterm_type ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_params ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_type ':' rules ';'
	Nonterm,              // nonterm : annotations identifier ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params nonterm_type ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_type ':' rules ';'
	Nonterm,              // nonterm : identifier ':' rules ';'
	SubType,              // nonterm_type : 'returns' symref
	InterfaceType,        // nonterm_type : 'interface'
	VoidType,             // nonterm_type : 'void'
	0,                    // nonterm_type : rawType
	Assoc,                // assoc : 'left'
	Assoc,                // assoc : 'right'
	Assoc,                // assoc : 'nonassoc'
	ParamModifier,        // param_modifier : 'explicit'
	ParamModifier,        // param_modifier : 'global'
	ParamModifier,        // param_modifier : 'lookahead'
	TemplateParam,        // template_param : '%' param_modifier param_type identifier '=' param_value ';'
	TemplateParam,        // template_param : '%' param_modifier param_type identifier ';'
	TemplateParam,        // template_param : '%' param_type identifier '=' param_value ';'
	TemplateParam,        // template_param : '%' param_type identifier ';'
	DirectivePrio,        // directive : '%' assoc references ';'
	DirectiveInput,       // directive : '%' 'input' inputref_list_Comma_separated ';'
	DirectiveAssert,      // directive : '%' 'assert' 'empty' rhsSet ';'
	DirectiveAssert,      // directive : '%' 'assert' 'nonempty' rhsSet ';'
	DirectiveSet,         // directive : '%' 'generate' identifier '=' rhsSet ';'
	0,                    // inputref_list_Comma_separated : inputref_list_Comma_separated ',' inputref
	0,                    // inputref_list_Comma_separated : inputref
	Inputref,             // inputref : symref 'no-eoi'
	Inputref,             // inputref : symref
	References,           // references : symref
	References,           // references : references symref
	0,                    // rules : rule0
	0,                    // rules : rules '|' rule0
	Rule,                 // rule0 : predicate rhsParts ruleAction rhsSuffixopt
	Rule,                 // rule0 : predicate rhsParts rhsSuffixopt
	Rule,                 // rule0 : predicate ruleAction rhsSuffixopt
	Rule,                 // rule0 : predicate rhsSuffixopt
	Rule,                 // rule0 : rhsParts ruleAction rhsSuffixopt
	Rule,                 // rule0 : rhsParts rhsSuffixopt
	Rule,                 // rule0 : ruleAction rhsSuffixopt
	Rule,                 // rule0 : rhsSuffixopt
	0,                    // rule0 : syntax_problem
	Predicate,            // predicate : '[' predicate_expression ']'
	RhsSuffix,            // rhsSuffix : '%' 'prec' symref
	RhsSuffix,            // rhsSuffix : '%' 'shift' symref
	RuleAction,           // ruleAction : '->' identifier '/' identifier
	RuleAction,           // ruleAction : '->' identifier
	0,                    // rhsParts : rhsPart
	0,                    // rhsParts : rhsParts rhsPart_OrSyntaxError
	0,                    // rhsPart : rhsAnnotated
	0,                    // rhsPart : command
	0,                    // rhsPart_OrSyntaxError : rhsAnnotated
	0,                    // rhsPart_OrSyntaxError : command
	0,                    // rhsPart_OrSyntaxError : syntax_problem
	0,                    // rhsAnnotated : rhsAssignment
	RhsAnnotated,         // rhsAnnotated : annotations rhsAssignment
	0,                    // rhsAssignment : rhsOptional
	RhsAssignment,        // rhsAssignment : identifier '=' rhsOptional
	RhsPlusAssignment,    // rhsAssignment : identifier '+=' rhsOptional
	0,                    // rhsOptional : rhsCast
	RhsOptional,          // rhsOptional : rhsCast '?'
	0,                    // rhsCast : rhsPrimary
	RhsCast,              // rhsCast : rhsPrimary 'as' symref_Args
	ListSeparator,        // listSeparator : 'separator' references
	RhsSymbol,            // rhsPrimary : symref_Args
	RhsNested,            // rhsPrimary : '(' rules ')'
	RhsPlusList,          // rhsPrimary : '(' rhsParts listSeparator ')' '+'
	RhsStarList,          // rhsPrimary : '(' rhsParts listSeparator ')' '*'
	RhsQuantifier,        // rhsPrimary : rhsPrimary '+'
	RhsQuantifier,        // rhsPrimary : rhsPrimary '*'
	RhsIgnored,           // rhsPrimary : '$' '(' rules ')'
	RhsPrimary,           // rhsPrimary : rhsSet
	RhsSet,               // rhsSet : 'set' '(' setExpression ')'
	SetSymbol,            // setPrimary : identifier symref_Args
	SetSymbol,            // setPrimary : symref_Args
	SetCompound,          // setPrimary : '(' setExpression ')'
	SetComplement,        // setPrimary : '~' setPrimary
	0,                    // setExpression : setPrimary
	SetOr,                // setExpression : setExpression '|' setExpression
	SetAnd,               // setExpression : setExpression '&' setExpression
	0,                    // annotation_list : annotation_list annotation
	0,                    // annotation_list : annotation
	Annotations,          // annotations : annotation_list
	AnnotationImpl,       // annotation : '@' identifier '=' expression
	AnnotationImpl,       // annotation : '@' identifier
	0,                    // annotation : '@' syntax_problem
	0,                    // nonterm_param_list_Comma_separated : nonterm_param_list_Comma_separated ',' nonterm_param
	0,                    // nonterm_param_list_Comma_separated : nonterm_param
	NontermParams,        // nonterm_params : '<' nonterm_param_list_Comma_separated '>'
	0,                    // nonterm_param : param_ref
	InlineParameter,      // nonterm_param : identifier identifier '=' param_value
	InlineParameter,      // nonterm_param : identifier identifier
	ParamRef,             // param_ref : identifier
	0,                    // argument_list_Comma_separated : argument_list_Comma_separated ',' argument
	0,                    // argument_list_Comma_separated : argument
	0,                    // argument_list_Comma_separated_opt : argument_list_Comma_separated
	0,                    // argument_list_Comma_separated_opt :
	SymrefArgs,           // symref_args : '<' argument_list_Comma_separated_opt '>'
	ArgumentImpl,         // argument : param_ref ':' param_value
	ArgumentImpl,         // argument : param_ref
	ArgumentTrue,         // argument : '+' param_ref
	ArgumentFalse,        // argument : '~' param_ref
	ParamType,            // param_type : 'flag'
	ParamType,            // param_type : 'param'
	0,                    // param_value : literal
	0,                    // param_value : param_ref
	0,                    // predicate_primary : param_ref
	PredicateNot,         // predicate_primary : '!' param_ref
	PredicateEq,          // predicate_primary : param_ref '==' literal
	PredicateNotEq,       // predicate_primary : param_ref '!=' literal
	0,                    // predicate_expression : predicate_primary
	PredicateAnd,         // predicate_expression : predicate_expression '&&' predicate_expression
	PredicateOr,          // predicate_expression : predicate_expression '||' predicate_expression
	0,                    // expression : literal
	0,                    // expression : symref_Args
	Array,                // expression : '[' expression_list_Comma_separated_opt ']'
	0,                    // expression : syntax_problem
	0,                    // expression_list_Comma_separated : expression_list_Comma_separated ',' expression
	0,                    // expression_list_Comma_separated : expression
	0,                    // expression_list_Comma_separated_opt : expression_list_Comma_separated
	0,                    // expression_list_Comma_separated_opt :
	0,                    // rawTypeopt : rawType
	0,                    // rawTypeopt :
	0,                    // integer_literalopt : integer_literal
	0,                    // integer_literalopt :
	0,                    // lexeme_attrsopt : lexeme_attrs
	0,                    // lexeme_attrsopt :
	0,                    // commandopt : command
	0,                    // commandopt :
	0,                    // rhsSuffixopt : rhsSuffix
	0,                    // rhsSuffixopt :
}
