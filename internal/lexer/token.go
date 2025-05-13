package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT // =
	EQUALS     // ==
	NOT        // !
	NOT_EQUALS // !=

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND

	DOT
	DOT_DOT
	SEMICOLON
	COLON
	QUESTION
	COMMA

	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS

	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	//RESERVED KEYWORD
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPE
	IN
)

var tokenKindStringMap = map[TokenKind]string{
    EOF:           "EOF",
    NUMBER:        "NUMBER",
    STRING:        "STRING",
    IDENTIFIER:    "IDENTIFIER",
    OPEN_BRACKET:  "OPEN_BRACKET",
    CLOSE_BRACKET: "CLOSE_BRACKET",
    OPEN_CURLY:    "OPEN_CURLY",
    CLOSE_CURLY:   "CLOSE_CURLY",
    OPEN_PAREN:    "OPEN_PAREN",
    CLOSE_PAREN:   "CLOSE_PAREN",
    ASSIGNMENT:    "ASSIGNMENT",
    EQUALS:        "EQUALS",
    NOT:           "NOT",
    NOT_EQUALS:    "NOT_EQUALS",
    LESS:          "LESS",
    LESS_EQUALS:   "LESS_EQUALS",
    GREATER:       "GREATER",
    GREATER_EQUALS:"GREATER_EQUALS",
    OR:            "OR",
    AND:           "AND",
    DOT:           "DOT",
    DOT_DOT:       "DOT_DOT",
    SEMICOLON:     "SEMICOLON",
    COLON:         "COLON",
    QUESTION:      "QUESTION",
    COMMA:         "COMMA",
    PLUS_PLUS:     "PLUS_PLUS",
    MINUS_MINUS:   "MINUS_MINUS",
    PLUS_EQUALS:   "PLUS_EQUALS",
    MINUS_EQUALS:  "MINUS_EQUALS",
    PLUS:          "PLUS",
    DASH:          "DASH",
    SLASH:         "SLASH",
    STAR:          "STAR",
    PERCENT:       "PERCENT",
    LET:           "LET",
    CONST:         "CONST",
    CLASS:         "CLASS",
    NEW:           "NEW",
    IMPORT:        "IMPORT",
    FROM:          "FROM",
    FN:            "FN",
    IF:            "IF",
    ELSE:          "ELSE",
    FOREACH:       "FOREACH",
    WHILE:         "WHILE",
    FOR:           "FOR",
    EXPORT:        "EXPORT",
    TYPE:          "TYPE",
    IN:            "IN",
}


type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind,value string) Token {
	return Token{kind,value}
}

func (tk *Token) isOneOrMany(expectedTokens ...TokenKind)bool {
	return slices.Contains(expectedTokens, tk.Kind)
}

func (tk *Token) Debug() {
	if tk.isOneOrMany(IDENTIFIER,NUMBER,STRING){
		fmt.Printf("%s (%s)\n",tk.tokenKindString(),tk.Value)
	}else {
		fmt.Printf("%s ()\n",tk.tokenKindString())
	}
}

func(tk *Token) tokenKindString() string {
	return tokenKindStringMap[tk.Kind]
}