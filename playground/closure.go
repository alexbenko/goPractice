package main

import "fmt"

func increment() func() int { //function that returns a function
	//obviously this function has no realworld purpose. Just made to practice the idea of anonymous functions and closure in go
	//and way to keep track of state
	numState := 0
	return func() int {
		numState++
		return numState
	}
}
func main() {
	next := increment() //next is now a function
	fmt.Println(next()) // => 1
	next()
	next()
	next()
	fmt.Println(next()) //=>5

}
