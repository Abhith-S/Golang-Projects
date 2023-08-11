package main

import "fmt"

func main(){

	//declare array
	var arr [4]int;

	arr[1] = 5;

	fmt.Println(arr);

	//declarea and initialize
	var arrNew = [6]string{"hia","hoi"};
	fmt.Println(arrNew);

	//array length
	fmt.Println("length of arr is ",len(arr));
}