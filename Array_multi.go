package main
import (
	"fmt"
)
func main(){
	//creating and init 2-dimensional array
	arr:= [3][3]string{{"C#", "C", "Python"},{"Java", "Scala", "Perl"},{"C++", "Go", "HTML"},}
	fmt.Println("Elements of Array 1")
	for x:= 0; x<3; x++{
		for y:= 0; y<3; y++{
			fmt.Println(arr[x][y])
		}
	}
	fmt.Println(arr[0][1]) //to find the specific element 
	// out of [3][3] multi dimensional array
	var arr1[2][2] int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	fmt.Println("Elements of Array 2")
	for p:= 0; p<2; p++{
		for q:= 0; q<2; q++{
			fmt.Println(arr1[p][q])
		}
	}
	myarr:= [2]int{2, 3}
	fmt.Println("Elements of the Array: ", myarr)

	arr2:= [3]int{9, 7, 6}
	arr3:= []int{9, 7, 8, 6, 1, 2, 3, 4, 5}
	arr4:= [3]int{8, 4, 5}
	fmt.Println("Length of the array 1 is:", len(arr2))
	fmt.Println("Length of the array 2 is:", len(arr3))
	fmt.Println("Length of the array 3 is:", len(arr4))

	myarray:= []int{29,48,76,39,
	30,49,32,80}
	for r:=0; r < len(myarray); r++{
		fmt.Printf("%d\n", myarray[r])
	} 
}