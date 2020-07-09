package main

import(
	"fmt"
	"os"
	"strconv"
)

func Fibonacci(start_point int, end_point int) {

	sequence := []int{0,1} // Fibonacci sequence starts with 0 and 1, the rest are calcluated

	in_range := []int{} // This will be slice of fib values in the desired range


	// This loop calculates fibonacci sequence up to end_point
	for true{
		next_value := sequence[len(sequence) -1] + sequence[len(sequence) -2]

		if next_value <= end_point{
			sequence = append(sequence, next_value)
		}else{ // If next value is greater than end_point we can break
			break
		}
	}

	// Add values that are greater than start point to in_range slice
	for _, value := range sequence {
		if value >= start_point{
			in_range = append(in_range, value)
		}
	}

	fmt.Println(in_range) // print slice
}

func main(){

	// get command line arguments and convert from string to int using os and strconv

	var start_point, _ = strconv.Atoi(os.Args[1])
	var end_point, _ = strconv.Atoi(os.Args[2])
	
	Fibonacci(start_point, end_point)
}