package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my time")
	curtime := time.Now() // method to get the current time almost same as Javascript
	fmt.Println("current time is ",curtime)
	fmt.Println("Current time is: ",curtime.Format("01-02-2006 15:04:05 Monday")) // standard mentioned in GO doc so follow this dates

}
