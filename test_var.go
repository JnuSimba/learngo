package main

import "fmt"
import "math"
import "unsafe"

var x1 int
var f1 float32 = 1.6
var s1, n1 = "abcd", math.MinInt16

var (
	a1 int
	b1 float32
)

// const
const (
	a2, b2 = "abc", len(a2)
	c1     = unsafe.Sizeof(b2)
)

//enum
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday   = 'a'  // 'a'
	Wednesday        // 'a'
	Thursday  = iota // 4
	Friday           // 5
	Saturday         // 6
)

/*
runtime.h
struct String
{
	byte* str;
	intgo len;
};
*/

func _uintptr() {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d)) // *struct --> Pointer --> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset
	p2 := unsafe.Pointer(p)          // uintptr --> Pointer
	px := (*int)(p2)                 // Pointer --> *int
	*px = 200

	fmt.Printf("%#v\n", d)
}

func ptr() {
	x := 0x12345678
	p := unsafe.Pointer(&x) // *int --> Pointer
	n := (*[4]byte)(p)      // Pointer --> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	println()
}

func main() {
	// in function body can use :=
	arr, i := [3]int{0, 1, 2}, 0
	i, arr[i] = 2, 100 // from left to right

	unused := 0
	_ = unused      // unless Error: unused declared and not used
	const x = "xxx" // unused const variable is ok.

	s := "abc" // double quote as string
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
	s = "hello" // re-assigment of s.
	{
		s := 1000 // definition. in another code block.
		fmt.Printf("in {} s = %d\n", s)
	}

	b := make([]int, 3) // make slice. malloc memory and initialise, return ref.
	b[1] = 10
	fmt.Println(b)
	c := new([]int) // malloc memory and memset to 0, return ptr.
	//not c[1] (index of type *[]int)
	fmt.Println(c)
	fmt.Printf("c=%p, *c=%v\n", c, *c)
	*c = make([]int, 10)
	fmt.Println(c)
	fmt.Println((*c)[2])

	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("p=%p, p.a=%v\n", p, p.a) // not p->a

	ptr()
	_uintptr()
}
