package main


import (
    "fmt"
    // "time"
)


var pl = fmt.Println



// type Counter struct{
//     total int
//     lastUpdated time.Time
// }

// func (c *Counter)Increment(){
//     c.total++
//     c.lastUpdated = time.Now()
// }

// func (c Counter)String() string{
//     return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
// }

// type Increment interface{
//     Increment()
// }

// type Incrementer interface{
// 	Incrementer()
// }

// var myStringer fmt.String
// var myIncrementer Incrementer
// pointerCounter := &Counter{}
// valueCounter := Counter{}

// myString = pointerCounter
// myStringer = valueCounter

// myIncrementer = pointerCounter
// myIncrementer = valueCounter




type Doubler interface{
    Double()
}

type DoubleInt int

func (d *DoubleInt) Double(){
    *d = *d * 2
}

type DoubleIntSlice [] int

func (d DoubleIntSlice)Double(){
    for i := range d {
        d[i] = d[i] *2
    }
}




func DoubleCompare(d1, d2 Doubler){
    pl(d1 == d2)
}

func main(){
    // var pointerCounter *Counter
    // fmt.Println(pointerCounter == nil) // prints true
    // var incrementer Increment
    // fmt.Println(incrementer == nil) // prints true
    // incrementer = pointerCounter
    // fmt.Println(incrementer == nil) // false

    var di DoubleInt = 10
    var di2 DoubleInt = 10
    var dis = DoubleIntSlice{1, 2, 3}
    var dis2 = DoubleIntSlice{1, 2, 3}

    DoubleCompare(&di, &di2)
    DoubleCompare(&di, dis)
    DoubleCompare(dis, dis2)

}