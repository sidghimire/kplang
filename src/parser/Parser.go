package parser

import (
	"fmt"
	"kplang/src/env"
	"kplang/src/interpreter"
	"kplang/src/lexer"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

const (
	TOKEN_EQUALS = "="
	TOKEN_PLUS   = "+"
	TOKEN_MINUS  = "-"
	TOKEN_MUL    = "*"
	TOKEN_DIV    = "/"
	TOKEN_NOT    = "!"
)

func Parser(tokens []lexer.Token, env *env.Environment) {
	i := 0

	for i < len(tokens) {

		if (tokens[i].Type == lexer.TOKEN_PRINT || tokens[i].Type == lexer.TOKEN_PRINTLN) && tokens[i+1].Type == lexer.TOKEN_STRING {
			if v, ok := tokens[i+1].Value.(int); ok {
				if tokens[i].Type == lexer.TOKEN_PRINT {
					interpreter.PRINT(v)
				} else {
					interpreter.PRINTLN(v)
				}
			} else if v, ok := tokens[i+1].Value.(string); ok {
				if tokens[i].Type == lexer.TOKEN_PRINT {
					interpreter.PRINT(v)
				} else {
					interpreter.PRINTLN(v)
				}
			}
		} else if (tokens[i].Type == lexer.TOKEN_PRINT || tokens[i].Type == lexer.TOKEN_PRINTLN) && tokens[i+1].Type == lexer.TOKEN_IDENTIFIER {
			identifierValue, _ := env.Get(tokens[i+1].Value.(string))
			if _, ok := tokens[i+1].Value.(int); ok {
				if tokens[i].Type == lexer.TOKEN_PRINT {
					interpreter.PRINT(identifierValue)
				} else {
					interpreter.PRINTLN(identifierValue)
				}
			} else if _, ok := tokens[i+1].Value.(string); ok {
				if tokens[i].Type == lexer.TOKEN_PRINT {
					interpreter.PRINT(identifierValue)
				} else {
					interpreter.PRINTLN(identifierValue)
				}
			}
		} else if tokens[i].Type == lexer.TOKEN_LET {
			if tokens[i+1].Type != lexer.TOKEN_IDENTIFIER {
				panic("Expected Identifier after 'let'")
			}
			varName := tokens[i+1].Value.(string)

			if tokens[i+2].Type != lexer.TOKEN_EQUALS {
				panic("Expected '=' after identifier")
			}
			var value interface{}
			j := i + 3
			compiled := ""
			stringControl := false
			for j < len(tokens) {
				if tokens[j].Type == lexer.TOKEN_EOL || tokens[j].Type == "EOF" {
					break
				}
				if tokens[j].Type == "IDENTIFIER" {
					identifierValue, _ := env.Get(tokens[j].Value.(string))
					compiled = compiled + fmt.Sprintf("%g", identifierValue)
				} else if tokens[j].Type == lexer.TOKEN_NUMBER {
					valStr := tokens[i+3].Value.(string)
					compiled = compiled + valStr
					j++

				} else if tokens[j].Type == "LEN" {

					if tokens[j+1].Type != lexer.TOKEN_BRACKET_OPEN {
						panic("Syntax Error")
					}

					if tokens[j+2].Type != lexer.TOKEN_STRING {
						panic("balen can only take a string")
					}
					if tokens[j+3].Type != lexer.TOKEN_BRACKET_CLOSE {
						panic("Syntax Error")
					}
					length := len(tokens[i+5].Value.(string))
					compiled = compiled + strconv.Itoa(length)
					j = j + 4
				} else if tokens[j].Type == "RANDOM" {
					if tokens[j+1].Type != lexer.TOKEN_BRACKET_OPEN {
						panic("Syntax Error")
					}

					if tokens[j+2].Type != lexer.TOKEN_NUMBER {
						panic("oli can only take a number")
					}
					if tokens[j+3].Type != lexer.TOKEN_COMMA {
						panic("Comma Xutyo")
					}
					if tokens[j+4].Type != lexer.TOKEN_NUMBER {
						panic("oli can only take a number")
					}
					if tokens[j+5].Type != lexer.TOKEN_BRACKET_CLOSE {
						panic("Syntax Error")
					}
					max, _ := strconv.Atoi(tokens[j+4].Value.(string))
					min, _ := strconv.Atoi(tokens[j+2].Value.(string))

					val := rand.Intn(max-min+1) + min
					compiled = compiled + strconv.Itoa(val)
					j = j + 6
				} else if tokens[j].Type == "RANDOMWORD" {
					if tokens[j+1].Type != lexer.TOKEN_BRACKET_OPEN {
						panic("Syntax Error")
					}
					k := j + 2
					wordList := []string{}
					endFound := false

					for !endFound {
						if tokens[k].Type != lexer.TOKEN_STRING {
							panic("deuba can only hardly take words")
						} else {
							wordList = append(wordList, tokens[k].Value.(string))
						}
						if tokens[k+1].Type != lexer.TOKEN_COMMA && tokens[k+1].Type == lexer.TOKEN_BRACKET_CLOSE {
							endFound = !endFound
						} else if tokens[k+1].Type != lexer.TOKEN_COMMA {
							panic("Missing Comma")
						}
						k = k + 2
					}
					randomIndex := rand.Intn(len(wordList))
					word := wordList[randomIndex]
					stringControl = true

					j = j + k - 2
					env.Set(varName, word)

				} else if tokens[j].Type == "SPLIT" {
					if tokens[j+1].Type != lexer.TOKEN_BRACKET_OPEN {
						panic("Syntax Error")
					}
					if tokens[j+2].Type != lexer.TOKEN_STRING {
						panic("prachanda can only take a string")
					}
					if tokens[j+3].Type != lexer.TOKEN_COMMA {
						panic("Comma Xutyo")
					}
					if tokens[j+4].Type != lexer.TOKEN_STRING {
						panic("prachanda can only take a string")
					}
					if tokens[j+5].Type != lexer.TOKEN_COMMA {
						panic("Comma Xutyo")
					}
					if tokens[j+6].Type != lexer.TOKEN_NUMBER {
						panic("prachanda can only take a number")
					}
					if tokens[j+7].Type != lexer.TOKEN_BRACKET_CLOSE {
						panic("Syntax Error")
					}

					text := tokens[j+2].Value.(string)

					delimiter := tokens[j+4].Value.(string)

					index, _ := strconv.Atoi(tokens[j+6].Value.(string))

					parts := strings.Split(text, delimiter)

					env.Set(varName, parts[index])
					j = j + 8

					stringControl = true

				} else if tokens[j].Type == "STRING" {
					value = tokens[i].Value.(string)
					env.Set(varName, value)
					return
				}
			}
			if !stringControl {
				expression, err := govaluate.NewEvaluableExpression(compiled)
				if err != nil {
					fmt.Println("Error in expression:", err)
					return
				}
				value, err = expression.Evaluate(nil)
				if err != nil {
					fmt.Println("Error evaluating expression:", err)
					return
				}

				env.Set(varName, value)
			}
			stringControl = !stringControl
		} else if tokens[i].Type == lexer.TOKEN_INPUT {
			if tokens[i+1].Type != lexer.TOKEN_IDENTIFIER {
				panic("Expected Identifier after kre")
			}
			varName := tokens[i+1].Value.(string)
			var value string
			fmt.Scan(&value)
			env.Set(varName, value)
		} else if tokens[i].Type == lexer.TOKEN_IF {
			foundOpening, foundClosing := false, false
			open, close := -1, -1
			nestedLevel := 0
			j := i

			for j < len(tokens) {
				if tokens[j].Type == "TOKEN_CURLY_OPEN" {
					if !foundOpening {
						open = j
						foundOpening = true
					}
					nestedLevel++
				}
				if tokens[j].Type == "TOKEN_CURLY_CLOSE" {
					nestedLevel--
					if nestedLevel == 0 {
						close = j
						foundClosing = true
						break
					}
				}
				j++
			}

			if !foundOpening || !foundClosing {
				panic("Incorrect Syntax: missing opening or closing braces")
			} else {
				compiled := ""
				for k := i + 1; k < open; k++ {
					if tokens[k].Type == lexer.TOKEN_IDENTIFIER {
						identifierValue, _ := env.Get(tokens[k].Value.(string))
						compiled += fmt.Sprintf("%v", identifierValue)
					} else {
						compiled += tokens[k].Value.(string)
					}
				}

				expression, err := govaluate.NewEvaluableExpression(compiled)
				if err != nil {
					fmt.Println("Error in expression:", err)
					return
				}
				value, err := expression.Evaluate(nil)
				if err != nil {
					fmt.Println("Error evaluating expression:", err)
					return
				}

				// Execute if condition is true
				if value == true {
					tempArray := tokens[open+1 : close]
					Parser(tempArray, env) // Recursive call for nested statements
				}
			}
			i = close + 1
		} else if tokens[i].Type == lexer.TOKEN_FUNCTION {

			foundClosing := false
			foundOpening := false
			open := -1
			close := -1
			j := i
			for j < len(tokens) {
				if tokens[j].Type == "TOKEN_CURLY_OPEN" {
					foundOpening = true
					open = j
				}
				if tokens[j].Type == "TOKEN_CURLY_CLOSE" {
					foundClosing = true
					close = j
				}
				if foundClosing == true && foundOpening == true {
					break
				}
				j++
			}
			if foundOpening == false || foundClosing == false {
				panic("Incorrect Syntax")
			} else {
				if tokens[i+1].Type == lexer.TOKEN_IDENTIFIER {
					tempArray := tokens[open+1 : close]
					varName := tokens[i+1].Value.(string)
					env.Set(varName, tempArray)
				}
			}
			i = close + 1
		} else if tokens[i].Type == lexer.TOKEN_WHILE {
		here:
			foundClosing := false
			foundOpening := false
			open := -1
			close := -1
			j := i
			for j < len(tokens) {
				if tokens[j].Type == "TOKEN_CURLY_OPEN" {
					foundOpening = true
					open = j
				}
				if tokens[j].Type == "TOKEN_CURLY_CLOSE" {
					foundClosing = true
					close = j
				}
				if foundClosing == true && foundOpening == true {
					break
				}
				j++
			}
			if foundOpening == false || foundClosing == false {
				panic("Incorrect Syntax")
			} else {
				j := i
				j++
				compiled := ""

				for j < len(tokens) {

					if tokens[j].Type == lexer.TOKEN_EOL {
						j++
						continue
					}
					if tokens[j].Type == lexer.TOKEN_IDENTIFIER {
						identifierValue, _ := env.Get(tokens[j].Value.(string))
						compiled = compiled + fmt.Sprintf("%g", identifierValue)
					} else if tokens[j].Type == lexer.TOKEN_NUMBER {
						compiled = compiled + tokens[j].Value.(string)
					} else if tokens[j].Type == lexer.TOKEN_CURLY_OPEN || tokens[j].Type == lexer.TOKEN_CURLY_CLOSE {
						break
					} else {
						compiled = compiled + tokens[j].Value.(string)
					}

					j++
				}
				expression, err := govaluate.NewEvaluableExpression(compiled)
				if err != nil {
					fmt.Println("Error in expression:", err)
					return
				}
				value, err := expression.Evaluate(nil)
				if err != nil {
					fmt.Println("Error evaluating expression:", err)
					return
				}
				if value == true {
					tempArray := tokens[open+1 : close]
					Parser(tempArray, env)
					goto here
				}
			}
			i = close + 1

		} else if tokens[i].Type == lexer.TOKEN_IDENTIFIER && tokens[i+1].Type == lexer.TOKEN_BRACKET_OPEN && tokens[i+2].Type == lexer.TOKEN_BRACKET_CLOSE {
			identifierValue, _ := env.Get(tokens[i].Value.(string))
			tokenArray, _ := identifierValue.([]lexer.Token)
			Parser(tokenArray, env)
		}
		i++
	}
}
