package main

import "fmt"

//slices are similar to arrays but are typed only by the elements they contain (not the number of elements) and have more built in methods.
func main() {
	slice := make([]string, 3)
	slice[0] = "F"
	slice[1] = "o"
	slice[2] = "o"
	fmt.Println(slice)

	//can call len on them to
	fmt.Println("Length of this slice: ", len(slice))

	//can add values to slice (like .push for js arrays)
	//append returns a new slice with the added data
	//appends has two required and infinite optional params
	//first param is the slice you want to add to, and the second is the value you want to add. Any other paramters passed in will get added to the slice
	slice = append(slice, "B")

	//can pass in multiple values
	slice = append(slice, "a", "r")
	fmt.Println(slice)

	//slices can be copied
	//could probably create your own method that does this cleaner
	slice_copy := make([]string, len(slice))
	copy(slice_copy, slice) //(destination,source)
	fmt.Println(slice_copy)

	//slices can be...sliced, works just like js just different syntax [low:high]
	//retrieves indexes starting with the low but excludes the high index
	sliced := slice[1:3] //slices up to the second index starting at the 1 index
	fmt.Println(sliced)  // => [ o o]

	sliced = slice[:4]  //slices up to but not including the 4th index
	fmt.Println(sliced) // => [F o o B]

	//slices can instatnly be declared with values to
	instant := []string{"F", "o", "o"}
	fmt.Println(instant) //=> [F o o]

	//slices can also be multi dimensioned
	//length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
