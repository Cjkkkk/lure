package memory

import (
	"awesomeProject/lexer"
)

type Environment struct{
	Enclosing *Environment
	Values map[string]interface{}
}

// todo throw runtime error
func (e *Environment) Assign(name lexer.Token, value interface{}) {
	if _, OK := e.Values[name.Lexeme]; OK {
		e.Values[name.Lexeme] = value
		return
	}
	// err.RunTimeError{name, "Undefined variable '" + name.Lexeme + "'."};
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

// todo throw runtimeError
func  (e *Environment) Get(name lexer.Token) interface{} {
	if data, OK := e.Values[name.Lexeme]; OK {
		return data
	}
	if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	}
	return nil
	// err.RunTimeError{ Token:name, Msg:"Undefined variable '" + name.Lexeme + "'."}
}