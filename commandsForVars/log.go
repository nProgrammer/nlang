package commandsForVars

import (
	"fmt"
	"os"
)

func Log(key string) {
	data := os.Getenv(key)
	fmt.Println(data)
}
