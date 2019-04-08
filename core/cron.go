package dbtools

import (
	"sync"
	"time"
)

type scheduled interface {
	nextRun() (time.Duration, error)
}

type Job struct {
	fn        func()
	Quit      chan bool
	SkipWait  chan bool
	err       error
	schedule  scheduled
	isRunning bool
	sync.RWMutex
}

type Recurrent struct {
	units  int
	period time.Duration
	done   bool
}

type Daily struct {
	hour int
	min  int
	sec  int
}

type Weekly struct {
	day time.Weekday
	d   Daily
}
