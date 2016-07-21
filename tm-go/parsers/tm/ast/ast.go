// generated by Textmapper; DO NOT EDIT

package ast

type Node interface {
	Start() int
	End() int
}

type Pos struct {
	Offset, Endoffset int
}

func (p Pos) Start() int { return p.Offset }
func (p Pos) End() int   { return p.Endoffset }

type LexemeAttribute int

const (
	LexemeAttribute_LSOFT LexemeAttribute = iota
	LexemeAttribute_LCLASS
	LexemeAttribute_LSPACE
	LexemeAttribute_LLAYOUT
)

type Assoc int

const (
	Assoc_LLEFT Assoc = iota
	Assoc_LRIGHT
	Assoc_LNONASSOC
)

type ParamModifier int

const (
	ParamModifier_LEXPLICIT ParamModifier = iota
	ParamModifier_LGLOBAL
	ParamModifier_LLOOKAHEAD
)

type ParamType int

const (
	ParamType_LFLAG ParamType = iota
	ParamType_LPARAM
)

type LexerPart interface {
	lexerPart()
}

type GrammarPart interface {
	grammarPart()
}

type NontermType interface {
	nontermType()
}

type NontermTypeAST struct {
	Reference *Symref
	Pos
}

func (*NontermTypeAST) nontermType() {}

type NontermTypeHint struct {
	Inline     bool
	Kind       NontermTypeHint_KindKind
	Name       *Identifier
	Implements []*Symref
	Pos
}

func (*NontermTypeHint) nontermType() {}

type NontermTypeHint_KindKind int

const (
	NontermTypeHint_LCLASS NontermTypeHint_KindKind = iota
	NontermTypeHint_LVOID
	NontermTypeHint_LINTERFACE
)

type NontermTypeRaw struct {
	TypeText string
	Pos
}

func (*NontermTypeRaw) nontermType() {}

type RhsPart interface {
	rhsPart()
}

type SetExpression interface {
	setExpression()
}

type SetBinary struct {
	Left  SetExpression
	Kind  SetBinary_KindKind
	Right SetExpression
	Pos
}

func (*SetBinary) setExpression() {}

type SetBinary_KindKind int

const (
	SetBinary_OR SetBinary_KindKind = iota
	SetBinary_AND
)

type NontermParam interface {
	nontermParam()
}

type InlineParameter struct {
	ParamType  string
	Name       *Identifier
	ParamValue ParamValue
	Pos
}

func (*InlineParameter) nontermParam() {}

type ParamValue interface {
	paramValue()
}

type PredicateExpression interface {
	predicateExpression()
}

type PredicateBinary struct {
	Left  PredicateExpression
	Kind  PredicateBinary_KindKind
	Right PredicateExpression
	Pos
}

func (*PredicateBinary) predicateExpression() {}

type PredicateBinary_KindKind int

const (
	PredicateBinary_ANDAND PredicateBinary_KindKind = iota
	PredicateBinary_OROR
)

type Expression interface {
	expression()
}

type Instance struct {
	ClassName *Name
	Entries   []*MapEntry
	Pos
}

func (*Instance) expression() {}

type Array struct {
	Content []Expression
	Pos
}

func (*Array) expression() {}

type Input struct {
	Header  *Header
	Imports []*Import
	Options []*Option
	Lexer   []LexerPart
	Parser  []GrammarPart
	Pos
}

type Header struct {
	Name             *Name
	Target           *Name
	ParsingAlgorithm *ParsingAlgorithm
	Pos
}

type ParsingAlgorithm struct {
	La int
	Pos
}

type Import struct {
	Alias string
	File  string
	Pos
}

type Option struct {
	Key           string
	Value         Expression
	SyntaxProblem *SyntaxProblem
	Pos
}

type Identifier struct {
	ID string
	Pos
}

type Symref struct {
	Name string
	Args *SymrefArgs
	Pos
}

func (*Symref) expression() {}
func (*Symref) paramValue() {}

type Pattern struct {
	REGEXP string
	Pos
}

type NamedPattern struct {
	Name    string
	Pattern *Pattern
	Pos
}

func (*NamedPattern) lexerPart() {}

type Lexeme struct {
	Name       *Identifier
	Type       string
	Pattern    *Pattern
	Transition *Stateref
	Priority   int
	Attrs      *LexemeAttrs
	Command    *Command
	Pos
}

func (*Lexeme) lexerPart() {}

type LexemeAttrs struct {
	Kind LexemeAttribute
	Pos
}

type StateSelector struct {
	States []*LexerState
	Pos
}

func (*StateSelector) lexerPart() {}

type Stateref struct {
	Name string
	Pos
}

type LexerState struct {
	Name              *Identifier
	DefaultTransition *Stateref
	Pos
}

type Nonterm struct {
	Annotations *Annotations
	Name        *Identifier
	Params      *NontermParams
	Type        NontermType
	Rules       []*Rule0
	Pos
}

func (*Nonterm) grammarPart() {}

type Inputref struct {
	Reference *Symref
	Noeoi     bool
	Pos
}

type Rule0 struct {
	Predicate PredicateExpression
	Prefix    *RhsPrefix
	List      []RhsPart
	Action    *RuleAction
	Suffix    *RhsSuffix
	Error     *SyntaxProblem
	Pos
}

type RhsPrefix struct {
	Annotations *Annotations
	Pos
}

type RhsSuffix struct {
	Kind   RhsSuffix_KindKind
	Symref *Symref
	Pos
}

type RhsSuffix_KindKind int

const (
	RhsSuffix_LPREC RhsSuffix_KindKind = iota
	RhsSuffix_LSHIFT
)

type RuleAction struct {
	Action    *Identifier
	Parameter string
	Pos
}

type Annotations struct {
	Annotations []*Annotation
	Pos
}

type Annotation struct {
	Name          string
	Expression    Expression
	SyntaxProblem *SyntaxProblem
	Pos
}

type NontermParams struct {
	List []NontermParam
	Pos
}

type ParamRef struct {
	Ref *Identifier
	Pos
}

func (*ParamRef) nontermParam() {}

type SymrefArgs struct {
	ArgList []*Argument
	Pos
}

type Argument struct {
	Name *ParamRef
	Val  ParamValue
	Bool Argument_BoolKind
	Pos
}

type Argument_BoolKind int

const (
	Argument_PLUS Argument_BoolKind = iota
	Argument_TILDE
)

type MapEntry struct {
	Name  string
	Value Expression
	Pos
}

type Literal struct {
	Value interface{}
	Pos
}

func (*Literal) paramValue() {}
func (*Literal) expression() {}

type Name struct {
	QualifiedId string
	Pos
}

type Command struct {
	Pos
}

func (*Command) rhsPart() {}

type SyntaxProblem struct {
	Pos
}

func (*SyntaxProblem) lexerPart() {}
func (*SyntaxProblem) grammarPart() {}
func (*SyntaxProblem) rhsPart() {}
func (*SyntaxProblem) expression() {}

type DirectiveBrackets struct {
	Opening *Symref
	Closing *Symref
	Pos
}

func (*DirectiveBrackets) lexerPart() {}

type TemplateParam struct {
	Modifier   ParamModifier
	ParamType  ParamType
	Name       *Identifier
	ParamValue ParamValue
	Pos
}

func (*TemplateParam) grammarPart() {}

type DirectivePrio struct {
	Assoc   Assoc
	Symbols []*Symref
	Pos
}

func (*DirectivePrio) grammarPart() {}

type DirectiveInput struct {
	InputRefs []*Inputref
	Pos
}

func (*DirectiveInput) grammarPart() {}

type DirectiveAssert struct {
	Kind   DirectiveAssert_KindKind
	RhsSet *RhsSet
	Pos
}

func (*DirectiveAssert) grammarPart() {}

type DirectiveAssert_KindKind int

const (
	DirectiveAssert_LEMPTY DirectiveAssert_KindKind = iota
	DirectiveAssert_LNONEMPTY
)

type DirectiveSet struct {
	Name   string
	RhsSet *RhsSet
	Pos
}

func (*DirectiveSet) grammarPart() {}

type RhsAnnotated struct {
	Annotations *Annotations
	Inner       RhsPart
	Pos
}

func (*RhsAnnotated) rhsPart() {}

type RhsAssignment struct {
	Id       *Identifier
	Addition bool
	Inner    RhsPart
	Pos
}

func (*RhsAssignment) rhsPart() {}

type RhsQuantifier struct {
	Inner      RhsPart
	Quantifier RhsQuantifier_QuantifierKind
	Pos
}

func (*RhsQuantifier) rhsPart() {}

type RhsQuantifier_QuantifierKind int

const (
	RhsQuantifier_QUEST RhsQuantifier_QuantifierKind = iota
	RhsQuantifier_PLUS
	RhsQuantifier_MULT
)

type RhsCast struct {
	Inner  RhsPart
	Target *Symref
	Pos
}

func (*RhsCast) rhsPart() {}

type RhsAsLiteral struct {
	Inner   RhsPart
	Literal *Literal
	Pos
}

func (*RhsAsLiteral) rhsPart() {}

type RhsUnordered struct {
	Left  RhsPart
	Right RhsPart
	Pos
}

func (*RhsUnordered) rhsPart() {}

type RhsClass struct {
	Identifier *Identifier
	Inner      RhsPart
	Pos
}

func (*RhsClass) rhsPart() {}

type RhsSymbol struct {
	Reference *Symref
	Pos
}

func (*RhsSymbol) rhsPart() {}

type RhsNested struct {
	Rules []*Rule0
	Pos
}

func (*RhsNested) rhsPart() {}

type RhsList struct {
	RuleParts  []RhsPart
	Separator  []*Symref
	AtLeastOne bool
	Pos
}

func (*RhsList) rhsPart() {}

type RhsIgnored struct {
	Rules []*Rule0
	Pos
}

func (*RhsIgnored) rhsPart() {}

type RhsSet struct {
	Expr SetExpression
	Pos
}

func (*RhsSet) rhsPart() {}

type SetSymbol struct {
	Operator string
	Symbol   *Symref
	Pos
}

func (*SetSymbol) setExpression() {}

type SetCompound struct {
	Inner SetExpression
	Pos
}

func (*SetCompound) setExpression() {}

type SetComplement struct {
	Inner SetExpression
	Pos
}

func (*SetComplement) setExpression() {}

type BoolPredicate struct {
	Negated  bool
	ParamRef *ParamRef
	Pos
}

func (*BoolPredicate) predicateExpression() {}

type ComparePredicate struct {
	ParamRef *ParamRef
	Kind     ComparePredicate_KindKind
	Literal  *Literal
	Pos
}

func (*ComparePredicate) predicateExpression() {}

type ComparePredicate_KindKind int

const (
	ComparePredicate_ASSIGNASSIGN ComparePredicate_KindKind = iota
	ComparePredicate_EXCLASSIGN
)
