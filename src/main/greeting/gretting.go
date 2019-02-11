package greeting

import "fmt"

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
