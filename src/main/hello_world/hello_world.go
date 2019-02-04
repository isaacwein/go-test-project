package main

import "fmt"

type subString string
type Person struct {
	name        string
	address     string
	phoneNumber string
}

func main() {
	var message1 string
	var int1, int2, int3 int
	const int4, int5, int6 = 4, 5, 6 // no need to declare the type if the value is declared right away.
	message1 = "hello message1\n"
	message2 := "Hello, message2" // short hand for var.
	fmt.Printf(message1)
	int1, int2, int3 = 1, 2, 3
	fmt.Println(message2)
	fmt.Println(int1, int2, int3, int4, int5, int6)

	//	pointers

	pointerMessage := "this is a pointer message"
	var pointerMessagePointer1 *string = &pointerMessage
	pointerMessagePointer2 := &pointerMessage
	fmt.Println("pointers", pointerMessage, pointerMessagePointer1, pointerMessagePointer2, *pointerMessagePointer1)

	//	custom types

	var message3 subString
	message3 = "message3"
	fmt.Println("custom types", message3)

	var p1 = Person{"isaac", "28 South 9", "347-563-5668"}
	p2 := Person{phoneNumber: "212-444-9911", name: "josh"}
	fmt.Println(p1, p2)

	//const varibles

	const (
		PI   = 3.14
		Lang = "en"
		A    = iota
		B
		C
	)

	fmt.Println("const", PI, Lang, A, B, C)

}
