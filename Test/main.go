package main

import "fmt"

//
//import (
//	"fmt"
//	"reflect"
//)
//
//type 袁神 int
//type 原神 = int
//
//func main() {
//	var a 袁神
//	var b 原神
//	fmt.Printf("type of a:%T\n", a)
//	fmt.Printf("type of b:%T\n", b)
//	rfTypeOf(a)
//	rfTypeOf(b)
//	TypeOf(a)
//	TypeOf(b)
//}
//func rfTypeOf(data interface{}) {
//	of := reflect.TypeOf(data)
//	fmt.Println(of)
//}
//func TypeOf(data interface{}) {
//	switch data.(type) {
//	case 袁神:
//		fmt.Println("Type is int")
//	case nil:
//		fmt.Println("Type is nil")
//	default:
//		fmt.Println("Type Not Found")
//	}
//}

//package main
//
//import "fmt"
//
//type 袁神 struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	a := 袁神{
//		Name: "yuan1",
//		Age:  18,
//	}
//	var b 袁神
//	b.Name = "yuan2"
//	b.Age = 18
//	c := struct {
//		Name string
//		Age  int
//	}{
//		"yuan3",
//		18,
//	}
//	d := NewYuanShen("袁神 tql", 18)
//	fmt.Printf("%#v\n", a)
//	fmt.Printf("%#v\n", b)
//	fmt.Printf("%#v\n", c)
//	fmt.Printf("%#v\n", d)
//}
//func NewYuanShen(name string, age int) *袁神 {
//	return &袁神{
//		Name: name,
//		Age:  age,
//	}
//	a := struct {
//		Name string
//		Age  int
//	}{
//		"yuan1",
//		18,
//	}
//	fmt.Println(a)
//	fmt.Println(a.Name)
//	fmt.Println(a.Age)
//}

//type People struct {
//	Name  string
//	Age   int
//	Books []Book
//}
//type Book struct {
//	Name string
//}
//
//func (w People) PrintName() {
//	fmt.Println(w.Name)
//}
//func (w People) PrintAge() {
//	fmt.Println(w.Age)
//}
//func (w People) PrintBook() {
//	fmt.Println(w.Books)
//}
//func (b Book) PrintBookName() {
//	fmt.Println(b.Name)
//}

//func main() {
//	a := 10
//	b := &a
//	fmt.Printf("a:%d ptr:%p\n", a, &a)
//	fmt.Printf("b:%p type:%T\n", b, b)
//	fmt.Println(&b)
//}

//func main() {
//	var p *string
//	fmt.Println(p)
//	fmt.Printf("p的值是%v\n", p)
//	if p != nil {
//		fmt.Println("非空")
//	} else {
//		fmt.Println("空值")
//	}
//}

//func main() {
//	a := new(string)
//	*a = "军哥NB"
//	fmt.Println(*a)
//
//	b := make(map[string]string)
//	b["军哥"] = "YYDS"
//	fmt.Println(b)
//}

//type Cat struct{}
//
//func (c Cat) Say() string { return "喵喵喵" }
//
//type Dog struct{}
//
//func (d Dog) Say() string { return "汪汪汪" }
//func main() {
//	c := Cat{}
//	fmt.Println("猫:", c.Say())
//	d := Dog{}
//	fmt.Println("狗:", d.Say())
//}

func main() {
	var x interface{}
	s := "YuanShen"
	x = s
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
