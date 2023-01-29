package main

import "fmt"

func main(){
	// pointers point to a memory spot in memory

	var c = 1;
	var d *int; // d is going to point to a variable
	d = &c // now d will point to c, so tey can be considered the same variable
	*d+=1; // this will change c

	fmt.Println(c); // c=2

}
