package main
import (
	"fmt"
	"sort"
)
func main(){
	//first index in a slice is always 0 // the last one is len(slice)-1
	arr:= [7]string{"If", "you've", "got", "one-shot,", "one opportunity", "Grab-it"}
	fmt.Println("Array: ",arr)
	fmt.Println(len(arr))
	myslice:= arr[1:6]
	fmt.Println("Slice:", myslice)
	fmt.Println("Length of the slice:",len(myslice))
	fmt.Println("Capacity of the slice:",cap(myslice))
	var my_slice1 =  []string{"Snap", "Back", "to reality"}
	fmt.Println("My slice1:",my_slice1)
	my_slice2:= [7]int{12, 132, 34, 54, 98, 76, 55}
	var val int
	var r int
	for i:=0; i<7; i++{
		val += my_slice2[i]
	}
	r = val/len(my_slice2)
	fmt.Println("Average of slice2: ", r)
	arr1 := [5]string{"Let's", "Get", "This", "Job", "Done"}
	my_slice3:= arr1[1:2]
	my_slice4:= arr1[3:5]
	my_slice5:= arr1[:4]
	fmt.Println(arr1)
	fmt.Println(my_slice3)
	fmt.Println(my_slice4)
	fmt.Println(my_slice5)
	fmt.Println(cap(arr1))
	fmt.Println(cap(my_slice4))
	fmt.Println(len(my_slice4))
	//creating a slice with capacity and length
	var myslice1 = make([]int, 4, 7)
	fmt.Printf("Slice 1= %v, \nlength = %d, \nCapacity = %d\n", 
myslice1, len(myslice1), cap(myslice1))
	var myslice2 = make([]int, 7)
	//to print just the length and the capacity of the array
	fmt.Printf("Slice 1= %v, \nlength = %d, \nCapacity = %d\n", 
myslice2, len(myslice2), cap(myslice2))
for e:= 0;e<len(arr1); e++{
	fmt.Println(arr1[e])
	slice:= []string{"This", "is the", "beginning", "of", "a new era", "of", "technology"}
	for _, ele:= range slice{
		fmt.Printf("Element = %s\n", ele)
	}
}
arrr:= []int{23,98,76,34,45,21}
slc:= arrr[0:4]
fmt.Println(arrr)
fmt.Println(slc)
//for modifying the array and slices
slc[0] = 100
slc[1] = 1000
slc[3] = 10
fmt.Println(arrr)
fmt.Println(slc)
var s3 []int
fmt.Println(s3 == nil)
fmt.Println(arrr == nil)
s1:= [][]int{{13,12},{32,45},{34,45},{98,67}}
fmt.Println(s1)
s2:= [][]string{
	{"freaks", "geeks"},
	{"Jocks", "Cracks"},
	{"GFG", "KGB"},
}
fmt.Println(s2)
sort.Ints(s1[0])
fmt.Println(s1)
sort.Ints(arrr)
fmt.Println(arrr)
sort.Strings(arr1[:])
fmt.Println(arr1)
}