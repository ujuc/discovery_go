package main

import "fmt"

func print(n int) {
	for i := 1; i <= n; i++ {
		n1 := i * 10
		n2 := (i + 1) * 10

		s := fmt.Sprintf("타잔이 %d 원짜리 팬티를 입고, %d 원짜리 칼을 차고 노래를 한다.", n1, n2)
		fmt.Println(s)
	}
}

func main() {
	print(10)
	pibo(10)
}
