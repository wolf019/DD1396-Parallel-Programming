package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	var t = time.Now() // Time will be constantly updated. t is created
	for {
		time.Sleep(delay)                                                    // wait for specified time to pass
		t = time.Now()                                                       // update time
		fmt.Printf("Klockan är %02d:%02d: %s\n", t.Hour(), t.Minute(), text) // print message
	}
}

func main() {
	go Remind("Dags att äta", time.Second*3) // Seconds instead of hours...
	go Remind("Dags att arbeta", time.Second*8)
	Remind("Dags att sova", time.Second*24) // last call doesn't need a goroutine
	select {}
}
