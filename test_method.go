package main

import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 10)}
}

// no receiver param name
func (*Queue) Push(e interface{}) error {
	panic("not implement")
}

// receive type : T or *T
func (self *Queue) length() int {
	return len(self.elements)
}

//-------------------------------------------
type Data struct {
	x int
}

//method is a special function
func (self Data) ValueTest() { // func ValueTest(self Data)
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data)
	fmt.Printf("Pointer: %p\n", self)
}

//-------------------------------------------------

// method value vs method expression
// instance.method(args...) ---> <type>.func(instance, args...)
// instance's type may T or *T

type User struct {
	id   int
	name string
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

//------------------------------
type Empty struct{}

func (Empty) TestValue()    {}
func (*Empty) TestPointer() {}

//-----------------------------

func main() {
	println("----------------------------")
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)

	//-----------------
	println("-----------------------------")
	u := User{1, "Tom"}
	myUser := u.TestValue // method value will copy receiver
	u.name, u.id = "kack", 2
	fmt.Printf("User: %p, %v\n", &u, u)
	myUser() // implicit deliver receiver

	// type T methods conclude 'receiver T method'
	// type (*T) methods conclude 'receiver T + *T method'
	mv := User.TestValue // method expression
	mv(u)                // explicit deliver receiver

	mp := (*User).TestPointer
	mp(&u)

	mp2 := (*User).TestValue
	mp2(&u)

	//-------------------------------
	println("-----------------------------------")

	var pp *Empty = nil
	pp.TestPointer()
	(*Empty)(nil).TestPointer() // method value
	(*Empty).TestPointer(nil)   // method expression

	// pp.TestValue() // Error: invalid memory address or nil pointer dereference
	// (Data)(nil).TestValue() // Error: cannot convert nil to type Data
	// Data.TestValue(nil) // Error: cannot use nil as type Data in function argument
}
