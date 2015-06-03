package main

import "unsafe"
import "time"
import "runtime"
import "fmt"

//  GODEBUG="gctrace=1" ../../bin/test_gc
type data struct {
	x [1024 * 1024]byte
}

type data2 struct {
	x [1024 * 1024]byte
	y int
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

// the object hold by uintptr will be gc
func test_uintptr() {
	const N = 1000
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
}

func test1() unsafe.Pointer {
	p := &data{}
	return unsafe.Pointer(p)
}

func test2() unsafe.Pointer {
	d := data2{}
	return unsafe.Pointer(&d.y)
}

// the object hold by unsafe.Pointer will not be gc
func test_unsafe_pointer() {
	const N = 1000
	cache := new([N]unsafe.Pointer)

	for i := 0; i < N; i++ {
		//cache[i] = test1()
		cache[i] = test2()
		time.Sleep(time.Millisecond)
	}
}

type Data struct {
	d [1024 * 1024]byte
	o *Data
}

// cycle ref + runtime.SetFinalizer will be memory leak
// go build -gcflags "-N -l"
// GODEBUG="gctrace=1" ../../bin/test_gc
func test_setfinalizer() {
	var a, b Data
	a.o = &b
	b.o = &a

	runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}

func main() {
	test_uintptr()
	println("===========================================")
	test_unsafe_pointer()
	println("===========================================")
	for {
		test_setfinalizer()
		time.Sleep(time.Millisecond)
	}
}
