package commandsForVars

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Log(key string) {
	data := os.Getenv(key)
	text := regexp.MustCompile("[/]").Split(data, -1)
	textS := strings.Join(text, " ")
	fmt.Println(textS)
}
