package main

import (
	"fmt"
	
)
func foo() (string, string, string){
	return "three", "values", "it is!"

}
func main(){
	fmt.Println(foo())
}