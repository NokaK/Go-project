package main

import (
	"fmt"
	"math/rand"
)


type Items struct {
	name string
	age int
}

var items = []Items{
	{name: "Noka", age: 30},
	{name: "Baco", age: 28},
	{name: "Gio", age: 25},
	{name: "Sandro", age: 30},
	{name: "Ana", age: 17},
}
func main() {
	var random = rand.Intn(len(items))
    fmt.Printf("Random item: %v\n", items[random])
}
