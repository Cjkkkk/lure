package parser

import (
	"awesomeProject/memory"
	"awesomeProject/lexer"
	"fmt"
	err "awesomeProject/error"
)

/*
Interpreter
*/

type Interpreter struct {
	environment memory.Environment
}
func MakeInterpreter() Interpreter{
	return Interpreter{environment: memory.Environment{ Enclosing: nil,Values:map[string]interface{}{}}}
}

func (i *Interpreter) Interpret_ (statements []Stmt) {
	defer func() {
		if e := recover();e != nil {
			fmt.Println(e)
		}
	}()
	for _, statement :=  range statements {
		i.Execute(statement)
	}
}
//-------------------
// 执行statements
func (i *Interpreter) Execute (stmt Stmt) {
	stmt.Eval(i)
}

func (i *Interpreter) VisitPrintStmt(stmt Print) {
	value := i.Evaluate(stmt.Expression)
	fmt.Println(value)
}
func (i *Interpreter) VisitExpressionStmt(stmt Expression){
	i.Evaluate(stmt.Expression)
}

func (i *Interpreter) VisitVarStmt(stmt Var){
	var value interface{}
	if stmt.Expression != nil {
		value = i.Evaluate(stmt.Expression)
	}
	i.environment.Define(stmt.Name.Lexeme, value)
}

func (i *Interpreter) executeBlock(stmts []Stmt, e *memory.Environment){
	// todo
}
// ---------------------------------------------------
// 化简expression的值 return interface{}

/* Evaluate expression*/
func (i *Interpreter) Evaluate(expr Expr) interface{}{
	return expr.Eval(i)
}

func (i *Interpreter) VisitBlockStmt(stmt Block){
	i.executeBlock(stmt.Statements, &memory.Environment{Enclosing: &i.environment, Values: map[string]interface{}{}})
}

func (i *Interpreter) VisitAssignExpr (expr Assign) interface{}{
	i.environment.Assign(expr.Name, i.Evaluate(expr.Expr))
	return true
}
func (i *Interpreter) VisitLiteralExpr(expr Literal) interface{}{
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr Grouping) interface{}{
	return expr.Expression.Eval(i)
}

func (i *Interpreter) VisitBinaryExpr(expr Binary) interface{}{
	left := expr.Left.Eval(i)
	right := expr.Right.Eval(i)
	switch expr.Operator.Type {
	case lexer.MINUS:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d - r_d
	case lexer.PLUS:
		r_d, r_ok := right.(float64)
		l_d, l_ok := left.(float64)
		if r_ok && l_ok {
			return l_d + r_d
		}
		r_s, r_ok := right.(string)
		l_s, l_ok := left.(string)
		if r_ok && l_ok {
			return l_s + r_s
		}
		e := err.RunTimeError{ Token: expr.Operator , Msg: "Operands must be two numbers or two strings."}
		panic(e.Error())
	case lexer.SLASH:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d / r_d
	case lexer.STAR:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d * r_d
	case lexer.BANG_EQUAL: return !i.isEqual(left, right)
	case lexer.EQUAL_EQUAL: return i.isEqual(left, right)
	case lexer.GREATER:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d >= r_d
	case lexer.GREATER_EQUAL:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d >= r_d
	case lexer.LESS:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d < r_d
	case lexer.LESS_EQUAL:
		l_d , r_d := checkNumberOperands(expr.Operator, left, right)
		return l_d <= r_d
	}
	// Unreachable.
	return nil
}
func (i *Interpreter) VisitUnaryExpr (expr Unary) interface{} {
	right := expr.Right.Eval(i)
	switch expr.Operator.Type {
	case lexer.BANG:
		return !i.isTruthy(right)
	case lexer.MINUS:
		return - checkNumberOperand(expr.Operator, right)
	}
	// Unreachable.
	return nil
}

func (i *Interpreter) VisitVariableExpr (expr Variable) interface{}{
	return i.environment.Get(expr.name)
}

// ------------------------------------------------------
// 辅助函数
func checkNumberOperand(operator lexer.Token, operand interface{}) float64{
	d, ok := operand.(float64)
	if ! ok {
		e := err.RunTimeError{ Token: operator , Msg: "Operand must be a number."}
		panic(e.Error())
	}
	return d
}

func checkNumberOperands(operator lexer.Token, left interface{}, right interface{}) (float64, float64){
	ld, l_ok := left.(float64)
	rd, r_ok := right.(float64)
	if !l_ok || !r_ok {
		e := err.RunTimeError{ Token: operator , Msg: "Operands must be a number."}
		panic(e.Error())
	}
	return ld, rd
}

func (i *Interpreter) isTruthy(object interface{}) bool{
	if object == nil {
		return false
	}
	if r, ok := object.(bool); ok {
		return r
	}
	return true
}

func (i *Interpreter) isEqual(a interface{}, b interface{}) bool {
	// nil is only equal to nil.
	if a == nil && b == nil {
		return true
	}
	if a == nil{
		return false
	}
	return a == b
}
