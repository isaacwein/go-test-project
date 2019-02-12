package greeting

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person struct {
	Name        string
	Address     string
	PhoneNumber string
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
	fmt.Println(addLineBreaks1(person.Name, person.Address, person.PhoneNumber))

	addLineBreaks2Value1, addLineBreaks2Value2 := addLineBreaks2(person.Name, person.Address, person.PhoneNumber)
	_, addLineBreaks2Value3 := addLineBreaks2(person.Name, person.Address, person.PhoneNumber)
	addLineBreaks2Value4, _ := addLineBreaks3(person.Name, person.Address, person.PhoneNumber)
	fmt.Println("addLineBreaks2", addLineBreaks2Value1, addLineBreaks2Value2, addLineBreaks2Value3, addLineBreaks2Value4)
}

func PrintPersons(person ...Person) {
	fmt.Println("PrintPersons ", len(person))
	fmt.Println(person)
}

func PrinterWithExclamationMark(message string, bool bool) {
	if ex := "!"; bool {
		fmt.Println("printing a message with a " + ex + " \"" + message + ex + "\"")
	} else {
		fmt.Println("printing a message without a " + ex + " \"" + message + "\"")
	}

}
func PrinterWithCharacter1(message string, int int) {

	// no need for break by default it is using the break to undo it use fallthrough
	switch int {
	case 0:
		fmt.Println(message + "?")
	case 1:
		fmt.Println(message + "!")
	case 2, 3:
		fmt.Println(message + ";")
	default:
		fmt.Println(message + ".")
	}

}
func PrinterWithCharacter2(message string, int int) {

	// no need for break by default it is using the break to undo it use fallthrough
	switch {
	case int == 0:
		fmt.Println(message + "?")
	case int == 1:
		fmt.Println(message + "!")
	case int == 2, int == 3:
		fmt.Println(message + ";")
	default:
		fmt.Println(message + ".")
	}

}

func CheckType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println(x, reflect.TypeOf(x), "int")
	case string:
		fmt.Println(x, reflect.TypeOf(x), "string")
	case bool:
		fmt.Println(x, reflect.TypeOf(x), "boolean")
	case float32:
		fmt.Println(x, reflect.TypeOf(x), "float32")
	case float64:
		fmt.Println(x, reflect.TypeOf(x), "float64")
	default:
		fmt.Println(x, reflect.TypeOf(x), "unknown")
	}
}

func FindSlice(num int) (Person, bool) {
	slice2 := map[int]Person{
		0: {"Isaac", "452 Broadway", "212-444-9911 #1079"},
		1: {"Shlomy", "452 Broadway", "212-444-9911 #1076"},
		2: {"Josh", "452 Broadway", "212-444-9911 #1024"},
	}
	delete(slice2, 1)

	fmt.Println(slice2)
	for k, v := range slice2 {
		fmt.Println("slice1", "[", strconv.Itoa(k), "] ", v)
	}
	if value, ex := slice2[num]; ex {
		return value, ex
	} else {
		return Person{"", "", ""}, ex
	}
}

type Renamable interface {
	RenameName(newName string)
}
type People struct {
	Name        string
	Address     string
	PhoneNumber string
}

func (people People) Combine() string {
	return people.Name + " - Address: " + people.Address + " - PhoneNumber: " + people.PhoneNumber
}
func (people *People) RenameName(r string) {
	people.Name = r
}

type Peoples []People

func RenameToFrog(r *People) {
	r.RenameName("frog")
	fmt.Println("pppp", r)
}
