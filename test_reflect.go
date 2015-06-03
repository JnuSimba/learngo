package main

import "fmt"
import "reflect"
import "unsafe"

type User struct {
	Username string `field:"username" type:"varchar(20)"`
	Age      int    `field:"age" type:"tinyint"`
}

type Admin struct {
	User
	title string
}

func test_type() {
	//	var u Admin
	u := new(Admin)
	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

func test_access() {
	var u Admin
	t := reflect.TypeOf(u)

	f, _ := t.FieldByName("title")
	fmt.Println(f.Name)

	f, _ = t.FieldByName("User") // embed field
	fmt.Println(f.Name)

	f, _ = t.FieldByName("Username") // member of embed field
	fmt.Println(f.Name)

	f = t.FieldByIndex([]int{0, 1}) // Admin[0] -> User[1] -> age
	fmt.Println(f.Name)
}

func (*User) ToString() {}
func (Admin) test()     {}

func test_method() {
	var u Admin

	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("----value interface ---")
	methods(reflect.TypeOf(u))

	fmt.Println("----pointer interface ---")
	methods(reflect.TypeOf(&u))
}

func test_metadata() {
	var u User
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Username")
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func get_type() {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)

	e := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(e)
}

type Data struct {
	b byte
	x int32
}

func (*Data) String() string {
	return ""
}

func test_impl() {
	var d *Data
	t := reflect.TypeOf(d)

	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(it))
}

func test_align() {
	var d Data
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())

	f, _ := t.FieldByName("b")
	fmt.Println(f.Type.FieldAlign())
}

func test_value() {
	u := &Admin{User{"Kacl", 23}, "NT"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String())
	fmt.Println(v.FieldByName("Age").Int())
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())

	arr := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, arr.Len(); i < n; i++ {
		fmt.Println(arr.Index(i).Int())
	}

	ma := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range ma.MapKeys() {
		fmt.Println(k.String(), ma.MapIndex(k).Int())
	}
}

func test_isvalid() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	f := v.FieldByName("a")
	fmt.Println(f.Kind(), f.IsValid())
	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
	println("---------------------------------")

	var pi *int
	var x interface{} = pi
	fmt.Println(x == nil)
	pp := reflect.ValueOf(pi)
	fmt.Println(pp.Kind(), pp.IsNil())
	println("---------------------------------")

	pe := reflect.ValueOf(&u).Elem()
	pe.FieldByName("Username").SetString("Tome")
	fa := pe.FieldByName("Age")
	fmt.Println(fa.CanSet())

	if fa.CanAddr() {
		age := (*int)(unsafe.Pointer(fa.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(fa.Addr().Pointer())
		*age = 88
	}

	fmt.Println(u, pe.Interface().(User))
	println("---------------------------------")

	arr := make([]int, 0, 10)
	av := reflect.ValueOf(&arr).Elem()
	av.SetLen(2)
	av.Index(0).SetInt(100)
	av.Index(1).SetInt(200)

	fmt.Println(av.Interface(), arr)

	av2 := reflect.Append(av, reflect.ValueOf(300))
	av2 = reflect.AppendSlice(av2, reflect.ValueOf([]int{400, 500}))
	fmt.Println(av2.Interface())
	m := map[string]int{"a": 1}
	mv := reflect.ValueOf(&m).Elem()
	mv.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) //update
	mv.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(mv.Interface(), m)
}

func main() {
	/*
					test_access()
					println("==========================")
					test_type()
					println("==========================")
					test_method()
					println("==========================")
					test_metadata()
					println("==========================")
					get_type()

				println("==========================")
				test_impl()

			println("==========================")
			test_align()

		println("==========================")
		test_value()
	*/

	println("==========================")
	test_isvalid()
}
