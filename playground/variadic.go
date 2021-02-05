package main

import "fmt"

//variadic functions takes in a varity of inputs, ie kwargs from python
func sum(nums ...int) int { // ... tells compiler that this function can take any amount of ints in
	sum := 0
	for _, num := range nums { //first value of range loop is the current index, since it wont need it declare it as underscore
		sum += num
	}
	return sum
}
func main() {
	fmt.Println(sum(100, 200, 300)) //can pass in values like this
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//will also work on iterable data types
	numbers := []int{1, 2, 3, 4, 5, 6, 10, 15, 20, 25, 30}
	fmt.Println(sum(numbers...)) //to pass in iterable, use ... syntax
}
