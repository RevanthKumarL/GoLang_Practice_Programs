//function
//defer func_name(parameter_list type)return_type{
//code
//}
//method
//defer func(receiver type) method_name(parameter_list){
	//code
//}
//defer func (parameter_list)(return_type){
	//code
	//}()

package main
import "fmt"
//functions

func mul(a1, a2 int) int{
	res:= a1 * a2
	fmt.Println("Result1: ", res)
	return 0
}

func add(a3, a4 int)int{
	result:= a3 + a4
	fmt.Println("Result2: ", result)
	return 0 
}

func show() {
	fmt.Println("Hello, GeeksforGeeks")
}
// first the func that is directly called 
// then the func show() and its print statement
// then the bottom_to_up of defer funcs declared
// after all of it, show() is written; end of the code
func main(){

	mul(23, 34)

	defer mul(23, 56)
	defer add(34, 56)
	defer add(56, 79)

	show()
}