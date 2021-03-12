package commandsForVars

import (
	"fmt"
	"os"
	"strconv"
)

func Multiply(value1 string, keyword string, value2 string) {
	value1 = os.Getenv(value1)
	a, _ := strconv.Atoi(value1)
	b, _ := strconv.Atoi(value2)
	if keyword == "by" {
		fmt.Println(a * b)
	}
}
