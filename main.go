package main

import (
	"fmt"
	"time"

	"github.com/ncruces/zenity"
)

func main() {
	prog, _ := zenity.Progress()
	// defer prog.Close()

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 50)
		if err := prog.Value(i); err != nil {
			// Dialog cancelled
			fmt.Println(err)
			break
		}
	}

	if err := prog.Complete(); err != nil {
		// Dialog cancelled
		fmt.Println(err)
	}

	<-prog.Done()
}
