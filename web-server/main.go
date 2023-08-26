package main

import (
	"fmt"
	"log"
	"net/http"
)

// w is response data
// r is request
func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Println(w, "Hello in terminal")
	fmt.Fprintf(w,"Hello in browser page");

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	//parse the request
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() error : %v",err);
		return
	}

	fmt.Println("POST request succesfull")

	//now get the values
	name := r.FormValue("name");
	age := r.FormValue("age");

	fmt.Fprintf(w, "Name = %s \n",name);
	fmt.Fprintf(w, "Age = %s",age);

}

func main() {

	//handle root path
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// /hello path
	http.HandleFunc("/hello", helloHandler)

	// /form path
	http.HandleFunc("/form", formHandler)

	//create server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
