package main

import (
	"fmt"
	foo "net/http" //Named Import (behaves like import pandas as pandu)
)

func main() {
	fmt.Println("Hello, Go Standard Library")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status:", resp.Status)
}

/*
-----Data Types----------

Integers
Floating Point Numbers
Complex Numbers
Booleans
Strings
Constants
Arrays
Structs
Pointers
Maps
Slices
Functions
Channels
JSON
Text & HTML Templates
*/