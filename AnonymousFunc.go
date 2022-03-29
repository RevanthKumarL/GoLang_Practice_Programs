// in order to create an inline function
//doesn't have to contain any name

//func(parameter_list)(return_type){
// code
//use return statement if return_type are given
//if return_type is not given, then don't	
//use return statement
//}

package main
import "fmt"

func main(){
	func(){
		//anonymous function
		fmt.Println("Welcome! to GeeksforGeeks")

	}()

	value:= func() {
		fmt.Println("Welcome to Riverside!")
	} 
	value()

	func(ele string){
		fmt.Println(ele)
	}("Freaks for Geeks")

	
}