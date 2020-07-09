package main

import(
	"fmt"
	"os" // because we want command line arguments
	"strconv" //to convert string to int
)

var hello = []string{"Hello", "Hola", "Bon Jour", "Ciao"} //dynamic array, slice


func main(){
	var index = 1
	if len(os.Args) > 1 {

		// The underscore is for return values that we don't need. 

		// In GO often there are two returns, in this case it is an error value, that we don't need

		index, _ = strconv.Atoi(os.Args[1]) // This converts command line argument from string to integer
	}

	if index < 1 || index > len(hello){
		index = 1
	}

	fmt.Println(hello[index-1])
}