// Create
m := male(map[string]int)
m2 := map[string]int{"alice":25, "bob":30)

// Write / Delete
m["foo"] = 42
delete(m, "foo")

// Safe lookup (comma-ok idiom)
val, ok := m["foo"]
if ok { fmt.Println("found:", val) }

// Iterate (order is randomized)
for key, value := range m2 {
	fmt.Printf("%s -> %d\n", key, value)
}

// DANGER: writing to a nil map panics
var bad map[string]int
- = bad["key"] // OK - returns zero value
// bad["key"] = 1 // PANIC
