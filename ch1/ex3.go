package main

import "fmt"

func pibo(n int) {
	p, q := 0, 1
	fmt.Println(p, q)

	for i := 1; i <= n; i++ {
		p, q = q, p+q
		fmt.Println(p, q)
	}
}
