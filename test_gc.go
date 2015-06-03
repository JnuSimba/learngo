package main

import "unsafe"
import "time"

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

func test_unsafe_pointer1() {
	const N = 1000
	cache := new([N]unsafe.Pointer)

	for i := 0; i < N; i++ {
		//cache[i] = test1()
		cache[i] = test2()
		time.Sleep(time.Millisecond)
	}
}

func main() {
	test_uintptr()
	println("===========================================")
	test_unsafe_pointer1()
}
