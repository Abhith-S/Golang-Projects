package main;

import "fmt"

func main() {

	sum, message := add(2,8);

	fmt.Println(message,sum);
}

func add(num1 int, num2 int) (int, string){

	return num1+num2,"Sum is";

}