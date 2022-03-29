package main
import "fmt"
//type definition
type data int
//method with non-struct type receiver
func (d1 data) multiply(d2 data) data{
	return d1*d2
}
/* return this code with main function */
//main function
func main() {
	value1:= data(23)
	value2:= data(34)
	res:= value1.multiply(value2)
	fmt.Println("Final result: ", res)
}