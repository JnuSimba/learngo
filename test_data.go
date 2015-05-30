package main

import "fmt"

// pointer array [n]*T

/*
runtime.h
struct Slice
{ // must not move anything
	byte* array; // actual data
	uintgo len; // number of elements
	uintgo cap; // allocated number of elements
};
*/

// array as param. pass by value
// use slice or array pointer *[n]T
func test_arr(x1 [2]int) {
	fmt.Printf("x1: %p\n", &x1)
	x1[1] = 1000
}

func test_slice() {
	s := make([]int, 0, 1) //make slice (type, len, cap)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i) // return new slice
		if n := cap(s); n > c {
			fmt.Printf("cap: %d --> %d\n", c, n)
			c = n
		}
	}
}

func test_map() {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	println(m[1].name)
	// m[1].name = "tom" error:cannot assign to m[1].name

	// if key is exist
	if v, ok := m[3]; ok {
		fmt.Println(v)
	}

	u := m[1]
	u.name = "tom"
	m[1] = u // replace value

	for k, v := range m {
		fmt.Println(k, v)
	}

	type user struct{ name string }
	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "jack" // replace value via pointer

}

func test_struct() {
	type Node struct {
		_    int
		id   int
		data *byte
		next *Node
	}

	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	fmt.Println(n2)

	type User struct {
		name string
		age  int
	}
	u1 := User{"Tom", 20} // sequence initialize. must have all field.
	fmt.Println(u1)

	type File struct {
		name string
		size int
		attr struct {
			perm  int
			owner int
		}
	}

	f := File{
		name: "test.txt",
		size: 1025,
		//attr: {0755, 1}, // Error: missing type in composite literal
	}

	f.attr.owner = 1
	f.attr.perm = 0755

	var attr = struct {
		perm  int
		owner int
	}{2, 0755}
	f.attr = attr

	fmt.Println(f)

	//--------------
	type Resource struct {
		id   int
		name string
	}

	type Classify struct {
		id int
	}

	type Manager struct {
		Resource // anonymous field. field name the same as type
		Classify
		name string
	}
	m := Manager{
		Resource{1, "people"},
		Classify{100},
		"jack",
	}

	println(m.name)
	println(m.Resource.name)
	println(m.Classify.id)

}

func main() {
	a1 := [2]int{}
	fmt.Printf("a: %p\n", &a1)
	test_arr(a1)
	fmt.Println(a1)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	sl := arr[1:4:5] //[low, high, max]
	sl[0] += 100     // change arr value
	sl[1] += 200
	p3 := &arr[3] // *int
	*p3 += 300
	fmt.Println(sl)
	fmt.Println(arr)

	test_slice()
	test_map()
	test_struct()
}
