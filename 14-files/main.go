package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//write to and create file
	content := "This is to add to text file"
	file, err := os.Create("./sampleFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is ", length)

	defer file.Close()

	readFile("./sampleFile.txt")

}

// read file
func readFile(filename string) {

	//data is byte format
	data, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("data in text file is : ", string(data))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
