package grammar

import (
	"sort"

	"github.com/inspirer/textmapper/tm-go/lex"
	"github.com/inspirer/textmapper/tm-go/status"
)

// Names of common terminals with predefined meaning.
const (
	Eoi          = "eoi"
	Error        = "error"
	InvalidToken = "invalid_token"
)

// Symbol is a grammar symbol.
type Symbol struct {
	Index   int
	ID      string // unique identifier to be used in generated code
	Name    string
	Comment string
	Type    string
	Origin  status.SourceNode
}

// PrettyType returns a user-friendly representation of the symbol type.
func (sym *Symbol) PrettyType() string {
	if sym.Type != "" {
		return sym.Type
	}
	return "<no type>"
}

// Grammar is a fully resolved and compiled Textmapper grammar.
type Grammar struct {
	Name       string // lowercase
	TargetLang string
	Syms       []Symbol
	NumTokens  int
	*Lexer
	*Options
}

// Tokens returns all lexical tokens defined in the grammar.
func (g *Grammar) Tokens() []Symbol {
	return g.Syms[:g.NumTokens]
}

// SemanticAction is a piece of code that will be executed upon some event.
type SemanticAction struct {
	Action  int
	Code    string
	Space   bool // this is a space token
	Comment string
	Origin  status.SourceNode
}

// ClassAction resolves class terminals into more specific tokens (such as keywords).
type ClassAction struct {
	Action int
	Custom map[string]int // maps constant terminals back into actions
}

// Lexer is a model of a generated lexer.
type Lexer struct {
	StartConditions []string
	Tables          *lex.Tables
	ClassActions    []ClassAction
	Actions         []SemanticAction
	InvalidToken    int
	RuleToken       []int // maps actions into tokens; empty if the mapping is 1:1
}

// SpaceActions returns a sorted list of space-only actions.
func (l *Lexer) SpaceActions() []int {
	var ret []int
	for _, a := range l.Actions {
		if a.Code == "" && a.Space {
			ret = append(ret, a.Action)
		}
	}
	sort.Ints(ret)
	return ret
}

// CodeActions returns a list of non-empty code actions.
func (l *Lexer) CodeActions() []SemanticAction {
	var ret []SemanticAction
	for _, a := range l.Actions {
		if a.Code != "" {
			ret = append(ret, a)
		}
	}
	return ret
}

type Options struct {
	TokenLine       bool // true by default
	TokenLineOffset bool
}
