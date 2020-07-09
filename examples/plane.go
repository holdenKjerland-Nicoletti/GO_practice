package main

import "fmt"

type Plane struct {
	enginetype string
	num_people int
}


// receiever
func(p Plane) Fly() {
	fmt.Println("flying..." + p.enginetype)
}


func main(){

	plane1:=Plane{"Rolls Royce", 125}
	plane1.Fly()

}