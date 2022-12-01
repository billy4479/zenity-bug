package main

import (
	"time"

	"github.com/ncruces/zenity"
)

func main() {
	prog, _ := zenity.Progress()
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100)
		prog.Value(i)
	}

	prog.Complete()
}
