package main

import (
	"fmt"
	"sort"
)

var phoneBook = make(map[string][]string) // name -> list of phone numbers

func main() {
	Add("Johannes", "+1234123123")
	Add("Gabriel", "088347534")
	Add("John", "+456456456456")
	Add("Johannes", "+345345345")
	Add("John Doe", "+9834534534")
	ListAll()
	fmt.Println("Johannes: ", Lookup("Johannes"))
	fmt.Println("John Doe", Lookup("John Doe"))
	Add("John Doe", Lookup("John")[0])
	Remove("John")
	fmt.Println(Lookup("John"))
	Add("John Doe", "+034958034985034")
	ListAll()
}

func Add(name, phone string) {
	phoneBook[name] = append(phoneBook[name], phone)
}

func Remove(name string) {
	delete(phoneBook, name)
}

func Lookup(name string) []string {
	return phoneBook[name]
}

func ListAll() {
	keys := make([]string, 0, len(phoneBook))
	for contact := range phoneBook {
		keys = append(keys, contact)
	}
	sort.Strings(keys)
	fmt.Println("------------")
	for _, contact := range keys {
		fmt.Println(contact, phoneBook[contact])
	}
	fmt.Println("-----------")
}
