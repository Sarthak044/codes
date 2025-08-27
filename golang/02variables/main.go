package main

import "fmt"

// First Character is Captial L becuase this is now public Variable

const LoginToken string = "asdasdasdasd"

func main() {
	var username string = "FotoBomb"
	// 0 to 255	
	var smallVal uint8 = 255
	// 0 to 65535
	var medVal uint16 = 65535
	// 0 to 4294967295
	var largeVal uint32 = 4294967295
	// 0 to 18446744073709551615
	var extraLargeVal uint64 = 18446744073709551615
	// -128 to 127
	var smallInt int8 = 127
	// -32768 to 32767
	var medInt int16 = 32767
	// -2147483648 to 2147483647
	var largeInt int32 = 2147483647
	// -9223372036854775808 to 9223372036854775807
	var extraLargeInt int64 = 9223372036854775807	
	//Set of all IEEE-754 32 bit floating point numbers
	var floatVal float32 = 255.576889789879869689
	//Set of all IEEE-754 64 bit floating point numbers
	var doubleVal float64 = 255.57697589964432632638
	// Set of all complex numbers with float32 real and imaginary parts
	var complexVal complex64 = 1 + 2i
	// Set of all complex numbers with float64 real and imaginary parts
	var complexDoubleVal complex128 = 1 + 2i
	// byte alias for uint8
	var asciiVal byte = 'A'
	// rune is an alias for int32
	var runeVal rune = 'A'
	// Boolean
	var isLoggedin bool = true

	//Implicit Type

	var website = "hello world"
	fmt.Println(website)

	// No var Type

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	// Public Variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
	fmt.Println(isLoggedin)
	fmt.Printf("These are the values: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v \n", smallVal, medVal, largeVal, extraLargeVal, smallInt, medInt, largeInt, extraLargeInt, floatVal, doubleVal, complexVal, complexDoubleVal, asciiVal, runeVal)
	fmt.Printf("These are the Types: %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T, %T \n", smallVal, medVal, largeVal, extraLargeVal, smallInt, medInt, largeInt, extraLargeInt, floatVal, doubleVal, complexVal, complexDoubleVal, asciiVal, runeVal)
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	fmt.Printf("Variabe is of type: %T \n", isLoggedin)
}
