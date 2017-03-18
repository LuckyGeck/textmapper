// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/inspirer/textmapper/tm-go/parsers/test"
	"github.com/inspirer/textmapper/tm-go/parsers/test/selector"
)

type Node interface {
	Type() test.NodeType
	// Child returns the first child node that matches the selector.
	Child(sel selector.Selector) Node
	Children(sel selector.Selector) []Node
	// Next returns the first element among the following siblings that matches the selector.
	Next(sel selector.Selector) Node
	// NextAll returns all following siblings of the node that match the selector.
	NextAll(sel selector.Selector) []Node
	Text() string
}

// Interfaces.

type TestNode interface {
	Node
	testNodeNode()
}

type Token struct {
	Node
}

// All types implement TestNode.
func (Block) testNodeNode()    {}
func (Decl1) testNodeNode()    {}
func (Decl2) testNodeNode()    {}
func (Negation) testNodeNode() {}
func (Test) testNodeNode()     {}
func (Token) testNodeNode()    {}

type Declaration interface {
	TestNode
	declarationNode()
}

// declarationNode() ensures that only the following types can be
// assigned to Declaration.
//
func (Block) declarationNode() {}
func (Decl1) declarationNode() {}
func (Decl2) declarationNode() {}

// Types.

type Block struct {
	Node
}

func (n Block) Negation() *Negation {
	if child := n.Child(selector.Negation); child != nil {
		return &Negation{child}
	}
	return nil
}

func (n Block) Declaration() []Declaration {
	nodes := n.Children(selector.Declaration)
	var result []Declaration = make([]Declaration, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, ToTestNode(node).(Declaration))
	}
	return result
}

type Decl1 struct {
	Node
}

func (n Decl1) Identifier() []Token {
	nodes := n.Children(selector.Identifier)
	var result []Token = make([]Token, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, Token{node})
	}
	return result
}

type Decl2 struct {
	Node
}

type Negation struct {
	Node
}

type Test struct {
	Node
}

func (n Test) Declaration() []Declaration {
	nodes := n.Children(selector.Declaration)
	var result []Declaration = make([]Declaration, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, ToTestNode(node).(Declaration))
	}
	return result
}
