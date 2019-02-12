package main

import (
	"fmt"
	"main/greeting"
)

type subString string

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

	var p1 = greeting.Person{"isaac", "28 South 9", "347-563-5668"}
	p2 := greeting.Person{PhoneNumber: "212-444-9911", Name: "josh"}
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

	//	functions
	greeting.PrintPerson(p1)
	greeting.PrintPersons(p1, p2, p1, p2, p1, p2)

	// if statements.
	fmt.Println("\n\nif statements.")
	greeting.PrinterWithExclamationMark("message 1 ", true)
	greeting.PrinterWithExclamationMark("message 2 ", false)

	greeting.PrinterWithCharacter1("0 message 3", 0)
	greeting.PrinterWithCharacter1("1 message 4", 1)
	greeting.PrinterWithCharacter1("2 message 5", 2)
	greeting.PrinterWithCharacter1("3 message 6", 3)
	greeting.PrinterWithCharacter1("3 message 7", 4)
	greeting.PrinterWithCharacter1("default message 8", 5)

	greeting.PrinterWithCharacter2("0 message 8", 0)
	greeting.PrinterWithCharacter2("1 message 9", 1)
	greeting.PrinterWithCharacter2("2 message 10", 2)
	greeting.PrinterWithCharacter2("3 message 11", 3)
	greeting.PrinterWithCharacter2("3 message 12", 4)
	greeting.PrinterWithCharacter2("default message 13", 5)

	fmt.Println("\nswitch on types")
	greeting.CheckType(1)
	greeting.CheckType("aaa")
	greeting.CheckType('d')
	var x float32 = 1.1
	greeting.CheckType(x)
	greeting.CheckType(1.12)
	var y subString = "a subString string"
	greeting.CheckType(y)

	fmt.Println("\nLoops")

	i1 := 0
	for { // like while (true) { // }
		if i1 == 5 {
			break
		}
		fmt.Println("for Loop", i1)
		i1++

	}

	for i2 := 0; i2 < 5; i2++ {
		fmt.Println("for Loop", i2)
	}
	var slice1 map[string]string
	slice1 = make(map[string]string)
	slice1["aaa"] = "A1"
	slice1["bbb"] = "B2"
	slice1["ccc"] = "C3"
	fmt.Println(slice1)

	toPrint1, toPrint2 := greeting.FindSlice(1)
	fmt.Println(toPrint1, toPrint2)
	toPrint1, toPrint2 = greeting.FindSlice(3)
	fmt.Println(toPrint1, toPrint2)

	fmt.Println("\n\narrays and Slices")

	var array1 [3]int
	array1[0] = 111
	array1[1] = 222
	array1[2] = 333
	fmt.Println(array1, len(array1))

	var slice3 []int
	slice3 = make([]int, 3, 5)
	slice3[0] = 111
	slice3[1] = 222
	slice3[2] = 333
	//slice3[3] = 444
	fmt.Println(slice3, len(slice3))

	slice4 := []int{111, 222, 333, 444}
	fmt.Println(slice4, len(slice4))
	slice5 := slice4[1:3] // [1:] go till the end, [:3] star from the beginning.
	fmt.Println(slice5, len(slice5))

	slice4 = append(slice4, 555)
	fmt.Println(slice4, len(slice4))
	slice4 = append(slice4, slice3...) // combine 2 arrays
	fmt.Println(slice4, len(slice4))

	slice4 = append(slice4[:5], slice4[6:]...) // remove slice4[5]
	fmt.Println(slice4, len(slice4))

	PeoplesList := greeting.Peoples{
		{"Isaac0 Weingarten", "28 South 9", "718-599-0720"},
		{"Isaac1 Weingarten", "29 South 9", "718-599-0721"},
		{"Isaac2 Weingarten", "30 South 9", "718-599-0722"},
		{"Isaac3 Weingarten", "31 South 9", "718-599-0723"},
	}
	fmt.Println(PeoplesList[0].Combine(), "\n", PeoplesList[1].Combine(), "\n", PeoplesList[2].Combine(), "\n", PeoplesList[3].Combine())

	PeoplesList[0].RenameName("Isaac")
	fmt.Println(PeoplesList[0].Combine())

	greeting.RenameToFrog(&PeoplesList[1])
	fmt.Println(PeoplesList[1].Combine())
}
