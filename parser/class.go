package parser

import (
	"awesomeProject/lexer"
)

type Program struct {

}

type Declaration interface {

}


/*
	Expr
*/
type Expr interface {
	//Accept(a *AstPrinter) string
	Eval(a *Interpreter) interface{}
}

type Binary struct {
	Left Expr
	Operator lexer.Token
	Right Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value interface{}
}

type Logical struct{
	Left Expr
	Operator lexer.Token
	Right Expr
}
type Unary struct{
	Operator lexer.Token
	Right Expr
}

type Variable struct {
	name lexer.Token
}

type Assign struct {
	Name lexer.Token
	Expr Expr
}
type Call struct {
	Paren lexer.Token
	Callee Expr
	Args []Expr
}
func (b Binary) Eval(a *Interpreter) interface{}{
	return a.VisitBinaryExpr(b)
}
func (g Grouping) Eval(a *Interpreter) interface{}{
	return a.VisitGroupingExpr(g)
}
func (l Literal) Eval(a *Interpreter) interface{}{
	return a.VisitLiteralExpr(l)
}
func (u Unary) Eval(a *Interpreter) interface{}{
	return a.VisitUnaryExpr(u)
}
func (v Variable) Eval(a *Interpreter) interface{}{
	return a.VisitVariableExpr(v)
}

func (as Assign) Eval(a *Interpreter) interface{}{
	return a.VisitAssignExpr(as)
}

func (l Logical) Eval(a *Interpreter) interface{}{
	return a.VisitLogicExpr(l)
}

func (c Call) Eval(a *Interpreter) interface{}{
	return a.VisitCallExpr(c)
}
//func (b Binary) Accept(a *AstPrinter) string{
//	return a.BinaryAccept(b)
//}
//func (g Grouping) Accept(a *AstPrinter) string{
//	return a.GroupingAccept(g)
//}
//func (l Literal) Accept(a *AstPrinter) string{
//	return a.LiteralAccept(l)
//}
//func (u Unary) Accept(a *AstPrinter) string{
//	return a.UnaryAccept(u)
//}

/*
	Stmt
*/
type Stmt interface {
	Eval(a *Interpreter)
}

type WhileStatement struct {
	Condition Expr
	Branch Stmt
}

type IfStatement struct {
	Condition Expr
	ThenBranch Stmt
	ElseBranch Stmt
}
type Expression struct {
	Expression Expr
}
type Print struct {
	Expression Expr
}

type Block struct {
	Statements []Stmt
}
type Var struct {
	Name lexer.Token
	Expression Expr
}
func (e Expression) Eval(a *Interpreter){
	a.VisitExpressionStmt(e)
}
func (p Print) Eval(a *Interpreter){
	a.VisitPrintStmt(p)
}

func (v Var) Eval(a *Interpreter){
	a.VisitVarStmt(v)
}

func (b Block) Eval(a *Interpreter){
	a.VisitBlockStmt(b)
}

func (i IfStatement) Eval(a *Interpreter){
	a.VisitIfStmt(i)
}

func (w WhileStatement) Eval(a *Interpreter){
	a.VisitWhileStmt(w)
}
