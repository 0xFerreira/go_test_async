package myPkg

import (
	"fmt"
	"time"
)

func doAsyncStuff() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Ping #", i)
	}
}

func DoStuff() {
	go doAsyncStuff()
}
