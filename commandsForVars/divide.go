package commandsForVars

import (
	"fmt"
	"strconv"
)

func Divide(value1 string, keyword string, value2 string) {
	a, _ := strconv.Atoi(value1)
	b, _ := strconv.Atoi(value2)
	if keyword == "by" {
		fmt.Println(a / b)
	}
}
