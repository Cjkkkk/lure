package parser
//
//import "fmt"
//
///*
//	AstPrinter struct 实现 Visitor接口
//*/
//type AstPrinter struct{}
//
//func (a *AstPrinter) Print(expr Expr) string{
//	return expr.Accept(a)
//}
////---------------------
//func (a *AstPrinter) BinaryAccept(expr Binary) string{
//	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
//}
//func (a *AstPrinter)LiteralAccept(expr Literal) string{
//	if expr.Value == nil {
//		return "nil"
//	}
//	if str, ok := expr.Value.(string); ok {
//		return str
//	} else if number, ok := expr.Value.(float64); ok{
//		return fmt.Sprintf("%f", number)
//	} else if number, ok := expr.Value.(int); ok {
//		return fmt.Sprintf("%d", number)
//	} else {
//		return "???"
//	}
//}
//func (a *AstPrinter) GroupingAccept(expr Grouping) string{
//	return a.parenthesize("group", expr.Expression)
//}
//func (a *AstPrinter) UnaryAccept(expr Unary) string{
//	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
//}
//
//func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
//	var builder = ""
//	builder = builder + "(" + name
//	for _, expr:= range exprs {
//		builder += " "
//		builder += expr.Accept(a)
//	}
//	builder += ")"
//	return builder
//}