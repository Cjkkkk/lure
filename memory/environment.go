package memory

import (
	"awesomeProject/lexer"
	Err "awesomeProject/error"
)

type Environment struct{
	Enclosing *Environment
	Values map[string]interface{}
}

// 赋值语句
func (e *Environment) Assign(name lexer.Token, value interface{}) {
	if _, OK := e.Values[name.Lexeme]; OK {
		e.Values[name.Lexeme] = value
		return
	}
	if e.Enclosing != nil {
		e.Enclosing.Assign(name, value)
		return
	}
	err := Err.RunTimeError{Msg:"Undefined variable '" + name.Lexeme + "'.", Token: name}
	panic(err.Error())
}

// 定义变量
func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

// 获取变量的值
func  (e *Environment) Get(name lexer.Token) interface{} {
	if data, OK := e.Values[name.Lexeme]; OK {
		return data
	}
	if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	}
	err := Err.RunTimeError{ Token:name, Msg:"Undefined variable '" + name.Lexeme + "'."}
	panic(err.Error())
}