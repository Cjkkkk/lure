package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var hadError bool = false


func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println("[line " + string(line) + "] Error" + where + ": " + message)
	hadError = true
}

func run(content string)  {
	if hadError {
		os.Exit(65)
	}
	var s scanner
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
		var input string
		fmt.Scanln(&input)
		run(input)
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