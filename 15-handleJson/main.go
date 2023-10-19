package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json video")
	EncodeJson();

}

func EncodeJson() {
	Courses := []course{
		{"React",
			20,
			"Edx",
			"1234",
			[]string{"js", "web"}},
		{"ML",
			20,
			"Edx",
			"1234",
			[]string{"js", "ML"}},
		{"AI",
			20,
			"Edx",
			"1234",
			nil},
	}

	//package this data as json data

	finalJson,err := json.MarshalIndent(Courses,"","\t");
	if err!=nil{
		panic(err);
	}
	fmt.Printf("%s \n",finalJson);

}
