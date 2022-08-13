package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Printf("slice1 : %v, len : %d, cap : %d\n", slice1, len(slice1), cap(slice2))
	fmt.Printf("slice2 : %v, len : %d, cap : %d\n", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // including slice2 will change

	fmt.Println("After change second element")
	fmt.Printf("slice1 : %v, len : %d, cap : %d\n", slice1, len(slice1), cap(slice2))
	fmt.Printf("slice2 : %v, len : %d, cap : %d\n", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 700)

	fmt.Println("After append 700")
	fmt.Printf("slice1 : %v, len : %d, cap : %d\n", slice1, len(slice1), cap(slice2))
	fmt.Printf("slice2 : %v, len : %d, cap : %d\n", slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 10, 11, 12)
	slice2[1] = 77
	slice2[3] = 777
	fmt.Println("Add 3 elements")
	fmt.Printf("slice1 : %v, len : %d, cap : %d\n", slice1, len(slice1), cap(slice2))
	fmt.Printf("slice2 : %v, len : %d, cap : %d\n", slice2, len(slice2), cap(slice2))

}
