package main;

import ("fmt")

func main()  {

	//fmt.Println("Pointer");
	num := 20;
	//create pointer
	ptr := &num;

	fmt.Println(ptr);
	fmt.Println(*ptr);
		
}