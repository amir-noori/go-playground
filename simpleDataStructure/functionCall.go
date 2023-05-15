package simpleDataStructure

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func CallPersonFactories() {
	p1 := createPersonV1()
	p2 := createPersonV2()

	fmt.Println("Person 1", p1.FirstName, p1.LastName)
	fmt.Println("Person 2", p2.FirstName, p2.LastName)
}

func createPersonV1() Person {
	p := Person{
		FirstName: "Jack",
		LastName:  "Black",
	}
	return p
}

func createPersonV2() *Person {
	p := Person{
		FirstName: "Joe",
		LastName:  "Brown",
	}
	return &p
}

const size = 2000

func CopyStackData() {
	s := "Hello"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	fmt.Println(c, s, *s)
	c++
	if c == 10 {
		return
	}
	stackCopy(s, c, a)
}
