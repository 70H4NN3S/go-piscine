// ARRAY - fixed size, value type
var arr [5]int
arr2 := [3]string{"a", "b", "c"}
arr3 := [...]int{1, 2, 3} // compiler infers size

// SLICE - dynamic, reference type
var a []int // nil SLICE
s2 := []int{1, 2, 3}
s3 := make([]int, 5)     //len=5, cap=5
s4 := make([]int, 3, 10) // len=3, cap=10

// Slicing - shares the backing array!
a := []int{0, 1, 2, 3, 4, 5}
b := a[1:4] // [1, 2, 3], len=3, cap=5
b[0] = 99   // modifies 'a' too!

// append may allocate a new backing array
a5 := []int{1, 2, 3}
a5 = append(a5, 4, 5)
a5 = append(a5, []int{6, 7}...)
