package lexer

import (
	"fmt"
	"strconv"
)

type TokenType int

const(
	// Single-character Tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character Tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

var KeyWords = map[string]TokenType {
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Token struct {
	Type TokenType
	Lexeme string
	Literal interface {}
	Line int
}

func (t *Token) toString() string{
	return string(t.Type) + " " + t.Lexeme
}

type Scanner struct {
	Source string
	Tokens []Token
	Current int //0
	Start int //0
	Line int //1
}

func (s *Scanner) isAtEnd() bool{
	return s.Current >= len(s.Source)
}

func (s *Scanner) advance() byte{
	s.Current += 1
	return s.Source[s.Current-1]
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		//error(s.Line, "Unterminated string.")
	}
	s.advance()
	Value := s.Source[s.Start + 1:s.Current - 1]
	s.addToken(STRING, Value)
}
func (s *Scanner) peek() byte{
	if s.isAtEnd() {
		return 0
	}
	return s.Source[s.Current]
}

func (s *Scanner) peekNext() byte{
	if s.Current + 1 >= len(s.Source) {
		return 0
	}
	return s.Source[s.Current + 1]
}
func (s *Scanner) match(expected byte) bool{
	if s.isAtEnd() {return false}
	if s.Source[s.Current]!= expected {return false}
	s.Current++
	return true
}

func (s *Scanner) addToken(t TokenType, i interface{}) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, Token{t, text,i, s.Line})
}
func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
		case '(': s.addToken(LEFT_PAREN, nil)
		case ')': s.addToken(RIGHT_PAREN, nil)
		case '{': s.addToken(LEFT_BRACE, nil)
		case '}': s.addToken(RIGHT_BRACE, nil)
		case ',': s.addToken(COMMA, nil)
		case '.': s.addToken(DOT, nil)
		case '-': s.addToken(MINUS, nil)
		case '+': s.addToken(PLUS, nil)
		case ';': s.addToken(SEMICOLON, nil)
		case '*': s.addToken(STAR, nil)
		case '!': if s.match('=') {s.addToken(BANG_EQUAL, nil)} else {s.addToken(BANG, nil)}
		case '=': if s.match('=') {s.addToken(EQUAL_EQUAL, nil)}else {s.addToken(EQUAL, nil)}
		case '<': if s.match('=') {s.addToken(LESS_EQUAL, nil)} else{s.addToken(LESS, nil)}
		case '>': if s.match('=') {s.addToken(GREATER_EQUAL, nil)}else{ s.addToken(GREATER, nil)}
		case '/':
		if s.match('/') {
			// A comment goes until the end of the Line.
			for s.peek() != '\n' && !s.isAtEnd(){
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
		case ' ':
		case '\r':
		case '\t':
		case '\n': s.Line++
		case '"': s.string()
		default:
			if isDigit(c) {
				s.number()
			} else if isAlpha(c) {
				s.identifier()
			} else{
				//error(s.Line, "Unexpected character [" + string(c) + "].")
			}
	}
}

func isDigit(c byte ) bool{
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool{
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner)identifier(){
	for isAlphaNumeric(s.peek()){
		s.advance()
	}
	var text = s.Source[s.Start: s.Current]
	var t, OK = KeyWords[text]
	if !OK { t = IDENTIFIER }
	s.addToken(t, nil)
}
func (s *Scanner)number() {
	for isDigit(s.peek()){
		s.advance()
	}
	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}
	f, _ := strconv.ParseFloat(s.Source[s.Start: s.Current], 64)
	s.addToken(NUMBER, f)
}

func (s *Scanner)ScanTokens() []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next Lexeme.
		s.Start = s.Current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, Token{EOF,"","", s.Line})
	return s.Tokens
}

/*
	visitor 接口
*/
type Visitor interface {
	BinaryAccept(expr Binary) string
	UnaryAccept(expr Unary) string
	GroupingAccept(expr Grouping) string
	LiteralAccept(expr Literal) string
}

/*
	Expr
*/
type Expr interface {
	accept(a Visitor) string
}

func (b Binary) accept(a Visitor) string{
	return a.BinaryAccept(b)
}
func (g Grouping) accept(a Visitor) string{
	return a.GroupingAccept(g)
}
func (l Literal) accept(a Visitor) string{
	return a.LiteralAccept(l)
}
func (u Unary) accept(a Visitor) string{
	return a.UnaryAccept(u)
}
type Binary struct {
	Left Expr
	Operator Token
	Right Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value interface{}
}
type Unary struct{
	Operator Token
	Right Expr
}

/*
	AstPrinter struct 实现 Visitor接口
*/
type AstPrinter struct{}

func (a *AstPrinter) Print(expr Expr) string{
	return expr.accept(a)
}

func (a *AstPrinter) BinaryAccept(expr Binary) string{
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}
func (a *AstPrinter)LiteralAccept(expr Literal) string{
	if expr.Value == nil {
		return "nil"
	}
	if str, ok := expr.Value.(string); ok {
		return str
	} else if number, ok := expr.Value.(float64); ok{
		return fmt.Sprintf("%f", number)
	} else if number, ok := expr.Value.(int); ok {
		return fmt.Sprintf("%d", number)
	} else {
		return "???"
	}
}
func (a *AstPrinter) GroupingAccept(expr Grouping) string{
	return a.parenthesize("group", expr.Expression)
}
func (a *AstPrinter) UnaryAccept(expr Unary) string{
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder = ""
	builder = builder + "(" + name
	for _, expr:= range exprs {
		builder += " "
		builder += expr.accept(a)
	}
	builder += ")"
	return builder
}

/*
Interpreter
*/

type Interpreter struct {
	AstPrinter
}

func (i *Interpreter) visitLiteralExpr(expr Literal) interface{}{
	return expr.Value
}

func (i *Interpreter) visitGroupingExpr(expr Grouping) interface{}{
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) evaluate(expr Expr) interface{}{
	return expr.accept(i)
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
// todo
func (i *Interpreter) visitBinaryExpr(expr Binary) interface{}{
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)
	r_d, r_ok := right.(string)
	l_d, l_ok := left.(string)
	if !l_ok || !r_ok{
		return nil
	}
	r, _ := strconv.ParseFloat(r_d, 64)
	l, _ := strconv.ParseFloat(l_d, 64)
	switch expr.Operator.Type {
		case MINUS:
			return l - r
	//case PLUS:
	//	if (left instanceof Double && right instanceof Double) {
	//	return (double)left + (double)right;
	//}
	//
	//	if (left instanceof String && right instanceof String) {
	//	return (String)left + (String)right;
	//}
	case SLASH:
			return l / r
		case STAR:
			return l * r
	}

	// Unreachable.
	return nil
}
func (i *Interpreter) visitUnaryExpr (expr Unary) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case BANG:
		return !i.isTruthy(right)
	case MINUS:
			d, ok := right.(string)
			if ok {
				r, _ := strconv.ParseFloat(d, 64)
				return - r
			}
	}

	// Unreachable.
	return nil
}