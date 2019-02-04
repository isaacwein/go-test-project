package main

import "fmt"

type subString string
type Person struct {
	name        string
	address     string
	phoneNumber string
}

func addLineBreaks1(name, address, phoneNumber string) string {
	return name + "\n" + address + "\n" + phoneNumber
}
func addLineBreaks2(name, address, phoneNumber string) (string, string) {
	return "string1\n" + name + "\n" + address + "\n" + phoneNumber,
		"string2\n" + name + "\n" + address + "\n" + phoneNumber
}
func addLineBreaks3(name, address, phoneNumber string) (string1 string, string2 string) {
	string1 = "string1\n" + name + "\n" + address + "\n" + phoneNumber
	string2 = "string2\n" + name + "\n" + address + "\n" + phoneNumber
	return
}
func PrintPerson(person Person) {
	fmt.Println("PrintPerson from func PrintPerson with addLineBreaks")
	fmt.Println(addLineBreaks1(person.name, person.address, person.phoneNumber))

	addLineBreaks2Value1, addLineBreaks2Value2 := addLineBreaks2(person.name, person.address, person.phoneNumber)
	_, addLineBreaks2Value3 := addLineBreaks2(person.name, person.address, person.phoneNumber)
	addLineBreaks2Value4, _ := addLineBreaks2(person.name, person.address, person.phoneNumber)
	fmt.Println("addLineBreaks2", addLineBreaks2Value1, addLineBreaks2Value2, addLineBreaks2Value3, addLineBreaks2Value4)
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

	//	functions
	PrintPerson(p1)

}
