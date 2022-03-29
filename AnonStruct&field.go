/*variable_name:= struct {
	//fields
} {// fields _values}*/

package main
import (
	"fmt"
)

type Student struct{
	int
	string
	float64
}
func main(){
	//create and initialize
	//anon structure
	Element:= struct {
		name string
		branch string
		language string
		particles int
	}{
		name: "Pikachu",
		branch: "Cybersecurity",
		language: "GoLang",
		particles: 498,
	}
	fmt.Println(Element)
	value:= Student{123, "Bud", 89898.003}
	fmt.Println("Enrollment number: ", value.int)
	fmt.Println("Student name: ", value.string)
	fmt.Println("Package price: ", value.float64)
}