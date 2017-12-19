package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human  实现的SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human  实现的Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la ... ", lyrics)
}

//Employee  重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

//因为这三个类型都实现了这两个方法。
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"sam", 36, "444-222-XXX"}, "Golang Inc", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Thing Ltd", 5000}
	// 定义Men类型的变量
	var i Men
	i = mike
	fmt.Println("This is Mike, a Student: ")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happpens")
	x := make([]Men, 3)
	//T 这三个都是不同类型的元素， 但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
}
