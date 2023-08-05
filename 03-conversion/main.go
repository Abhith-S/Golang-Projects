package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	

	reader := bufio.NewReader(os.Stdin);

	fmt.Println("Provide any number");

	input,_ := reader.ReadString('\n');

	fmt.Println("your input is ",input);

	inputInNum, err := strconv.ParseFloat(strings.TrimSpace(input),64);

	if(err != nil){
		fmt.Println(err);
	}else{

		//now add 1 to the input
	//here we take it as string , so we have to convert it
	//also that the input will have an empty space attched to it , so remove before converting
	fmt.Println("your input after adding 1 is ", inputInNum + 1);

	}

	

	


}