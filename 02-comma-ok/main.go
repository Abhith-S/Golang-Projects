package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	//take input
	fmt.Println("Enter any number");
	reader := bufio.NewReader(os.Stdin);


	rating,_ := reader.ReadString('\n');

	fmt.Println("rating is ",rating);


}