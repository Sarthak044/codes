package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the class"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")
	
	// comma ok syntax or error ok syntax
	input, _ := reader.ReadString('\n')
	// fmt.Println(input)
	fmt.Printf("Thanks for rating, %s", input)
}