package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input program"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:",input)
	fmt.Printf("The type of this is %T ",input)
}
