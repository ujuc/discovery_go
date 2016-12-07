package strings

import "fmt"

func Example_printBytes() {
	s := "가나다"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}

	fmt.Println()

	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")

	fmt.Println("pre", b)

	b[2]++

	fmt.Println("post", b)

	fmt.Println(string(b))
	// Output:
	// pre [234 176 128 235 130 152 235 139 164]
	// post [234 176 129 235 130 152 235 139 164]
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"

	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}
