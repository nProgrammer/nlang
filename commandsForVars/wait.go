package commandsForVars

import (
	"os"
	"strconv"
	time2 "time"
)

func Wait(time string) {
	t := os.Getenv(time)
	s, _ := strconv.Atoi(t)
	i := 0
	for i < s {
		time2.Sleep(1 * time2.Second)
		i++
	}
}
