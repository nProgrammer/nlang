package commands

import (
	"fmt"
	"regexp"
	"strings"
)

func Log(data string) {
	text := regexp.MustCompile("[/]").Split(data, -1)
	textS := strings.Join(text, " ")
	textS = textS[1 : len(textS)-1]
	fmt.Println(textS)
}
