package main

import "fmt"

func main(){
	// basic instance
	purchases := [5]float32{1,2,3,4,5}
	// indicies
	secondPurchase := purchases[1];
	// len
	fmt.Println(len(purchases));
	fmt.Println(purchases);
	fmt.Println(secondPurchase);

	// slices (arrays are immutable so we need slices to reduce or add something)
	slice := purchases[:] // copy of purch
	slice = append(slice, 12.32, 23,14);
	fmt.Println(slice);
}
