package main

import (
	"fmt"
)

/*
When you have multiple consecutive parameters of the same type, 
you may omit the type name for the like-typed parameters up to \
the final parameter that declares the type.
*/
func plusPlus(a, b, c int) int {
    return a + b + c
}

// single return value * it won’t automatically return the value of the last expression.
func add(a int,b int) int {
	return a + b
}

// multiple return values
func foo(a string) (string,string) {
	return a,"extra"
}

func foobar(a string) (string,string,string) {
	return a,"extra","more"
}

// Variadic Functions
// can be called with any number of trailing arguments
func sum(nums ...int) int {
	fmt.Println(nums, " ")
	sum := 0
	for _,num := range(nums){
		sum += num
	}
	return sum
}

// Closures
// which can form closures.
//  Anonymous functions are useful when you want to define a function inline without having to name it
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}


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

	fmt.Println(add(1,2))

	_,f,g := foobar("n")
	fmt.Println(f,g)


	nums := []int{1,2,3,4}
	fmt.Println(sum(nums...))

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
