package main

import(
	"fmt"
)

type Astruct struct {
	a int
}


// If we want the object to actually change we need to use a pointer
// For some reason with receivers, we don't need to dereference the pointer
func (obj *Astruct) MethodA(){ 
	obj.a = 10
}

type Bstruct struct {

	b int

	// obj AStruct this gets ugly
	Astruct // embedding take off the type allows us not to do obj.obj.a
}

func (obj *Bstruct) MethodB(){
	obj.b = 10
}

func main(){

	obj := Astruct{}

	obj.a = 15

	obj.MethodA()

	fmt.Println(obj.a)
}