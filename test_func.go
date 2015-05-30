package main

import "fmt"
import "os"
import "errors"

func test1(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

// anonymous func as param
func test2(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string // define function type

func test3(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

// var param is slice, can only have one
func test4(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}

// named return para
func test_named(x, y int) (z int) {
	// in same level, 'z redeclared in this block'

	{
		var z = x + y
		// return  Error: z is shadowed during return.
		return z // must be explicit not implicit
	}
}

func test_defer(x, y int) (z int) {
	defer func() {
		println(z)
	}() // must have param even empty

	z = x + y
	return z + 200 // (z = z + 200) --> (call defer) --> return
}

func test_closure() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func test_deferlifo() error {
	defer println("a")
	defer println("b")

	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// f.Close() will be call before return
	defer f.Close() // register call not register func.

	f.WriteString("Hello world!")
	return nil
}

func test_recover() {
	defer func() {
		if err := recover(); err != nil { // recover must be call in defer func
			println(err.(string)) // interface() --> string
		}
	}()

	panic("panic error~!")
}

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func test_error() {
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		fmt.Println(err)
	}
}

func main() {

	n1, s1 := test1(20, 30, "%d")

	s2 := test2(func() int { return 100 })

	s3 := test3(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(n1, s1, s2, s3)
	println(test4("sum: %d", 1, 2, 3))
	si := []int{4, 5, 6}
	println(test4("sum: %d", si...))

	println(test_defer(1, 2))

	// -- function variable ---
	fn := func() { println("Hello world!") }
	fn()

	// -- function collection
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	//-- function as field
	d := struct {
		fn func() string
	}{
		fn: func() string { return "hello World!" },
	}
	println(d.fn())

	// -- channel of function --
	fc := make(chan func() string, 2)
	fc <- func() string { return "hi!!" }
	println((<-fc)())

	fcl := test_closure()
	fcl()

	test_deferlifo()

	test_recover()

	test_error()
}
