package main

import "fmt"

func main(){
	// we use make to make a map 
	cart := make(map[string]int);// now we have a map

	// map population
	cart["gigi"] = 1;
	cart["lando"] = 13;

	fmt.Println(cart);
	fmt.Println(cart["gigi"]);

}
