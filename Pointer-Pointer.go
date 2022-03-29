package main
import "fmt"

func main(){
	var v int = 101

	var pt1 *int = &v

	var pt2 **int = &pt1

	fmt.Println("The value of variable v:", v)

	*pt1 = 202

	fmt.Println("Value stored in v after changing pt1:",v)

	**pt2 = 303
	
	fmt.Println("Value stored in v after changing pt2:",v)




}