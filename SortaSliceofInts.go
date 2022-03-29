package main
import (
	"fmt"
	"sort"
)
func main(){
	//syntax: func Ints(slc []int)
	//create and init the slices using shorthand declaration
	slc1:= []int{400,600,100,300,500,200,900}
	slc2:= []int{-23,567,-34,67,0,12,-5}
	fmt.Println("Slices(Before):")
	fmt.Println("Slice 1: ",slc1)
	fmt.Println("Slice 2: ",slc2)
	//sorting the slice of ints 
	//using the ints function
	sort.Ints(slc1)
	sort.Ints(slc2)
	fmt.Println("Slices(After):")
	fmt.Println("Slice 1:",slc1)
	fmt.Println("Slice 2:",slc2)
	res1:= sort.IntsAreSorted(slc1)
	res2:= sort.IntsAreSorted(slc2)
	fmt.Println("Results:")
	fmt.Println("Is Slice1 is sorted?: ",res1)
	fmt.Println("Is Slice2 is sorted?: ",res2)
}