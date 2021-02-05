package main

import "fmt"

func vals() (int, int, int) { //how multiple return values are declared
	return 1, 2, 3
}

func main() {
	//fmt.Println("H")
	user1, user2, user3 := vals()
	fmt.Println(user1) //=> 1
	fmt.Println(user2) //=> 2
	fmt.Println(user3) //=>3

	_, _, user := vals() //if you want to ignore one or more of multiple return values
	fmt.Println(user)    //=> 3
}
