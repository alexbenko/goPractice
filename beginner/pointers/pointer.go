package main

import "fmt"

func overwrite(val int) {
	//this will not change the value passed in
	val += 2
}

//*int means that val is an integer pointer
func incrementByTwoUsingPointer(valPtr *int) {
	//this function will change the value passed in
	//this dereferences the pointer from memory to current value
	//assining a value to a dereferenced ponter changes the value at that address
	*valPtr += 2
}

func main() {
	test := 0
	fmt.Println("Start: ", test) //0

	overwrite(test)
	fmt.Println("Non Pointer Func : ", test) //0

	incrementByTwoUsingPointer(&test)
	fmt.Println("After Pointer Func : ", test) // 2

	fmt.Println("Pointer Adress : ", &test)
}
