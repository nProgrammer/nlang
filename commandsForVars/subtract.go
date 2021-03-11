package commandsForVars

import (
	"fmt"
	"strconv"
)

func Subtract(value1 string, keyword string, value2 string) {
	a, _ := strconv.Atoi(value1)
	b, _ := strconv.Atoi(value2)
	if keyword == "from" {
		fmt.Println(b - a)
	}
}
