//Basic type: Numbers, strings, and booleans
//Aggregate type: Array and structs
//Reference type: Pointers, slices, maps, functions, and channels

package main

import (
	"fmt"
	
)

func main() {
	// 8bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	//16bit signed int
	var Y int16 = 32767
	fmt.Println(Y-2, Y+2)

	a := 20.45
	b := 34.89

	c := b-a
	fmt.Printf("Result is: %f", c)
	fmt.Printf("\nThe type of c is: %T", c)

	var d complex128 = complex(6, 4)
	var e complex64 = complex(9, 2)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Printf("The type of d is %T and "+"the type of e is %T\n", d, e)

	//variables
	str1 := "brosbeforehoes"
	str2 := "babesbeforebros"
	str3 := "mebeforeanyone"
	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println( result1)
	fmt.Println( result2)

	//display the type of result1 and result2
	fmt.Printf("The type of result1 is %T and "+
	"the type of result2 is %T\n", result1, result2)

	//str variable which stores the strings

	str := "brosbeforehoes"
	//length of the string
	fmt.Printf("Length of the string is: %d", len(str))
	//display the string
	fmt.Printf("\nString is: %s", str)
	//display the type of str variable
	fmt.Printf("\nType of str is: %T", str)

}


