package parser

import "awesomeProject/lexer"

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
type Unary struct{
	Operator lexer.Token
	Right Expr
}

type Variable struct {
	name lexer.Token
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

type Expression struct {
	Expression Expr
}
type Print struct {
	Expression Expr
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
