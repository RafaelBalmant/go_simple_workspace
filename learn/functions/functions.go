package functions

import (
	"log"
	"sort"
)

func Sum(a int, b int) int {
	return a + b
}

func SumAndString(a int, b int) (int, string) {
	return a + b, "Sum"
}

func SumNomead(a, b int) (result int) {
	result = a + b
	return
}

// I can pass a slice of int and use the ... operator to pass the slice as a variadic parameter
func SumAll(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

func GenerateFunction(value int) func() int {
	return func() int {
		return value + 10
	}
}

func BinarySearch(arr []string, item string) int {
	low := 0
	high := len(arr) - 1
	sort.Strings(arr[:])
	for low <= high {
		log.Printf("Low: %d", low)
		log.Printf("High: %d", high)
		middle := (low + high) / 2
		log.Printf("Middle: %d", middle)
		try := arr[middle]
		log.Printf("Try: %s", try)
		if try == item {
			log.Printf("Item found at index: %d", middle)
			return middle
		} else if try > item {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return 0
}
