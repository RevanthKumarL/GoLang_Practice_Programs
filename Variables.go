// var variable_name type  = expression

package main
import "fmt"

func main() {
	//variable declared and initialized without explicit type

	var myvariable1 = 20
	var myvariable2 = "fuckbitches"
	var myvariable3 = 34.80

	//display values and types of variables
	fmt.Printf("The value of myvariable1 is: %d\n", myvariable1)
	fmt.Printf("The type of variable1 is: %T\n", myvariable1)
	fmt.Printf("The value of variable2 is: %s\n", myvariable2)
	fmt.Printf("The type of variable2 is: %T\n", myvariable2)
	fmt.Printf("The value of variable3 is: %f\n", myvariable3)
	fmt.Printf("The type of variable3 is: %T\n", myvariable3)

	var myvariable4 int
	var myvariable5 string
	var myvariable6 float64

	//display the zero value of the variables
	fmt.Printf("The value of myvariable4 is: %d\n", myvariable4)
	fmt.Printf("The value of myvariable5 is: %s\n", myvariable5)
	fmt.Printf("The value of myvariable6 is: %f\n", myvariable6)

	//display mutiple variables of the same type in a single line
	var myvariable7, myvariable8, myvariable9 int = 2, 454, 67
	fmt.Printf("The value of myvariable7 is : %d\n", myvariable7)
	fmt.Printf("The value of myvariable8 is : %d\n", myvariable8)
	fmt.Printf("The value of myvariable9 is : %d\n", myvariable9)

	var myvar1, myvar2, myvar3 = 2, "GFG", 67.89
	fmt.Printf("\nThe value of myvar1 is : %d\n", myvar1)
	fmt.Printf("The type of myvar1 is : %T\n", myvar1)
	fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
	fmt.Printf("The type of myvar2 is : %T\n", myvar2)
	fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
	fmt.Printf("The value of myvar3 is : %T\n", myvar3)


	myv1 := 39
	myv2 := "GeeksforGeeks"
	myv3 := 34.67

	myv4, myv5, myv6 := 800, 34, 56

	fmt.Printf("\nThe value of myv1 : %d\n", myv1)
	fmt.Printf("The type of myv1 : %T\n", myv1)
	fmt.Printf("\nThe value of myv2 : %s\n", myv2)
	fmt.Printf("The type of myv2 : %T\n", myv2)
	fmt.Printf("\nThe value of myv3 : %f\n", myv3)
	fmt.Printf("The type of myv3 : %T\n", myv3)

	fmt.Printf("\nThe value of myv4 is : %d\n", myv4)
	fmt.Printf("The type of myv4 is : %T\n", myv4)
	fmt.Printf("\nThe value of myv5 is : %d\n", myv5)
	fmt.Printf("The type of myv5 is : %T\n", myv5)
	fmt.Printf("\nThe value of myv6 is : %d\n", myv6)
	fmt.Printf("The type of myv6 is : %T\n", myv6)



}