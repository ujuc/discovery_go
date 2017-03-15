package main

import (
	"fmt"
)

func fac(n int) int {
	if n <= 0 {
		return 1
	}

	return n * fac(n-1)
}

func facItr(n int) int {
	result := 1

	for n > 0 {
		result *= n
		n--
	}

	return result
}

func facItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func ex1_6_1(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("타잔이 %v원짜리 팬티를 입고, %v원짜리 칼을 차고 노래를 한다.\n", i*10, (i+1)*10)
	}
}

func ex1_6_2(disk_num int) {
	fmt.Printf("Number of disks: %v\n", disk_num)
}

func ex1_6_3(n int) {
	p, q := 0, 1
	i := 1

	fmt.Printf("%v %v\n", p, q)
	for i <= n {
		p, q = q, p+q
		fmt.Printf("%v %v\n", p, q)
		i++
	}
}

func main() {
	fmt.Println(fac(5))
	fmt.Println(facItr(5))
	fmt.Println(facItr2(5))

	ex1_6_1(10)
	ex1_6_2(3)
	ex1_6_3(10)
}
