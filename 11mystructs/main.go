package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	var dexter = Pet{"dexter", "dog", 12, behaviour{true, false}}
	fmt.Printf("dexter %+v", dexter)
}

type Pet struct {
	Name string
	Type string
	age  int
	behaviour
}

type behaviour struct {
	isFriendly bool
	softSpoken bool
}
