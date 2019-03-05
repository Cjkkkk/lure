package error

import (
	"awesomeProject/lexer"
	"fmt"
	"strconv"
)

var HadError = false

type ParseError struct {
	
}

func (p ParseError) Error() string{
	return "2333"
}
func Error(token lexer.Token, message string) {
	if token.Type == lexer.EOF {
		Report(token.Line, " at end", message)
	} else {
		Report(token.Line, " at '" + token.Lexeme + "'", message)
	}
}

func Report(line int, where string, message string) {
	fmt.Println("[line " + strconv.Itoa(line) + "] Error" + where + ": " + message)
	HadError = true
}