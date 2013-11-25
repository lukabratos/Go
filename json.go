package main

import (
	"fmt"
	"encoding/json"
	)

type person struct {
	Name string
	Surname string
	Age uint
}

func main() {
	p := person{}
	src := []byte(`{"name":"Jack","surname":"Jones","age":30}`)
	err := json.Unmarshal(src, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %v %v you are %v years old!\n", p.Name, p.Surname, p.Age)
}