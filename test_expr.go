package main

import "fmt"

func test_range() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}

		a[i] = v + 10 // index, value copy from origin, so v doesn't change
	}
	fmt.Println(a)
}

func test_ref() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			ss := s[:3] // ss is a ref
			ss[2] = 100
		}

		println(i, v)
	}
}

// no break, but have fallthrough
func test_switch() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]: // not const is ok
		println("a")
	case 1, 3:
		println("b")
	default:
		println("c")
	}
}

func test_bkgtcont() {
L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}

			print(x, ":", y, " ")
		}

		println()
	}
}

func main() {
	x := 0
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[0])
	}

	s := "abc"
	// '++ --' is statement but not expression, so need ';' before them in oneline
	for i, nu := 0, len(s); i < nu; i++ {
		println(s[i])
	}

	// index
	for i := range s {
		println(s[i])
	}

	// index, value
	for _, c := range s {
		println(c)
	}

	le := len(s)
	for le > 0 {
		println(s[le-1])
		le--
	}

	// C-like for (; ;)
	//for {
	//	println(s)
	//}

	// key, value
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v)
	}

	test_range()
	test_ref()
	test_switch()
	test_bkgtcont()
}
