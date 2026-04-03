// Package arrays is a simple package to test arrays and write a SUM function
package arrays

func Sum(arr []int) (res int) {
	for i := range arr {
		res += arr[i]
	}
	return
}
