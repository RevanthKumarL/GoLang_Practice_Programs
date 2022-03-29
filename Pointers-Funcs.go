package main
import "fmt"

//take a function with a pointer
//type pointer as a parameter
func ptf(a *int) {
	//dereferencing
	*a = 748
}
//main function
func main() {
	//take a normal variable
	var x = 100
	fmt.Println("The value of x before the function call:",x)

	//take a pointer variable
	//assign the address of x to it

	var ap *int = &x

	//calling the function by passing the pointer to func
	ptf(ap)

	ptf(&x)
	//both are same
	fmt.Println("The value of x after the func call:",x)
}