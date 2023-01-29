package main

import "fmt"

/* a struct is the most similar way to a class in go*/
// struct are like interfaces in ts
type Animal struct{
	// a generic animal
	class string
	age int
	gender bool
}

func main(){

	var bear = Animal{
		class: "bear",
		age: 12,
		gender: true,
	}

	fmt.Println(bear.class, bear.age, bear.gender);
}
