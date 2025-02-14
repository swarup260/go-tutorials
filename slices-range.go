package main

import (
	"fmt"
)
func foo(a string) (string,string) {
	return a,"extra"
}

func SliceRange() {
	var a []string               // define a slice
	b := []string{"a", "b", "c"} // define a slice with initial values
	/*
		func make([]T, len, cap) []T
		** cap is optional argument
	*/
	c := make([]byte, 5, 5) // define a slice using make function
	d := make([]byte, 5)    // define a slice using make function

	fmt.Printf("A Slices : %s | Slices Len : %d |  Slice Cap : %d \n", a, len(a), cap(a))
	fmt.Printf("B Slices : %s | Slices Len : %d |  Slice Cap : %d \n", b, len(b), cap(b))
	fmt.Printf("C Slices : %s | Slices Len : %d |  Slice Cap : %d \n", c, len(c), cap(c))
	fmt.Printf("D Slices : %s | Slices Len : %d |  Slice Cap : %d \n", d, len(d), cap(d))


	e := make(map[string]int) // define a map 
	n := map[string]int{"foo": 1, "bar": 2} // define a map with value initialize 
	// set the value
	e["k1"] = 1
	e["k2"] = 1

	fmt.Println(e,n)

	fmt.Println(e["k3"]) // if value is not present then zero values

	_,prs := e["k3"]
	_,prs1 := e["k2"] // check if the key is present or not . the variable prs1 return bool value 

	fmt.Println(prs,prs1)
	fmt.Println(len(e)) // len of map
	delete(e,"k2")
	fmt.Println(e) // len of map
	clear(e)
	fmt.Println(e) // len of map
	
	fmt.Println(foo("a"))

	_,extra := foo("b")

	fmt.Println(extra)

}
