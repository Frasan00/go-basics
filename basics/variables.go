package main

import "fmt"

func testingHello() {
	greet := "hello world"
	fmt.Println(greet)
	a := 4
	b := 5
	fmt.Println(a + b)
	array := [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println(array[i])
	}
}

func main() {
	testingHello()
}

// go run to run file (only main package can be executed)
// only main function is executed
