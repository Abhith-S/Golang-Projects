package main;

import "fmt"

func main() {

	//declare maps
	languages := make(map[string]string)

	//add values
	languages["JS"] = "Javascript";
	languages["PY"] = "Python";


	fmt.Println(languages);

	delete(languages,"JS");
	fmt.Println(languages);

}