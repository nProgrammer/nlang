package main

import (
	"fmt"
	"io/ioutil"
	"nlang/commands"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./run.nlang")
	if err != nil {
		fmt.Println(err)
	}
	code := regexp.MustCompile("[ ;]").Split(string(data), -1)
	i := 0
	for i < len(code) {
		code[i] = strings.TrimSpace(code[i])
		if code[i] == "log" {
			commands.Log(code[i+1])
			i += 2
		} else if code[i] == "add" {
			commands.Add(code[i+1], code[i+2], code[i+3])
			i++
		} else if code[i] == "wait" {
			commands.Wait(code[i+1])
			i++
		} else if code[i] == "subtract" {
			commands.Subtract(code[i+1], code[i+2], code[i+3])
			i++
		} else if code[i] == "multiply" {
			commands.Multiply(code[i+1], code[i+2], code[i+3])
			i++
		} else if code[i] == "divide" {
			commands.Divide(code[i+1], code[i+2], code[i+3])
			i++
		} else {
			i++
		}
	}
}
