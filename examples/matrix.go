package main

import(
	"fmt"
)

func main(){

	// Use closure to set variables as functions

	a := func (x int, y int) int{
		return x + y
	}
	d := func (x int, y int) int{
		return x / y
	}
	m := func (x int, y int) int{
		return x * x
	}
	s := func (x int, y int) int{
		return x - y
	}

	// Initialize Matrix we want to solve

	matrix := [][]interface{}{
		{a, 4, 5, 0},
		{d, 0, 3, 0},
		{m, 4, 8, 0},
		{s, 1, 4, 0}}

	// Use type assertion on empty interface values to calcutate 4th column

	for index, value := range matrix{
		matrix[index][3] = value[0].((func(int, int) int))(value[1].(int), value[2].(int))
	}

	fmt.Println(matrix)
}