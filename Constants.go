package main

import "fmt"


func main() {

	const PI = 3.14

	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)

	fmt.Println("Happy", PI, "Day")

	const Correct= true
	fmt.Println("Go rules?", Correct)

	const A = "GFG"
	const B = "GeeksforGeeks"

	var helloWorld = A+ " " + B
	helloWorld += "!"
	fmt.Println(helloWorld)
	
	fmt.Println(A == "GFG")
	fmt.Println(B < A)

	const trueConst = true
	type myBool bool
	var defaultbool = trueConst
	var customBool myBool = trueConst

	fmt.Println(defaultbool)
	fmt.Println(customBool)

}