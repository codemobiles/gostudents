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
	fn1()
}

func fn1()  {
	fmt.Printf("Hey %s\n", "lek")
}
