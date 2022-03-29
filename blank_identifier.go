package main
import "fmt"

func main(){
	//calling the function
	//function returns two values which are
	//assigned to mul and div identifier

	//example: mul, div := mul_div(105, 7)

	mul, _ := mul_div(105, 7)

	fmt.Println("105 * 7 = ", mul)
}

//function returning two
//values of integer type
func mul_div(n1 int, n2 int) (int, int){
	//returning the values
	return n1*n2, n1/n2
}