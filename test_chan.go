package main

import "fmt"
import "os"

import "runtime"
import "math/rand"
import "time"

func test_syn() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data { // receive from data until close
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		exit <- true // send exit
	}()

	data <- 1 // send data
	data <- 2
	data <- 3
	close(data) // close chan

	fmt.Println("send over.")
	<-exit //wait for exit.
}

func test_asyn() {
	data := make(chan int, 3) // buffer of 3 elements
	exit := make(chan bool)

	data <- 1 //will not block
	data <- 2
	data <- 3
	go func() {
		// if d, ok := <-data; ok
		for d := range data { // will not block until empty
			fmt.Println(d)
		} // exit until close
		exit <- true
	}()

	data <- 4 // blocked if buffer is full
	data <- 5
	close(data)
	<-exit
}

func test_feature() {
	/*
		var a, b chan int = make(chan int), make(chan int, 3) // buffer size not count
		a <- 1
		fmt.Println(len(a), cap(a)) // 0 0
		fmt.Println(len(b), cap(b)) // 1 3
	*/
	/*
		c := make(chan int, 3)
		var send chan<- int = c // send-ony
		var recv <-chan int = c // receive-only
		send <- 1
		// <-send // Error: receive from send-only type chan<- int
		<-recv
		// recv <- 2 // Error: send to receive-only type <-chan int
	*/
}

func test_select() {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""

		for {
			select { //random select available chan
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			case <-time.After(time.Second * 3):
				fmt.Println("timeout.")
			}

			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case a <- i:
		case b <- i:
		}
	}

	close(a)
	close(b)
}

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()

	return c
}

// chan as struct field
type Request struct {
	data []int
	ret  chan int
}

func NewRequest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}

func main() {
	runtime.GOMAXPROCS(2)

	test_syn()
	println("=========================================")
	test_asyn()
	println("=========================================")
	test_select()
	println("=========================================")
	//	t := NewTest()
	//	println(<-t)
	println("=========================================")
	req := NewRequest(10, 20, 30)
	Process(req)
	fmt.Println(<-req.ret)
}
