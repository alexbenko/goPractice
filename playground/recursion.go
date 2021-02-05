package main

import "fmt"

func evenMaker(num int) int { //subtracts 1 from number until it is even
	fmt.Println(num)
	if num%2 == 0 {
		return num
	}
	return evenMaker(num - 1)
}

func main() {
	fmt.Println(evenMaker(3333))
}
