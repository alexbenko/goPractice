package main

import "fmt"

func main() {
	//range is pretty similar to .forEach in js with different syntax
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers { //_ allows you to declare an unused variable since the Go compiler forbids it
		sum += num
	}
	fmt.Println(sum) //=> 15

	//just like js , gives you access to the index of the current loop this time i wont use underscore
	for i, num := range numbers {
		if num == 5 {
			fmt.Println("Found 5 at index: ", i)
		}
	}

	//works on maps (objects)
	user := map[string]string{"first_name": "Alex", "last_name": "Benko"}
	for key, value := range user {
		fmt.Println("%s -> %s\n", key, value)
	}
}
