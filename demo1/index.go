package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var t1 string = "test"
	var t2 int = 10
	var t3 bool = true
	t4 := "lek"
	t1 = "codemobiles"
	const t5 string = "xxx"
	// t5 = "yyy"

	fmt.Println(t1, t2, t3, t4, t5)
	fn1("lek")
	fn2("hello", 11, false)
}

func fn1(p1 string) {
	fmt.Printf("Hey %s\n", p1)
}

func fn2(p1 string, p2 int, p3 bool) {
	fmt.Printf("Hey %s, %d, %v\n", p1, p2, p3)
}

func fn3(p1 string, p2 int, p3 bool) string {
	fmt.Printf("Hey %s, %d, %v\n", p1, p2, p3)
	return "I am return function"
}
