package main
import "fmt"

func main(){
	r:= [3]int{8,7,6}

	if len(r)==0{
		fmt.Println("Empty!")
	} else {
		fmt.Println("Not Empty!")
	}
}