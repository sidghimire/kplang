package main

import (
	errorhandler "kplang/src/ErrorHandler"
	"kplang/src/env"
	"kplang/src/lexer"
	"kplang/src/parser"
	"os"
)

func main() {

	code, err := os.ReadFile("./hello.sid")
	errorhandler.CheckError(err, "Failed to open file")
	code_string := string(code)
	tokens := lexer.Lexer(code_string)
	//fmt.Println(tokens)
	env := env.NewEnvironment()
	parser.Parser(tokens, env)
	//fmt.Println(env.Get("HelloWorld"))
}
