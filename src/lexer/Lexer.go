package lexer

import (
	"strings"
	"unicode"
)

const (
	TOKEN_LET            = "LET"
	TOKEN_IDENTIFIER     = "IDENTIFIER"
	TOKEN_STRING_VALUE   = "STRING_VALUE"
	TOKEN_NUMBER         = "NUMBER"
	TOKEN_PRINT          = "PRINT"
	TOKEN_PRINTLN        = "PRINTLN"
	TOKEN_STRING         = "STRING"
	TOKEN_EOF            = "EOF"
	TOKEN_EOL            = "EOL"
	TOKEN_EQUALS         = "EQUALS"
	TOKEN_EQUALS_TO      = "EQUALS_TO"
	TOKEN_GREATER        = "GREATER"
	TOKEN_LESSER         = "LESSER"
	TOKEN_GREATER_EQUALS = "GREATER_EQUALS"
	TOKEN_LESSER_EQUALS  = "LESSER_EQUALS"
	TOKEN_NOT_EQUALS     = "NOT_EQUALS"
	TOKEN_PLUS           = "PLUS"
	TOKEN_MINUS          = "MINUS"
	TOKEN_MUL            = "MULTIPLY"
	TOKEN_DIV            = "DIVIDE"
	TOKEN_NOT            = "NOT"
	TOKEN_BRACKET_OPEN   = "TOKEN_BRACKET_OPEN"
	TOKEN_BRACKET_CLOSE  = "TOKEN_BRACKET_CLOSE"
	TOKEN_CURLY_OPEN     = "TOKEN_CURLY_OPEN"
	TOKEN_CURLY_CLOSE    = "TOKEN_CURLY_CLOSE"

	TOKEN_FUNCTION   = "FUNCTION"
	TOKEN_IF         = "IF"
	TOKEN_ELSE       = "ELSE"
	TOKEN_WHILE      = "WHILE"
	TOKEN_INPUT      = "INPUT"
	TOKEN_TRY        = "TRY"
	TOKEN_CATCH      = "CATCH"
	TOKEN_LEN        = "LEN"
	TOKEN_RANDOM     = "RANDOM"
	TOKEN_COMMA      = "COMMA"
	TOKEN_SPLIT      = "SPLIT"
	TOKEN_RANDOMWORD = "RANDOMWORD"
)

type Token struct {
	Type  string
	Value interface{}
}

func Lexer(code string) []Token {
	tokens := []Token{}
	i := 0
	for i < len(code) {
		if unicode.IsSpace(rune(code[i])) && code[i] != '\n' {
			i++
			continue
		}
		if code[i] == '\n' {
			tokens = append(tokens, Token{Type: TOKEN_EOL, Value: string(code[i])})
			i++
		} else if strings.HasPrefix(code[i:], "lekh_arko") {
			tokens = append(tokens, Token{Type: TOKEN_PRINTLN, Value: "println"})
			i = i + 9
		} else if strings.HasPrefix(code[i:], "kre?") {
			tokens = append(tokens, Token{Type: TOKEN_INPUT, Value: "scanln"})
			i = i + 4
		} else if strings.HasPrefix(code[i:], "balen") {
			tokens = append(tokens, Token{Type: TOKEN_LEN, Value: "len"})
			i = i + 5
		} else if strings.HasPrefix(code[i:], "oli") {
			tokens = append(tokens, Token{Type: TOKEN_RANDOM, Value: "random"})
			i = i + 3
		} else if strings.HasPrefix(code[i:], ",") {
			tokens = append(tokens, Token{Type: TOKEN_COMMA, Value: "comma"})
			i = i + 1
		} else if strings.HasPrefix(code[i:], "prachanda") {
			tokens = append(tokens, Token{Type: TOKEN_SPLIT, Value: "split"})
			i = i + 9
		} else if strings.HasPrefix(code[i:], "deuba") {
			tokens = append(tokens, Token{Type: TOKEN_RANDOMWORD, Value: "randomword"})
			i = i + 5
		} else if strings.HasPrefix(code[i:], "herta") {
			tokens = append(tokens, Token{Type: TOKEN_TRY, Value: "try"})
			i = i + 5
		} else if strings.HasPrefix(code[i:], "pakdyo") {
			tokens = append(tokens, Token{Type: TOKEN_CATCH, Value: "catch"})
			i = i + 6
		} else if strings.HasPrefix(code[i:], "kaam") {
			tokens = append(tokens, Token{Type: TOKEN_FUNCTION, Value: "function"})
			i = i + 4
		} else if strings.HasPrefix(code[i:], "yedi") {
			tokens = append(tokens, Token{Type: TOKEN_IF, Value: "if"})
			i = i + 4
		} else if strings.HasPrefix(code[i:], "jabasamma") {
			tokens = append(tokens, Token{Type: TOKEN_WHILE, Value: "while"})
			i = i + 9
		} else if strings.HasPrefix(code[i:], "lekh") {
			tokens = append(tokens, Token{Type: TOKEN_PRINT, Value: "print"})
			i += 4
		} else if strings.HasPrefix(code[i:], "la") {
			tokens = append(tokens, Token{Type: TOKEN_LET, Value: "let"})
			i += 2
		} else if code[i] == '"' {
			i++
			start := i
			for i < len(code) && code[i] != '"' {
				i++
			}
			word := code[start:i]
			tokens = append(tokens, Token{Type: TOKEN_STRING, Value: word})
			i++
		} else if i+2 < len(code) && strings.HasPrefix(code[i:i+2], "==") {
			tokens = append(tokens, Token{Type: TOKEN_EQUALS_TO, Value: "=="})
			i = i + 2
		} else if strings.HasPrefix(code[i:i+1], ">") {
			tokens = append(tokens, Token{Type: TOKEN_GREATER, Value: ">"})
			i++
		} else if strings.HasPrefix(code[i:i+1], "<") {
			tokens = append(tokens, Token{Type: TOKEN_LESSER, Value: "<"})
			i++
		} else if i+2 < len(code) && strings.HasPrefix(code[i:i+2], ">=") {
			tokens = append(tokens, Token{Type: TOKEN_GREATER_EQUALS, Value: ">="})
			i = i + 2
		} else if i+2 < len(code) && strings.HasPrefix(code[i:i+2], "<=") {
			tokens = append(tokens, Token{Type: TOKEN_LESSER_EQUALS, Value: "<="})
			i = i + 2
		} else if i+2 < len(code) && strings.HasPrefix(code[i:i+2], "!=") {
			tokens = append(tokens, Token{Type: TOKEN_NOT_EQUALS, Value: "!="})
			i = i + 2
		} else if code[i] == '=' {
			tokens = append(tokens, Token{Type: TOKEN_EQUALS, Value: "="})
			i++
		} else if code[i] == '{' {
			tokens = append(tokens, Token{Type: TOKEN_CURLY_OPEN, Value: "{"})
			i++
		} else if code[i] == '}' {
			tokens = append(tokens, Token{Type: TOKEN_CURLY_CLOSE, Value: "}"})
			i++
		} else if code[i] == '(' {
			tokens = append(tokens, Token{Type: TOKEN_BRACKET_OPEN, Value: "("})
			i++
		} else if code[i] == ')' {
			tokens = append(tokens, Token{Type: TOKEN_BRACKET_CLOSE, Value: ")"})
			i++
		} else if code[i] == '+' {
			tokens = append(tokens, Token{Type: TOKEN_PLUS, Value: "+"})
			i++
		} else if code[i] == '-' {
			tokens = append(tokens, Token{Type: TOKEN_MINUS, Value: "-"})
			i++
		} else if code[i] == '*' {
			tokens = append(tokens, Token{Type: TOKEN_MUL, Value: "*"})
			i++
		} else if code[i] == '/' {
			tokens = append(tokens, Token{Type: TOKEN_DIV, Value: "/"})
			i++
		} else if code[i] == '!' {
			tokens = append(tokens, Token{Type: TOKEN_NOT, Value: "!"})
			i++
		} else if unicode.IsLetter(rune(code[i])) {
			start := i
			for i < len(code) && (unicode.IsLetter(rune(code[i])) || unicode.IsDigit(rune(code[i]))) {
				i++
			}
			tokens = append(tokens, Token{Type: TOKEN_IDENTIFIER, Value: code[start:i]})
		} else if unicode.IsDigit(rune(code[i])) {
			start := i
			for i < len(code) && unicode.IsDigit(rune(code[i])) {
				i++
			}
			tokens = append(tokens, Token{Type: TOKEN_NUMBER, Value: code[start:i]})
		} else {
			i++
		}
	}
	tokens = append(tokens, Token{Type: TOKEN_EOF, Value: ""})
	return tokens
}
