package main

import (
	"fmt"
	"time"
)

func Goroutines() {
	tryAccessSpawnedStack()
}

func tryAccessSpawnedStack() {
	// Q: Does a goroutine have real time access to the stack that it was spawned from?
	a := 12
	go func() {
		fmt.Println(a) // 12
		time.Sleep(100 * time.Millisecond)
		fmt.Println(a) // 11
	}()
	time.Sleep(50 * time.Millisecond)
	a = 11
	// keep main process alive
	time.Sleep(100 * time.Millisecond)

	// A: yes, the goroutine could access `a` as if that variable was assigned on the goroutine's stack, and responded
	// 		to updates to the variable as well.
}
