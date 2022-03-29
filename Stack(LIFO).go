package main
import "fmt"

func main(){
	//create
	var stack []string

	//push
	stack = append(stack, "Praise the lord")

	stack = append(stack, "And break the law")

	fmt.Println(stack)

	for len(stack) >0 {

		//print the top

		n:= len(stack)-1

		fmt.Println(stack[n])

		//pop
		stack = stack[:n]
	}
}