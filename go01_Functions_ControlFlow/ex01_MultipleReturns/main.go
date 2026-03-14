package main

import "fmt"

func main() {
	strings := []string{
		"This is the first string",
		"42 is my favorite number",
		"because it's a Palyndrom like Sheldon said in BBT",
		"WELL, IF I REMEMBER CORRECTLY",
		"double  spaces  are  more  FUN!",
		"12312312312312312312312312312313",
	}
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "length", "upper", "lower",
		"digits", "spaces")
	for _, str := range strings {
		length, upper, lower, digits, spaces := analyze(str)
		fmt.Printf("%-10d %-10d %-10d %-10d %-10d\n",
			length, upper, lower, digits, spaces)
	}
}

func analyze(s string) (length, upper, lower, digits, spaces int) {
	for _, l := range s {
		switch {
		case l >= '0' && l <= '9':
			digits++
		case l >= 'a' && l <= 'z':
			lower++
		case l >= 'A' && l <= 'Z':
			upper++
		case l == ' ':
			spaces++
		}
		length++
	}
	return
}
