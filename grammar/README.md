# grammar
```BNF=
program  → statement* EOF;

statement → exprStmt
          | printStmt
          | block ;

block     → "{" declaration* "}" ;
expression → assignment ;
assignment → IDENTIFIER "=" assignment
           | equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → addition ( ( ">" | ">=" | "<" | "<=" ) addition )* ;
addition       → multiplication ( ( "-" | "+" ) multiplication )* ;
multiplication → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "false" | "true" | "nil"
               | "(" expression ")" ;
```