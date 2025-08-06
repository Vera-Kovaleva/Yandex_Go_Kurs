package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime time.Time
	Durations []time.Duration
}

func (stopWatch *Stopwatch) Start() {
	stopWatch.StartTime = time.Now()
}

func (stopWatch *Stopwatch) SaveSplit() {
	stopWatch.Durations = append(stopWatch.Durations, time.Since(stopWatch.StartTime))
}

func (stopWatch Stopwatch) GetResults() []time.Duration {
	return stopWatch.Durations
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
