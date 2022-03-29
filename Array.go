package main
import (
	"fmt"
)
//var array_name[length]type{item1, item2, ....itemN}
//var array_name[index]= element
func main(){
	var myarr[3]string
	myarr[0] = "Don't need a GF"
	myarr[1] = "Need a Job"
	myarr[2] = "And a car"

	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", myarr[0])
	fmt.Println("Element 2: ", myarr[1])
	fmt.Println("Element 3: ", myarr[2])
//array_name:= [length]type{item1, item2, ...itemN}
array:= [5]string{"Meta", "Google", "Microsoft", "Snapchat", "Netflix"}

fmt.Println("Elements of the Array:")
// for i:=1; i<4; i++{ print statement }
for i:=0; i<5; i++{
	fmt.Println(array[i])

}
}