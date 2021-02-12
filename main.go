package main

import (
	"go_test_async/myPkg"
	"time"
)

func main() {
	myPkg.DoStuff()
	time.Sleep(10 * time.Second)
}
