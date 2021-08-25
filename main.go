package main

import (
	"fmt"
	//"math"
	"time"
)

type Person struct {
	Name string
	DOB  string
}

func (p Person) Age() int {
	test, _ := time.Parse("02/01/2006", p.DOB)
	duration := time.Since(test)
	return int(duration.Seconds() / 31540000)
}

func main() {
	p := Person{"Bob", "16/11/1988", 10}
	fmt.Println(p.Age())
}
