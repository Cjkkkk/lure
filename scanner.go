package main

type tokenType int

const(
	// Single-character tokens.
	LEFT_PAREN tokenType = iota
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

	// One or two character tokens.
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
type token struct {
	Type tokenType
	lexeme string
	literal string
	line int
}

func (t *token) toString() string{
	return string(t.Type) + " " + t.lexeme + " " + t.literal
}

type scanner struct {
	source string
	tokens []token
	current int //0
	start int //0
	line int //1
}

func (s *scanner) isAtEnd() bool{
	return s.current >= len(s.source)
}

func (s *scanner) advance() byte{
	s.current += 1
	return s.source[s.current-1]
}

func (s *scanner) addToken(t tokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token{t, text,nil, s.line})
}
func (s *scanner) scanToken() {
	c := s.advance()
	switch c {
		case '(': s.addToken(LEFT_PAREN)
		case ')': s.addToken(RIGHT_PAREN)
		case '{': s.addToken(LEFT_BRACE)
		case '}': s.addToken(RIGHT_BRACE)
		case ',': s.addToken(COMMA)
		case '.': s.addToken(DOT)
		case '-': s.addToken(MINUS)
		case '+': s.addToken(PLUS)
		case ';': s.addToken(SEMICOLON)
		case '*': s.addToken(STAR)
		default: error(s.line, "Unexpected character.")
	}
}
func (s *scanner)scanTokens() []token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token{EOF,"",nil, s.line})
	return s.tokens
}

