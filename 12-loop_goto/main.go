package main;

import (
	"fmt"

)

func main() {
	//fmt.Println("Loops");

	fruits  := []string{"Mango","apple","peach","berry","palm"};

	//type 1
	// for i := 0; i< len(fruits); i++ {
	// 	fmt.Println(fruits[i]);
	// }

	//type 2
	// for i := range fruits{
	// 	fmt.Println(fruits[i]);
	// }

	//type 3
	// for i,fruit := range fruits{
	// 	fmt.Println(i);
	// 	fmt.Println(fruit);
	// }

	//goto
	for i := range fruits{

		if i == 3{
			goto myTest;
		}
		fmt.Println(fruits[i]);
	}

	myTest : fmt.Println("Go to worked");
}