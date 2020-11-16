package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	username string `json:"username"`
	password string `json:"password"`
	age      int    `json:"age"`
	isAdmin  bool   `json:"isAdmin"`
}

type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	u2 := account{}
	jsonString2 := `{"username":"codemobiles", "password":"777"}`
	if e := json.Unmarshal([]byte(jsonString2), &u2); e != nil {
		fmt.Println(e)
	}

	fmt.Println(u2)
}

func f1(value user) {
	fmt.Println("Values in Struct :", value)
}
