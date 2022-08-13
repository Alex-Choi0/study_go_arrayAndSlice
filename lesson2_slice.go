package main

import (
	"fmt"
)

func main() {
	// 슬라이스 생성하기
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{}
	slice2 = append(slice1, 77, 88)
	fmt.Println(slice1) // 해당 슬라이스1 출력
	fmt.Println(slice2) // 해당 슬라이스1 출력

	for i, v := range slice1 {
		fmt.Println(i, v) // 인덱스와 슬라이스 값을 1줄씩 출역
	}

	var x int = 3
	fmt.Printf("%d인덱스의 슬라이스값은 : %d\n", x, slice1[x])

	slice1[x] = 10
	fmt.Printf("%d인덱스에 변경된 슬라이스값은 : %d\n", x, slice1[x])

	fmt.Printf("슬라이스 길이 : %d, 슬라이스의 용량 : %d\n", len(slice1), cap(slice1))
	fmt.Printf("원소의 타입 : %T\n", slice1[0])

	slice1[0] = 70

	slice1 = append(slice1, 6, 7)
	fmt.Printf("slice1 : %v\n", slice1)
	fmt.Printf("cap : %d, len :%d\n", cap(slice1), len(slice1))

	fmt.Printf("slice2 : %v\n", slice2)
	fmt.Printf("cap : %d, len :%d\n", cap(slice2), len(slice2))

}
