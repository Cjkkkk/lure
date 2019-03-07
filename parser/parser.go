package parser

import(
	"awesomeProject/lexer"
	err "awesomeProject/error"
	"errors"
)
type Parser struct {
	Tokens []lexer.Token
	Current int
}

func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser)check(t lexer.TokenType ) bool{
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() lexer.Token{
	if !p.isAtEnd() {
		p.Current += 1
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool{
	return p.peek().Type == lexer.EOF
}

func (p *Parser) peek() lexer.Token {
	return p.Tokens[p.Current]
}

func (p Parser) previous() lexer.Token{
	return p.Tokens[p.Current - 1]
}

func (p *Parser)expression() Expr {
	return p.equality()
}

func (p *Parser)equality() Expr {
	expr := p.comparison()

	for p.match(lexer.BANG_EQUAL, lexer.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = Binary{Left:expr, Operator:operator, Right:right}
	}
	return expr
}

func (p *Parser) comparison() Expr{
	expr := p.addition()

	for p.match(lexer.GREATER, lexer.GREATER_EQUAL, lexer.LESS, lexer.LESS_EQUAL) {
		operator := p.previous()
		right := p.addition()
		expr = Binary{
			Left:expr,
			Operator:operator,
			Right:right,
		}
	}

	return expr
}


func (p *Parser) addition() Expr{
	expr := p.multiplication()
	for p.match(lexer.MINUS, lexer.PLUS) {
		operator := p.previous()
		right := p.multiplication()
		expr = Binary{Left:expr, Operator:operator,Right: right}
	}
	return expr
}

func (p *Parser) multiplication() Expr{
	expr := p.unary()

	for p.match(lexer.SLASH, lexer.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = Binary{Left:expr, Operator:operator, Right:right}
	}
	return expr
}


func (p *Parser) unary() Expr{
	if p.match(lexer.BANG, lexer.MINUS) {
		operator := p.previous()
		right := p.unary()
		return Unary{Operator:operator, Right:right}
	}
	if r, e := p.primary(); e != nil {
		// todo
		return nil
	} else{
		return r
	}
}

func (p *Parser) primary() (Expr, error){
	if p.match(lexer.FALSE){
		return Literal{Value:false}, nil
	}
	if p.match(lexer.TRUE) {
		return Literal{Value:true}, nil
	}
	if p.match(lexer.NIL) {
		return Literal{Value:nil}, nil
	}
	if p.match(lexer.NUMBER, lexer.STRING) {
		return Literal{Value:p.previous().Literal}, nil
	}
	if p.match(lexer.LEFT_PAREN) {
		expr := p.expression()
		p.consume(lexer.RIGHT_PAREN, "Expect ')' after expression.")
		return Grouping{Expression:expr}, nil
	}
	if p.match(lexer.IDENTIFIER) {
		return Variable{p.previous()}, nil
	}
	return nil, p.error(p.peek(), "Expect expression.")
}

func (p *Parser) consume (t lexer.TokenType, message string) (lexer.Token, error){
	if p.check(t) {
		return p.advance(), nil
	}
	return lexer.Token{}, p.error(p.peek(), message)
}

func (p *Parser) error(token lexer.Token, message string) error{
	err.Error(token, message)
	return errors.New("can't work with 42")
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
	if p.previous().Type == lexer.SEMICOLON {
		return
	}

	switch p.peek().Type {
		case lexer.CLASS:
		case lexer.FUN:
		case lexer.VAR:
		case lexer.FOR:
		case lexer.IF:
		case lexer.WHILE:
		case lexer.PRINT:
		case lexer.RETURN:
		return
	}
	p.advance()
	}
}
func (p *Parser) statement() Stmt{
	if p.match(lexer.PRINT) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() Stmt{
	value := p.expression()
	p.consume(lexer.SEMICOLON, "Expect ';' after value.")
	return Print{Expression: value}
}

func (p *Parser) expressionStatement() Stmt{
	expr := p.expression()
	p.consume(lexer.SEMICOLON, "Expect ';' after expression.")
	return Expression{Expression: expr}
}
func (p *Parser) Parse() []Stmt {
	var statements []Stmt
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
	//return p.expression()
	// todo error handling
}

func (p *Parser) declaration() Stmt{
	if p.match(lexer.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
	// synchronize();
}

func (p *Parser) varDeclaration() Stmt{
	name, _ := p.consume(lexer.IDENTIFIER, "Expect variable name.")

	var initializer Expr
	if p.match(lexer.EQUAL) {
		initializer = p.expression()
	}

	p.consume(lexer.SEMICOLON, "Expect ';' after variable declaration.")
	return Var{Name:name, Expression:initializer }
}