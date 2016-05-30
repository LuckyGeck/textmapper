package tm

import (
	"bytes"
	"unicode/utf8"
	"strconv"
)

const (
	State_initial = 0
	State_afterAt = 1
	State_afterAtID = 2
)

// ErrorHandler is called every time a lexer or parser is unable to process
// some part of the input.
type ErrorHandler func(line, offset, len int, msg string)

// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source []byte
	err    ErrorHandler

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token offset
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
	lineOffset  int  // current line offset
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable
}

const bom = 0xfeff // byte order mark, permitted as a first character only
var bomSeq = []byte{0xEF, 0xBB, 0xBF}

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
//
// Note that Init may call err one or more times if there are errors in the
// first few characters of the text.
func (l *Lexer) Init(source []byte, err ErrorHandler) {
	l.source = source
	l.err = err

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
	l.line = 1
	l.tokenLine = 1
	l.lineOffset = 0
	l.scanOffset = 0
	l.State = 0

	if bytes.HasPrefix(source, bomSeq) {
		l.scanOffset += len(bomSeq)
	}

skipChar:
	l.offset = l.scanOffset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = utf8.DecodeRune(l.source[l.offset:])
			if r == utf8.RuneError && w == 1 || r == bom {
				l.invalidRune(r, w)
				l.scanOffset += w
				goto skipChar
			}
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}

// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() Token {
restart:
	l.tokenLine = l.line
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
	hash := uint32(0)
	for state >= 0 {
		var ch int
		switch {
		case l.ch < 0:
			state = int(tmLexerAction[state*tmNumClasses])
			if state == -1 {
				l.err(l.line, l.tokenOffset, l.offset-l.tokenOffset, "Unexpected end of input reached")
			}
			continue
		case int(l.ch) < tmRuneClassLen:
			ch = int(tmRuneClass[l.ch])
		default:
			ch = 1
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state < -1 {
			break
		}
		hash = hash*uint32(31) + uint32(l.ch)

		if l.ch == '\n' {
			l.line++
			l.lineOffset = l.offset
		}
	skipChar:
		// Scan the next character.
		// Note: the following code is inlined to avoid performance implications.
		l.offset = l.scanOffset
		if l.offset < len(l.source) {
			r, w := rune(l.source[l.offset]), 1
			if r >= 0x80 {
				// not ASCII
				r, w = utf8.DecodeRune(l.source[l.offset:])
				if r == utf8.RuneError && w == 1 || r == bom {
					l.invalidRune(r, w)
					l.scanOffset += w
					goto skipChar
				}
			}
			l.scanOffset += w
			l.ch = r
		} else {
			l.ch = -1 // EOI
		}
	}
	if state >= -2 {
		if state == -1 {
			l.err(l.tokenLine, l.tokenOffset, l.offset-l.tokenOffset, "invalid token")
			goto restart
		}
		if state == -2 {
			return EOI
		}
	}

	rule := -state - 3
	switch rule {
	case 0:
		hh := hash&63
		switch hh {
		case 2:
			if hash == 0x43733a82 && bytes.Equal([]byte("lookahead"), l.source[l.tokenOffset:l.offset]) {
				rule = 61
				break
			}
			if hash == 0x6856c82 && bytes.Equal([]byte("shift"), l.source[l.tokenOffset:l.offset]) {
				rule = 49
				break
			}
		case 3:
			if hash == 0x41796943 && bytes.Equal([]byte("returns"), l.source[l.tokenOffset:l.offset]) {
				rule = 50
				break
			}
		case 6:
			if hash == 0xac107346 && bytes.Equal([]byte("assert"), l.source[l.tokenOffset:l.offset]) {
				rule = 56
				break
			}
			if hash == 0x688f106 && bytes.Equal([]byte("space"), l.source[l.tokenOffset:l.offset]) {
				rule = 69
				break
			}
		case 7:
			if hash == 0x32a007 && bytes.Equal([]byte("left"), l.source[l.tokenOffset:l.offset]) {
				rule = 52
				break
			}
		case 10:
			if hash == 0x5fb57ca && bytes.Equal([]byte("input"), l.source[l.tokenOffset:l.offset]) {
				rule = 51
				break
			}
		case 11:
			if hash == 0xfde4e8cb && bytes.Equal([]byte("brackets"), l.source[l.tokenOffset:l.offset]) {
				rule = 46
				break
			}
		case 12:
			if hash == 0x621a30c && bytes.Equal([]byte("lexer"), l.source[l.tokenOffset:l.offset]) {
				rule = 73
				break
			}
		case 13:
			if hash == 0x5c2854d && bytes.Equal([]byte("empty"), l.source[l.tokenOffset:l.offset]) {
				rule = 57
				break
			}
			if hash == 0x658188d && bytes.Equal([]byte("param"), l.source[l.tokenOffset:l.offset]) {
				rule = 62
				break
			}
		case 14:
			if hash == 0x36758e && bytes.Equal([]byte("true"), l.source[l.tokenOffset:l.offset]) {
				rule = 39
				break
			}
		case 20:
			if hash == 0x375194 && bytes.Equal([]byte("void"), l.source[l.tokenOffset:l.offset]) {
				rule = 68
				break
			}
		case 24:
			if hash == 0x9fd29358 && bytes.Equal([]byte("language"), l.source[l.tokenOffset:l.offset]) {
				rule = 71
				break
			}
		case 25:
			if hash == 0xb96da299 && bytes.Equal([]byte("inline"), l.source[l.tokenOffset:l.offset]) {
				rule = 47
				break
			}
		case 28:
			if hash == 0x677c21c && bytes.Equal([]byte("right"), l.source[l.tokenOffset:l.offset]) {
				rule = 53
				break
			}
		case 31:
			if hash == 0xc4ab3c1f && bytes.Equal([]byte("parser"), l.source[l.tokenOffset:l.offset]) {
				rule = 74
				break
			}
		case 32:
			if hash == 0x1a9a0 && bytes.Equal([]byte("new"), l.source[l.tokenOffset:l.offset]) {
				rule = 41
				break
			}
			if hash == 0x540c92a0 && bytes.Equal([]byte("nonempty"), l.source[l.tokenOffset:l.offset]) {
				rule = 58
				break
			}
			if hash == 0x34a220 && bytes.Equal([]byte("prec"), l.source[l.tokenOffset:l.offset]) {
				rule = 48
				break
			}
		case 34:
			if hash == 0x1bc62 && bytes.Equal([]byte("set"), l.source[l.tokenOffset:l.offset]) {
				rule = 45
				break
			}
		case 35:
			if hash == 0x5cb1923 && bytes.Equal([]byte("false"), l.source[l.tokenOffset:l.offset]) {
				rule = 40
				break
			}
			if hash == 0xb5e903a3 && bytes.Equal([]byte("global"), l.source[l.tokenOffset:l.offset]) {
				rule = 59
				break
			}
		case 37:
			if hash == 0xb96173a5 && bytes.Equal([]byte("import"), l.source[l.tokenOffset:l.offset]) {
				rule = 44
				break
			}
			if hash == 0x6748e2e5 && bytes.Equal([]byte("separator"), l.source[l.tokenOffset:l.offset]) {
				rule = 42
				break
			}
		case 38:
			if hash == 0xc846f566 && bytes.Equal([]byte("reduce"), l.source[l.tokenOffset:l.offset]) {
				rule = 75
				break
			}
		case 40:
			if hash == 0x53d6f968 && bytes.Equal([]byte("nonassoc"), l.source[l.tokenOffset:l.offset]) {
				rule = 54
				break
			}
		case 42:
			if hash == 0xbddafb2a && bytes.Equal([]byte("layout"), l.source[l.tokenOffset:l.offset]) {
				rule = 70
				break
			}
			if hash == 0x35f42a && bytes.Equal([]byte("soft"), l.source[l.tokenOffset:l.offset]) {
				rule = 65
				break
			}
		case 44:
			if hash == 0x2fff6c && bytes.Equal([]byte("flag"), l.source[l.tokenOffset:l.offset]) {
				rule = 63
				break
			}
		case 50:
			if hash == 0xc32 && bytes.Equal([]byte("as"), l.source[l.tokenOffset:l.offset]) {
				rule = 43
				break
			}
		case 51:
			if hash == 0xc1e742f3 && bytes.Equal([]byte("no-eoi"), l.source[l.tokenOffset:l.offset]) {
				rule = 64
				break
			}
		case 52:
			if hash == 0x8d046634 && bytes.Equal([]byte("explicit"), l.source[l.tokenOffset:l.offset]) {
				rule = 60
				break
			}
		case 53:
			if hash == 0x6be81575 && bytes.Equal([]byte("generate"), l.source[l.tokenOffset:l.offset]) {
				rule = 55
				break
			}
		case 56:
			if hash == 0x5a5a978 && bytes.Equal([]byte("class"), l.source[l.tokenOffset:l.offset]) {
				rule = 66
				break
			}
		case 57:
			if hash == 0x1df56d39 && bytes.Equal([]byte("interface"), l.source[l.tokenOffset:l.offset]) {
				rule = 67
				break
			}
		case 59:
			if hash == 0x3291bb && bytes.Equal([]byte("lalr"), l.source[l.tokenOffset:l.offset]) {
				rule = 72
				break
			}
		}
	}

	token := tmToken[rule]
	space := false
	switch rule {
	case 0: // ID: /[a-zA-Z_]([a-zA-Z_\-0-9]*[a-zA-Z_0-9])?|'([^\n\\']|\\.)*'/
		{ l.value = l.Text(); }
	case 1: // regexp: /\/{reFirst}{reChar}*\//
		{ text := l.Text(); l.value = text[1:len(text)-2] }
	case 2: // scon: /"([^\n\\"]|\\.)*"/
		{ text := l.Text(); l.value = text[1:len(text)-2] }
	case 3: // icon: /\-?[0-9]+/
		{ l.value, _ = strconv.ParseInt(l.Text(), 10, 64) }
	case 5: // _skip: /[\n\r\t ]+/
		space = true
	case 6: // _skip_comment: /#.*(\r?\n)?/
		space = true
	case 7: // _skip_multiline: /\/\*{commentChars}\*\//
		space = true
	}
	if space {
		goto restart
	}
	return token
}

func (l *Lexer) invalidRune(r rune, w int) {
	switch r {
	case utf8.RuneError:
		l.err(l.line, l.offset, w, "illegal UTF-8 encoding")
	case bom:
		l.err(l.line, l.offset, w, "illegal byte order mark")
	}
}

// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}

// Line returns the line number of the last token returned by Next().
func (l *Lexer) Line() int {
	return l.tokenLine
}

// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return string(l.source[l.tokenOffset:l.offset])
}

func (l *Lexer) Value() interface{} {
	return l.value
}
