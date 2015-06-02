package main

//import "fmt"
import "math"
import "sync"
import "runtime"

func sum(id int) {
	var x int64
	var i uint32
	for i = 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	println(id, x)
}

func main() {
	wg := new(sync.WaitGroup)
	/*	wg.Add(2)

		for i := 0; i < 2; i++ {
			go func(id int) {
				defer wg.Done()
				sum(id)
			}(i)
		}
	*/

	/*
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer println("A.defer")

			func() {
				defer println("B.defer")
				runtime.Goexit() //stop execution
				println("B")     // no print
			}()

			println("A") // no print
		}()
	*/

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				runtime.Gosched()
			} // similar to 'yield'. suspend execution
		}
	}()

	go func() {
		defer wg.Done()
		println("hello world!~")
	}()

	wg.Wait()
}
