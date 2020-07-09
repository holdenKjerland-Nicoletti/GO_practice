package main

import "fmt"

func Swap(a int, b int) (int, int){
	// return b, a
	temp:= a
	a = b
	b = temp
	return a, b
}

func Swap_pointers(a *int, b *int){
	temp:= *a
	*a = *b
	*b = temp
}


/* can't do short syntax assignment for global variables

e := 10 can't do this
*/

func main(){

	var c int = 5 // we don't have to do this, the data type can be implied
	var d = 10 
	
	
	e:=5
	f:=10 // less type heavy way of doing above


	Swap(c, d) // this won't swap the values

	fmt.Println(c, d)

	Swap_pointers(&e, &f)

	fmt.Println(e, f)
	// fmt.Println(Swap(5,10))
}