package lexer

import (
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
