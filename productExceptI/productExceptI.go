/*
Given an array of integers, return a new array such that each element at index i of the new array
is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be
 [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

 must be better way than this!!
*/

package main

import "fmt"

func mult(in []int) int {
	sum := 1
	for i := 0; i < len(in); i++ {
		sum *= in[i]
	}
	return sum
}

func productExceptI(data []int) []int {
	var res []int
	for i := 0; i < len(data); i++ {
		a, b := data[:i], data[i+1:]
		var c []int
		c = append(c, a...)
		c = append(c, b...)
		res = append(res, mult(c))
	}
	return res
}

func main() {

	v := productExceptI([]int{1, 2, 3, 4, 5})
	v2 := productExceptI([]int{3, 2, 1})
	fmt.Println(v)
	fmt.Println(v2)

}
