package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	username string
	password string
	age      int
	isAdmin  bool
}

func (u user) clear() user {
	u.username = ""
	u.password = ""
	u.age = 0
	u.isAdmin = false
	return u
}

func main() {
	fmt.Println("Struct")
	var u1 user = user{username: "lek", password: "666", age: 30, isAdmin: false}
	fmt.Println(u1)

	u1.username = "admin"
	u1.password = "1234"
	u1.age = 20
	u1.isAdmin = true
	// fmt.Println(u1)
	f1(u1)

	u1 = u1.clear()
	fmt.Println("Clear", u1)

	jsonString := `{"username":"codemobiles", "password":"777", "age":10, "isAdmin":true}`
	if e := json.Unmarshal([]byte(jsonString), &u1); e != nil {
		fmt.Println(e)
	}

	fmt.Println(u1)
}

func f1(value user) {
	fmt.Println("Values in Struct :", value)
}
