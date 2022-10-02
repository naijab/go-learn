package main

import (
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here

	m := make(map[int32]int32)

	for _, valAr := range ar {
		m[valAr] += 1
	}

	var result int32 = 0

	for key, valM := range m {
		fmt.Printf("key[%d], val: %d\n", key, valM)
		if valM%2 == 1 {
			result++
		}
	}

	return result
}

func main() {
	const n int32 = 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	result := sockMerchant(n, ar)
	fmt.Println("actual case 1: ", result)

	const expected int32 = 3
	fmt.Println("expected case 1: ", expected)

	fmt.Println("is pass: ", expected == result)
	fmt.Println("====================")
	fmt.Println("")

	const n2 int32 = 10
	ar2 := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}

	result2 := sockMerchant(n2, ar2)
	fmt.Println("actual case 2: ", result2)

	// TODO: Still not pass
	const expected2 int32 = 4
	fmt.Println("expected case 2: ", expected2)

	fmt.Println("is pass: ", expected2 == result2)
	fmt.Println("====================")
}
