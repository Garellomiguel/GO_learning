package main

import "fmt"

type contectInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	// contact  contectInfo
	contectInfo
}

func main() {
	// alex := person{"Alex", "Anders"}
	alex := person{firstName: "Alex", lastName: "Anders"}
	// var alex person

	alex.contectInfo = contectInfo{email: "Alex@hla", zip: 1025}

	alexPointer := &alex // Reference to memory addres
	alexPointer.updateName("Jorge")
	alex.updateName("Mauro") // no need for the last two lines, if we add the * to the func works the same
	alex.print()
}

func (pointerToPerson *person) updateName(newName string) { // if the * is on the type means that we are reciving a pointer
	(*pointerToPerson).firstName = newName // if the * is on the actual pointer, we are geting the value that is store in that pointer(memory slot)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
