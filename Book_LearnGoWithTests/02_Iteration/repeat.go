// Package iteration shows basic usage of iteration in golang
package iteration

func Repeat(str string) (result string) {
	for range 5 {
		result += str
	}
	return
}
