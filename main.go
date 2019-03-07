package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	error "awesomeProject/error"
	lexer "awesomeProject/lexer"
	parser "awesomeProject/parser"
)


func run(content string)  {
	if error.HadError {
		os.Exit(65)
	}
	if error.HadRuntimeError{
		os.Exit(70)
	}
	s := lexer.Scanner{content, []lexer.Token{}, 0 ,0 , 1}
	s.ScanTokens()
	p := parser.Parser{s.Tokens , 0}
	stmts := p.Parse()
	//a := interpreter.AstPrinter{}
	b := parser.MakeInterpreter()
	b.Interpret_(stmts)
}
func runFile(filename string) {
	fmt.Println("running file...")
	//f, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend|os.ModePerm)
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	run(string(str))
	fmt.Println("done processing file...")
}

func runPromot()  {
	for ;; {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)
		run(text)
		error.HadError = false
	}
}
func main(){
	 args := os.Args
	 //var expression = lexer.Binary{
		// lexer.Unary{
		//	 lexer.Token{lexer.MINUS, "-", nil, 1},
		//	 lexer.Literal{123},
	 //		},
		//lexer.Token{lexer.STAR, "*", nil, 1},
		// lexer.Grouping{
		//lexer.Literal{45.67},
	 //	},
	 //}

	//fmt.Println( lexer.AstPrinter{}.Print(expression))
	if len(args) > 2 {
		fmt.Println("Usage: lox [script]")
	}else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPromot()
	}
}