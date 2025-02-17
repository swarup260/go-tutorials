package main

import (
	"fmt"
)

func PointerError() {

	// define a pointer
	var p *int

	fmt.Println(&p) // print the address of the pointer
	fmt.Println(p) // print the value of the pointer


}