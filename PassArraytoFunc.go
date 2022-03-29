package main
import (
	"fmt"
)
//func function_name(variable_name [size]type) { //cpde }

func myfun(a [6]int, size int) int{
	var k, val, r int

	for k=0; k<size; k++{
		val += a[k]
	}
	r=val/size
	return r
}
//main function
func main(){
	var arr = [6]int{67, 59, 29, 65, 4, 98}
	var res int
	//passing an array as an argument
	res = myfun(arr, 6)
	fmt.Printf("Final result is: %d", res)
}