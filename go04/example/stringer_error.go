package main

import "fmt"

type Temperature struct{ Celsius float64 }

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C / %.1f°F", t.Celsius, t.Celsius*9/5+32)
}

type ValidationError struct{ Field, Message string }

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func main() {
	t := Temperature{34}
	e := ValidationError{"Dunno what", "THis is the message"}
	fmt.Println(t.String())
	fmt.Println(e.Error())
}
