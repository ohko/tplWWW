package demo1

import (
	"fmt"
	"time"
)

// Demo1 ...
type Demo1 struct {
	start time.Time
	last  time.Time
}

// Register ...
func (o *Demo1) Register() error {
	o.start = time.Now()
	go o.do()

	return nil
}

// Status ...
func (o *Demo1) Status() string {
	return fmt.Sprintf("demo1: %s ~ %s", o.start, o.last)
}

func (o *Demo1) do() {
	for {
		time.Sleep(time.Hour)
		o.last = time.Now()
	}
}
