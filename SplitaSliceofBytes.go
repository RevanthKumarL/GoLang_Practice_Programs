package main
import (
	"bytes"
	"fmt"
)
func main(){
	//syntax: func Split(o_slice, sep []byte) [][]byte
	//creating and init the slice of bytes
	slice1:= []byte{'!','!','I','L','o','v','e','Y','o','u','#','#'}
	slice2:= []byte{'T','o','M','o','o','n','&','B','a','c','k'}
	slice3:= []byte{'%','%','P','.','S',':','R','e','v'}
	fmt.Println("Original Slices:")
	fmt.Printf("%s\n",slice1)
	fmt.Printf("%s\n",slice2)
	fmt.Printf("%s\n",slice3)
	res1:= bytes.Split(slice1, []byte("!!"))
	res2:= bytes.Split(slice2, []byte(""))
	res3:= bytes.Split(slice3, []byte("%"))
	fmt.Printf("\nAfter Splitting:\n")
	fmt.Printf("%s\n",res1)
	fmt.Printf("%s\n",res2)
	fmt.Printf("%s\n",res3)
	res4 := bytes.Split([]byte("****Welcome, to, GeeksforGeeks****"), 
                                                         []byte("****"))
      
    res5 := bytes.Split([]byte("Learning x how x to x trim x a x slice of bytes"),
                                                                     []byte("x"))
    
    res6 := bytes.Split([]byte("GeeksforGeeks, Geek"), []byte(","))    
    res7 := bytes.Split([]byte(""), []byte(","))
    // Display the results
    fmt.Printf("\nFinal Value:\n")
    fmt.Printf("Slice 1: %s", res4)
    fmt.Printf("\nSlice 2: %s", res5)
    fmt.Printf("\nSlice 3: %s", res6)
    fmt.Printf("\nSlice 4: %s", res7)
}