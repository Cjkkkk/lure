package Memory

import (
	"awesomeProject/lexer"
)

type Environment struct{
	Values map[string]interface{}
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}
// todo throw runtimeError
func  (e *Environment) Get(name lexer.Token) interface{} {
	if data, OK := e.Values[name.Lexeme]; OK {
		return data
	}
	return nil
	// error.RunTimeError{ Token:name, Msg:"Undefined variable '" + name.Lexeme + "'."}
}