package commands

import (
	"strconv"
	time2 "time"
)

func Wait(time string) {
	s, _ := strconv.Atoi(time)
	i := 0
	for i < s {
		time2.Sleep(1 * time2.Second)
		i++
	}
}
