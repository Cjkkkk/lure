package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	lexer "awesomeProject/lexer"
)

var hadError = false


func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println("[line " + strconv.Itoa(line) + "] Error" + where + ": " + message)
	hadError = true
}

func run(content string)  {
	if hadError {
		os.Exit(65)
	}
	s := lexer.Scanner{content, []lexer.Token{}, 0 ,0 , 1}
	s.ScanTokens()
	for _, t := range s.Tokens {
		fmt.Println(t.Type)
	}
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
		hadError = false
	}
}
func main(){
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: lox [script]")
	}else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPromot()
	}
}