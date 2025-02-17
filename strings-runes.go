package main

import (
	"fmt"
)

func print(sample []byte){
	fmt.Println("Println:")
	fmt.Println(sample)
	// byte-by-byte loop
	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Printf("\n")
	fmt.Println("Bytes loop quoted:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%+q ", sample[i])
	}
	fmt.Printf("\n")
	fmt.Println("Byte loop quoted:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	// A shorter way to generate presentable output for a messy string
	// is to use the %x (hexadecimal) format verb of fmt.Printf.
	// It just dumps out the sequential bytes of the string as hexadecimal digits, two per byte.
	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	// The %q (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

func StringExample() {
	// https://go.dev/blog/strings
	// String

	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	bytes := []byte{0xbd,0xb2,0x3d,0xbc,0x20,0xe2,0x8c,0x98}

	print(bytes)

	
}
