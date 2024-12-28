package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println


type Counter struct{
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment(){
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter)String() string{
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main(){
	// var c Counter
	// pl(c.String())
	// c.Increment()
	// pl(c.String())
	
	c := &Counter{}
	pl(c.String())
	c.Increment()
	pl(c.String())
}