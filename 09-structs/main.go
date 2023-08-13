package main

import "fmt"

// declaration
type Abhith struct {
	Name     string
	Age      int
	isSingle bool
}

func main() {

	//usage
	abhith := Abhith{"Abhi", 23, false}

	fmt.Println(abhith)
	fmt.Println(abhith.Name)
	fmt.Printf("Struct is %+v", abhith)

}
