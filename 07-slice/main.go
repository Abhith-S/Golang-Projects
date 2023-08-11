package main

import "fmt"

func main() {

	//declare slice
	var newSlice = []string{"Apple","Mango","Pine"};

	fmt.Println(newSlice);

	//add to slice
	newSlice = append(newSlice, "Orange");
	fmt.Println(newSlice);

	//remove mango
	newSlice = append(newSlice[:1],newSlice[2:]...)
	fmt.Println(newSlice);
}