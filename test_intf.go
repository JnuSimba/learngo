package main

import "fmt"
import "unsafe"
import "reflect"

/*
runtime.h
struct Iface
{
	Itab* tab;
	void* data;
};
struct Itab
{
	InterfaceType* inter;
	Type* type;
	void (*fun[])(void);
};
*/

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer //embed intf
	Print()
}

type User struct {
	id   int
	name string
}

// implement interface
func (self *User) String() string {
	return fmt.Sprintf("user: %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

// empty interface. all type implemented.
func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

type Tester struct {
	s interface { // anonymous interface as field
		String() string
	}
}

func main() {

	println("------------------------------")
	var t Printer = &User{1, "tom"} //(*User) methods conclude String() & Print()
	t.Print()

	println("------------------------------")
	Print(1)
	Print("Hello world!")

	println("------------------------------")
	tr := Tester{&User{2, "jack"}}
	fmt.Println(tr.s.String())

	println("------------------------------")
	u := User{3, "simba"}
	var vi, pi interface{} = u, &u
	// vi.(User).name = "SIMBA" // Error: cannot assign to vi.(User).name
	pi.(*User).name = "SIMBA"
	fmt.Printf("%v\n", vi.(User))  // vi is a readonly copy of u
	fmt.Printf("%v\n", pi.(*User)) // Printf will call u.String()

	println("------------------------------")
	if i, ok := pi.(fmt.Stringer); ok {
		fmt.Println(i)
		println("it's ok....")
	}

	switch v := pi.(type) {
	case nil: // 0 == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println(v)
	case func() string: //func
		fmt.Println(v())
	case *User: // *struct
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")

	}

	println("------------------------------")
	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab contain *int typeinfo, data = nil
	type iface struct {
		itab, data uintptr
	}

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}
