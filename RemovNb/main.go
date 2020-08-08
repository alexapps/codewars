package main

import "fmt"

/**
A friend of mine takes a sequence of numbers from 1 to n (where n > 0).
Within that sequence, he chooses two numbers, a and b.
He says that the product of a and b should be equal to the sum of all numbers in the sequence, excluding a and b.
Given a number n, could you tell me the numbers he excluded from the sequence?
*/

func main()  {
	fmt.Println("Test", RemovNb(26))
	
}

func RemovNb(m uint64) [][2]uint64 {
	var retValue [][2]uint64
	sum := sumAllExpectTwoNumbers(m) 
	// your code
	for i := uint64(1); i <= m; i++ {
		for j := uint64(1); j <= m; j++ {
			if (sum - (i + j))  == i * j {
				// fmt.Println(fmt.Sprintf("a = %d, b = %d", i, j))
				retValue = append(retValue, [2]uint64{i,j})
			}
		}
	}

	return retValue
}

func sumAllExpectTwoNumbers(m uint64) (sum uint64) {
	// For is too long
	// for i := uint64(1); i <= m; i++ {
	// 	sum = sum + i
	// }
	sum = (m*(m + 1))/2  
	return
}

