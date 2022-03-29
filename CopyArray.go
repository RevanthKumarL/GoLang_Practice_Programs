package main
import (
	"fmt"
)
func main(){
	myarr1:= [5]string{"Scala", "Perl", "Java", "Python", "Go"}
	myarr2:= myarr1
	fmt.Println("Array1: ", myarr1)
	fmt.Println("Array2: ", myarr2)
	myarr1[0] = "C++"

	fmt.Println("\nArray1: ", myarr1)
	fmt.Println("Array2: ", myarr2)

	myarr3:= [6]int{12, 456, 67, 65, 34, 43}
	//copying the array into new var and elements passed by ref*
	myarr4:= &myarr3

	fmt.Println("\nArray3: ", myarr3)
	fmt.Println("Array4: ", *myarr4)
	myarr3[5] = 100000

	fmt.Println("\nArray3: ", myarr3)
	fmt.Println("Array4: ", *myarr4)

}