package main

import "fmt"

func testingHello() {
	const book = "Harry Potter";
	var num =  "2";
	labarda := "lalalala";

	fmt.Println(book, num, labarda);
}

func main() {
	testingHello()
}

// go run to run file (only main package can be executed)
// only main function is executed
