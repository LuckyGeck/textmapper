package js

import (
	"fmt"
)

type Token int

const (
	UNAVAILABLE Token = iota - 1

	// An end-of-input marker token.
	EOI

	WHITESPACE
	LINETERMINATORSEQUENCE
	MULTILINECOMMENT
	SINGLELINECOMMENT
	IDENTIFIER
	BREAK // break
	CASE // case
	CATCH // catch
	CLASS // class
	CONST // const
	CONTINUE // continue
	DEBUGGER // debugger
	DEFAULT // default
	DELETE // delete
	DO // do
	ELSE // else
	EXPORT // export
	EXTENDS // extends
	FINALLY // finally
	FOR // for
	FUNCTION // function
	IF // if
	IMPORT // import
	IN // in
	INSTANCEOF // instanceof
	NEW // new
	RETURN // return
	SUPER // super
	SWITCH // switch
	THIS // this
	THROW // throw
	TRY // try
	TYPEOF // typeof
	VAR // var
	VOID // void
	WHILE // while
	WITH // with
	YIELD // yield
	AWAIT // await
	ENUM // enum
	NULL // null
	TRUE // true
	FALSE // false
	AS // as
	FROM // from
	GET // get
	LET // let
	OF // of
	SET // set
	STATIC // static
	TARGET // target
	LBRACE // {
	RBRACE // }
	LPAREN // (
	RPAREN // )
	LBRACK // [
	RBRACK // ]
	DOT // .
	SEMICOLON // ;
	COMMA // ,
	LT // <
	GT // >
	LTASSIGN // <=
	GTASSIGN // >=
	ASSIGNASSIGN // ==
	EXCLASSIGN // !=
	ASSIGNASSIGNASSIGN // ===
	EXCLASSIGNASSIGN // !==
	PLUS // +
	MINUS // -
	MULT // *
	DIV // /
	REM // %
	PLUSPLUS // ++
	MINUSMINUS // --
	LTLT // <<
	GTGT // >>
	GTGTGT // >>>
	AND // &
	OR // |
	XOR // ^
	EXCL // !
	TILDE // ~
	ANDAND // &&
	OROR // ||
	QUEST // ?
	COLON // :
	ASSIGN // =
	PLUSASSIGN // +=
	MINUSASSIGN // -=
	MULTASSIGN // *=
	DIVASSIGN // /=
	REMASSIGN // %=
	LTLTASSIGN // <<=
	GTGTASSIGN // >>=
	GTGTGTASSIGN // >>>=
	ANDASSIGN // &=
	ORASSIGN // |=
	XORASSIGN // ^=
	ASSIGNGT // =>
	NUMERICLITERAL
	STRINGLITERAL
	NOSUBSTITUTIONTEMPLATE
	TEMPLATEHEAD
	TEMPLATEMIDDLE
	TEMPLATETAIL
	REGULAREXPRESSIONLITERAL
	JSXSTRINGLITERAL
	JSXIDENTIFIER
	JSXTEXT

	terminalEnd
)

var tokenStr = [...]string{
	"EOF",

	"WHITESPACE",
	"LINETERMINATORSEQUENCE",
	"MULTILINECOMMENT",
	"SINGLELINECOMMENT",
	"IDENTIFIER",
	"break",
	"case",
	"catch",
	"class",
	"const",
	"continue",
	"debugger",
	"default",
	"delete",
	"do",
	"else",
	"export",
	"extends",
	"finally",
	"for",
	"function",
	"if",
	"import",
	"in",
	"instanceof",
	"new",
	"return",
	"super",
	"switch",
	"this",
	"throw",
	"try",
	"typeof",
	"var",
	"void",
	"while",
	"with",
	"yield",
	"await",
	"enum",
	"null",
	"true",
	"false",
	"as",
	"from",
	"get",
	"let",
	"of",
	"set",
	"static",
	"target",
	"{",
	"}",
	"(",
	")",
	"[",
	"]",
	".",
	";",
	",",
	"<",
	">",
	"<=",
	">=",
	"==",
	"!=",
	"===",
	"!==",
	"+",
	"-",
	"*",
	"/",
	"%",
	"++",
	"--",
	"<<",
	">>",
	">>>",
	"&",
	"|",
	"^",
	"!",
	"~",
	"&&",
	"||",
	"?",
	":",
	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"<<=",
	">>=",
	">>>=",
	"&=",
	"|=",
	"^=",
	"=>",
	"NUMERICLITERAL",
	"STRINGLITERAL",
	"NOSUBSTITUTIONTEMPLATE",
	"TEMPLATEHEAD",
	"TEMPLATEMIDDLE",
	"TEMPLATETAIL",
	"REGULAREXPRESSIONLITERAL",
	"JSXSTRINGLITERAL",
	"JSXIDENTIFIER",
	"JSXTEXT",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
