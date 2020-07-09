package main

import "fmt"

func Cool(a int, b...string){ // the first parameter is an int, the rest of the parameters are strings, as many as we want

	fmt.Println(a)

	for index, value:=range b { // range is an iterator function
		fmt.Println(index, value)
	}

}

func main(){
	Cool(44, "x", "y")
}