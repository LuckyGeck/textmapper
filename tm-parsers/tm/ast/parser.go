package ast

import (
	"errors"

	"github.com/inspirer/textmapper/tm-parsers/tm"
)

// Parse parses a given utf-8 content into an AST.
func Parse(path, content string, eh tm.ErrorHandler) (*Tree, error) {
	var l tm.Lexer
	l.Init(content)
	var p tm.Parser
	b := newBuilder(path, content)
	p.Init(eh, b.addNode)
	err := p.Parse(&l)
	if err != nil {
		return nil, err
	}
	return b.build()
}

type builder struct {
	tree  *Tree
	stack []*Node
	err   error
}

func newBuilder(path, content string) *builder {
	return &builder{
		tree:  newTree(path, content),
		stack: make([]*Node, 0, 512),
	}
}

func (b *builder) addNode(t tm.NodeType, offset, endoffset int) {
	if t == tm.File {
		offset, endoffset = 0, len(b.tree.content)
	}

	start := len(b.stack)
	end := start
	for start > 0 && b.stack[start-1].offset >= offset {
		start--
		if b.stack[start].offset >= endoffset {
			end--
		}
	}
	out := &Node{
		tree:      b.tree,
		t:         t,
		offset:    offset,
		endoffset: endoffset,
	}
	if start < end {
		out.firstChild = b.stack[start]
		var prev *Node
		for i := end - 1; i >= start; i-- {
			n := b.stack[i]
			n.parent = out
			n.next = prev
			prev = n
		}
	}
	if end == len(b.stack) {
		b.stack = append(b.stack[:start], out)
	} else if start < end {
		b.stack[start] = out
		l := copy(b.stack[start+1:], b.stack[end:])
		b.stack = b.stack[:start+1+l]
	} else {
		b.stack = append(b.stack, nil)
		copy(b.stack[start+1:], b.stack[start:])
		b.stack[start] = out
	}
}

var errNumRoots = errors.New("exactly one root node is expected")

func (b *builder) build() (*Tree, error) {
	if b.err != nil {
		return nil, b.err
	}
	if len(b.stack) != 1 {
		return nil, errNumRoots
	}
	b.tree.root = b.stack[0]
	return b.tree, nil
}
