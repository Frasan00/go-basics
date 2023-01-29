package main

import "fmt"

func main(){
	// classic for
	for i:=0; i<10; i++{
		fmt.Println("first loop");
	}

	// while (there isn't a while in go, but you can use for in this way)
	j := 0;
	for j<5{
		fmt.Println("second loop");
		j++;
	}

}