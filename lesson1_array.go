package main

import (
	"fmt"
)

func main() {
	// 배열 생성하기
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr) // 해당 배열 출력

	for i, v := range arr {
		fmt.Println(i, v) // 인덱스와 배열값을 1줄씩 출역
	}

	var x int = 3
	fmt.Printf("%d인덱스의 배열값은 : %d\n", x, arr[x])

	arr[x] = 10
	fmt.Printf("%d인덱스에 변경된 배열값은 : %d\n", x, arr[x])

	fmt.Printf("배열 길이 : %d, 배열의 용량 : %d\n", len(arr), cap(arr))
	fmt.Printf("원소의 타입 : %T\n", arr[0])
}
