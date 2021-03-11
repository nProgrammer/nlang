package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./run.nlang")
	if err != nil {
		fmt.Println(err)
	}
	code := strings.Split(string(data), " ")
	fmt.Println(code[0])
}
