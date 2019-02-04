package main

import (
	"fmt"
)

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
	fmt.Println(pointerMessage, pointerMessagePointer1, pointerMessagePointer2, *pointerMessagePointer1)
}
