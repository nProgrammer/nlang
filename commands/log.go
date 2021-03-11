package commands

import (
	"fmt"
)

func Log(data string) {
	data = data[1 : len(data)-1]
	fmt.Println(data)
}
