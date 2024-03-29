package selects

import (
	"fmt"
	"time"
)

func printWithDelay(ch1, ch2 chan string, delayMs1, delayMs2 int) {
	for {
		select {
		case str1 := <-ch1:
			time.Sleep(time.Duration(delayMs1) * time.Millisecond)
			fmt.Println(str1)
		case str2 := <-ch2:
			time.Sleep(time.Duration(delayMs2) * time.Millisecond)
			fmt.Println(str2)
		}
	}
}

func BlockSelectCase() {
	// Q: does a select instantly handle all cases, or can a case block other cases?
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go printWithDelay(ch1, ch2, 0, 500)
	ch2 <- "ch2"
	ch1 <- "ch1"
	time.Sleep(time.Second)

	// STDOUT
	// > ch1
	// > ch2

	// A: writing to both ch2 and ch1 at approx the same time causes a race condition, so sometimes:
	// - "ch1" prints
	// - ~500ms of silence
	// - "ch2" prints
	// and sometimes:
	// - ~500ms of silence
	// - "ch2" prints
	// - "ch1" prints immediately after

	// If we assume that these two paths are the only expected outcomes, then we can deduce that all
	// cases run synchronous to each other and DO block one another.
}

func PreventRaceCondition() {
	// With the results of the experiment BlockSelectCase in mind,
	// Im curious if [HYPOTHESIS](the select case conditions run in the order which they are listed).
	// Above, we saw that sometimes ch1 would rec first, and that could be due to both channels getting data pushed onto
	// them BEFORE the next cycle in the select loop.
	// If we trigger a race condition by pushing to both channels "simultaneously" and the select executes
	// case conditions in the order listed, then we should see the first channel handled in the select statement handled
	// first everytime.

	ch3 := make(chan string, 1)
	ch4 := make(chan string, 1)
	go printWithDelay(ch3, ch4, 500, 0)
	ch3 <- "ch3"
	ch4 <- "ch4"
	time.Sleep(time.Second)

	ch5 := make(chan string)
	ch6 := make(chan string)
	go printWithDelay(ch5, ch6, 500, 0)
	ch5 <- "ch5"
	ch6 <- "ch6"
	time.Sleep(time.Second)

	// RESULTS:
	// After about 10 runs:
	// - 2 of 10 runs printed "ch4" before "ch3", which indicates there's still a race condition with buffered channels
	//	 *note that the frequency of the "unexpected" route in the race condition is much lower than the 50/50 split we
	//	 saw on the previous experiment (BlockSelectCase)
	// - 10 of 10 runs printed "ch5" before "ch6", which indicates unbuffered channels prevent a race condition
}
