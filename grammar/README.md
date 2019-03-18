# grammar
```BNF=
program  → statement* EOF;

declaration → varDecl
            | statement ;
varDecl → "var" IDENTIFIER ( "=" expression )? ";" ;            
statement → exprStmt
          | printStmt
          | forStmt
          | whileStmt
          | ifStmt
          | block ;
exprStmt  → expression ";" ;
printStmt → "print" expression ";" ;
forStmt -> "for" "(" (varDecl | exprStmt | ";")
                expression? ";"
                expression? ")" statement;
whileStmt -> "while" "(" expression ")" statement
ifStmt -> "if" "("expression")" statement "else" statement
block     → "{" declaration* "}" ;

expression → assignment ;
assignment → IDENTIFIER "=" assignment
           | logic_or ;
logic_or -> logic_and ( "or" logic_and )*;
logic_and -> equality ( "and" equality )*;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → addition ( ( ">" | ">=" | "<" | "<=" ) addition )* ;
addition       → multiplication ( ( "-" | "+" ) multiplication )* ;
multiplication → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "false" | "true" | "nil"
               | "(" expression ")" ;
```