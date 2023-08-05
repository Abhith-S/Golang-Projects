package main

import (
	"fmt"
	"time"
)

func main() {
	
	//time now
	presentTime := time.Now();
	fmt.Println(presentTime)

	//give it a format , can;t change format value
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) 

	//create a time and date
	createdTime := time.Date(2000, time.February,10,23,45,0,0,time.UTC);

	fmt.Println(createdTime.Format("02-01-2006 15:04:05 Monday"));
}