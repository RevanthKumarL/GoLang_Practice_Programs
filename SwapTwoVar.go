package main
import "fmt"

func swap() []int{
	a,b := 15,10
	b,a = a,b
	return []int{a,b}
}
func main(){
	fmt.Println(swap())
	
}