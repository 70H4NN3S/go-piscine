package main

func main() {
	c := ConsoleLogger{}
	f := FileLogger{"logs"}

	c.Log("INFO", "ConsoleLogger works fine")
	c.Log("WARN", "About to use the FileLogger")
	f.Log("DEBUG", "Let's see if this works")
	f.Log("ERROR", "You can't divide a number by 0")
	f.Log("INFO", "Appending the third line to the log file")
}
