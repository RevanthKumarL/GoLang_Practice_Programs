package main
import (
	"bytes"
	"fmt"
)
//syntax: func Trim(ori_slice[]byte, cut_string string) []byte
func main(){
	slice1:= []byte{'!','!','@','I','L','o','v','e','Y','o','u'}
	slice2:= []byte{'*','(','T','o','T','h','e','M','o','o','n','&','B','a','c','k'}
	slice3:= []byte{'^','%','#','P','.','S',':','R','e','v'}
	//displaying slices
	fmt.Printf("\nOriginal Slices: \n")
	fmt.Printf("Slice 1 is: %s\n",slice1)
	fmt.Printf("Slice 2 is: %s\n",slice2)
	fmt.Printf("Slice 3 is: %s\n",slice3)
	//trimming the specified leading
	res1:= bytes.Trim(slice1, "!@)")
	res2:= bytes.Trim(slice2, "*(")
	res3:= bytes.Trim(slice3, "^%#")
	fmt.Printf("\nNew Slice:\n")
	fmt.Printf("Slice 1: %s\n", res1)
	fmt.Printf("Slice 2: %s\n", res2)
	fmt.Printf("Slice 3: %s\n", res3)
	//use the printf to get the strings printed 
	//or else the byte value is printed rather than the value
	result1:= bytes.Trim([]byte("###@@@Labyrinth###@@@"), "#@")
	result2:= bytes.Trim([]byte("^^^***Maze^^^***"), "^*")
	result3:= bytes.Trim([]byte("%LoneWolf%"), "%")
	fmt.Printf("\nFinal Slice:")
	fmt.Printf("\n%s\n", result1)
	fmt.Printf("%s\n", result2)
	fmt.Printf("%s", result3)
}