package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	d1 := "Lek"
	fmt.Println(d1)

	var d1Pointer *string
	d1Pointer = &d1
	d1Clone := d1

	*d1Pointer = "Kan"
	clear(&d1)

	fmt.Println("d1Pointer :", *d1Pointer)
	fmt.Println("d1 :", d1)
	fmt.Println("d1Clone :", d1Clone)
}


func clear(v1 *string)  {
	*v1 = ""
}
