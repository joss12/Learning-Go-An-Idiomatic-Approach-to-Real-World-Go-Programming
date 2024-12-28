package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last update: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	pl("in doUpdateWrong", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	pl("in updateRight:", c.String())
}

func main() {
	var c Counter
	// c := &Counter{}
	doUpdateWrong(c)
	pl("in main:", c.String())
	doUpdateRight(&c)
	pl("in main:", c.String())
}
