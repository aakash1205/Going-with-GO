package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my pizza app")
	fmt.Println("Please rate our pizza between 1 and 5:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for reading ", input)
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // strings is a package
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("rating is updated as: ", numrating+1)
	}

}
