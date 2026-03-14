package main

import "fmt"

func main() {
	registry := Registry{}
	registry.Add("Alpha", 2, 3, 4, 2, 3, 4, 5, 3, 2, 1)
	registry.Add("Beta", 9, 8, 6, 5, 6)
	registry.Add("Gamma", 10, 10, 10, 10, 9)
	registry.Add("Theta", 10, 10)
	registry.Add("Delta", 5, 6, 7, 8, 9, 10, 10, 10)
	registry.Add("Epsylon", 9, 8, 7, 6, 5, 4)
	registry.Add("Zetta", 1, 2, 3, 4, 5, 6)
	registry.Add("Phi", 9, 8, 7, 6, 7, 8, 9)
	registry.Add("Pi", 3, 3, 4, 3, 3, 3, 3, 3, 3)
	registry.Add("Omega", 10, 1, 10, 10, 10)

	for i, student := range registry.TopN(10) {
		fmt.Printf("%d. Student: %s; Average: %0.2f; Grade: %s\n", i+1,
			student.Name, student.Average(), student.Grade())
	}
}
