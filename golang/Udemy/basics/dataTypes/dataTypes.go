package main

import "fmt"

var globalVar string = "Bro"

func main(){

	var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"

	fmt.Println("",age,name,name1,count,lastName)
	/* 
	Default Values
	Numeric Types: 0
	Boolean Types: False
	String Type: ""
	Pointers, slices, maps, functions, and structs: nil

	:= can only be used in a function
	*/
	printName()
}

func printName(){
	firstName := "sarthak"
	fmt.Println(firstName)
	fmt.Println(globalVar)
}
