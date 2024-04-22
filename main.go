package main

import (
	"fmt"
	"go_playground/concurrency"
)

func main() {
	fmt.Println("Mutex Impl")
	concurrency.GoRoutineScript()
	fmt.Println("Channel Impl")
	concurrency.ChannelImpl()
}
