package main

func main() {
	rectangle := Rectangle{3.14, 7.89}
	circle := Circle{10}
	triangle := Triangle{4, 5, 6}

	rectangle.PrettyPrint()
	circle.PrettyPrint()
	triangle.PrettyPrint()
}
