package commands

import (
	"fmt"
	"strconv"
)

func Add(value1 string, keyword string, value2 string) {
	a, _ := strconv.Atoi(value1)
	b, _ := strconv.Atoi(value2)
	if keyword == "to" {
		fmt.Println(a + b)
	}
}
