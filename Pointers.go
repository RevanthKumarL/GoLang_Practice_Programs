//pointer= variable used to store the memory address of another var
package main
import "fmt"
func main(){
	x:= 0xFF
	y:= 0x9C
	fmt.Printf("Type: %T", x)
	fmt.Printf("\nValue of x in hex: %X\n", x)
	fmt.Printf("Value of x in decimal: %v\n", x)
	fmt.Printf("\nType: %T\n", y)
	fmt.Printf("Value of y in hex: %X\n",y)
	fmt.Printf("Value of y in decimal: %v\n\n",y)
	// * is the dereferencing operator to declare pointer var
	// and access the value stored in the address
	// & is the address operator used to return the addr of var
	/* var a = 45
	init of pointer s with memory address of var a
	var s *int = &a	*/
	var r int = 5748
	//declaration of pointer
	var p *int 
	//init of pointer
	p = &r
	// can also be written as : var p *int = &r
	fmt.Println("Value stored in r:",r)
	fmt.Println("Address stored in r:",&r)
	fmt.Println("Value stored in p:",p)
	fmt.Println("")
	var z = 458
	var q = &z
	// also be written like: 
	//z:= 458
	//q:= &z
	fmt.Println("Value stored in y:", z)
	fmt.Println("Addr of y:",&z)
	fmt.Println("Addr stored in pointer var p:",q)
	//derefering the pointer
	//using the * operator before a pointer
	//var to access the value stored
	//at var which is pointing at
	fmt.Println("Value stored in z(*q):",*q) //*q = z value
	//we can also change the value of the z 
	//assign the new value to the pointer
	*q = 500
	fmt.Println("")
	fmt.Println("Value stored in z(*p) after change:",*q)

}