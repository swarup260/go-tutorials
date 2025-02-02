package main

import (
	"fmt"
)

func SliceRange() {
	// https://go.dev/blog/slices-intro
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

	//copy slice 
	e := make([]byte,len(d))
	copy(e,d)

	fmt.Printf("D Slices : %s | Slices Len : %d |  Slice Cap : %d \n", e, len(e), cap(e))
	
	// append
	e = append(e,1,2,3)
	fmt.Println(e)

	// loop 
	for Byte := range(e) {
		fmt.Println(e[Byte])
	}




	

}
