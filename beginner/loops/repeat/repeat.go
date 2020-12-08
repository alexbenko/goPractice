package repeat

func Repeat(character string, totalLoops int) string {
	output := ""
	for i := 0; i < totalLoops; i++ {
		output = output + character
	}

	return output
}
