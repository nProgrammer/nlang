package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"nlang/commands"
	"nlang/commandsForVars"
	"nlang/vars"
	"regexp"
	"strings"
)

func main() {
	path := flag.String("path", "", "a path")
	flag.Parse()
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err)
	}
	code := regexp.MustCompile("[ ;]").Split(string(data), -1)
	i := 0
	fmt.Println(len(code))
	for i < len(code)-1 {
		code[i] = strings.TrimSpace(code[i])
		if code[i] == "log" {
			if code[i+1] == "%" {
				commandsForVars.Log(code[i+2])
			} else {
				commands.Log(code[i+1])
			}
			i += 2
		} else if code[i] == "add" {
			if code[i+1] == "%" {
				commandsForVars.Add(code[i+2], code[i+3], code[i+4])
			} else {
				commands.Add(code[i+1], code[i+2], code[i+3])
			}
			i++
		} else if code[i] == "wait" {
			if code[i+2] == "%" {
				commandsForVars.Wait(code[i+2])
			} else {
				commands.Wait(code[i+1])
			}
			i++
		} else if code[i] == "subtract" {
			if code[i+2] == "%" {
				commandsForVars.Subtract(code[i+2], code[i+3], code[i+4])
			} else {
				commands.Subtract(code[i+1], code[i+2], code[i+3])
			}
			i++
		} else if code[i] == "multiply" {
			if code[i+1] == "%" {
				commandsForVars.Multiply(code[i+2], code[i+3], code[i+4])
			} else {
				commands.Multiply(code[i+1], code[i+2], code[i+3])
			}
			i++
		} else if code[i] == "divide" {
			if code[i+1] == "%" {
				commandsForVars.Divide(code[i+2], code[i+3], code[i+4])
			} else {
				commands.Divide(code[i+1], code[i+2], code[i+3])
			}
			i++
		} else if code[i+1] == "<-" {
			vars.VariableString(code[i], code[i+2])
			i += 2
		} else {
			i++
		}
	}
}
