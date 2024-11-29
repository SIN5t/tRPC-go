package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	array := make([]int, 0, 2)

	array = append(array, 1)
	fmt.Println(&array[0])
	fmt.Println(array)

	slice := append(array, 2)
	fmt.Println(&slice[0])
	fmt.Println(slice)

	slice1 := append(array, 3)
	fmt.Println(&slice1[0])
	fmt.Println(slice1)

}
