package main

import (
	"fmt"
)

func VariablesArray() {
	fmt.Println("Hello WOrld")

	// single var
	var a = 1
	// multiple variables
	var b, c = 2, 3
	/* Go will infer the type of initialized variables. */
	var d = true

	fmt.Println(a, b, c, d)
	// shorthand
	e := "Golang"

	fmt.Println(e + " Programming")
	// Variables declared without a corresponding initialization are zero-valued for numeric
	var f int
	var g string // empty string
 	var h float32

	fmt.Println(f, g, h)

	// constant
	const i = "HEL"

	/* for is Go’s only looping construct */
	/* 
	for without a condition will loop repeatedly until you 
	break out of the loop or return from the enclosing function.
	*/
	// basic loop
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// range loop 
	for k := range(5) {
		fmt.Println(k)
	}

	// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
	var l = 9
	if l > 5 {
		fmt.Println("HE")
	}else {
		fmt.Println("HO")
	}

	// switch case

	var m = 1
	switch m {
	case 1,2: // multiple expressions
		fmt.Println("M")
	case 3: 
		fmt.Println("N")
	default:
		fmt.Println("O")
	}

	// array 
	// By default an array is zero-valued, which for ints means 0s.
	
	var n [5]int //declare an array with 5 elements 
	o := [5]int{1,2,3,4,5} // Use this syntax to declare and initialize an array in one line.
	p := [...]int{1,2,3,4,5} // the compiler count the number of elements for you
	q := [...]int{1,3:45,5} // If you specify the index with :, the elements in between will be zeroed
	var r [5][5]int //2D array
	s := [2][3]int{  //2D array
		{1,2,3},
		{1,2,3},
	} 

	fmt.Println(n)
	// assign value
	n[0] = 1
	n[4] = 1
	fmt.Println(n)
	// len
	fmt.Println("Array Len : ",len(n))
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)


	var name string = "Alex"
    var age int = 25
    var height float64 = 5.9
    var isStudent bool = true

    fmt.Printf("Name: %s\n", name)         // %s for strings
    fmt.Printf("Age: %d\n", age)           // %d for integers
    fmt.Printf("Height: %.1f\n", height)   // %.1f for floating-point numbers
    fmt.Printf("Is Student: %t\n", isStudent) // %t for booleans

}
