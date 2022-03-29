package main
import (
	"fmt"	
)
func area(length, width int)int{
	Ar:= length*width
	return Ar
}
func area1(len, breadth, height int)int{
	Area:= len*breadth*height
	return Area
}
func areaTri(b, h int)int{
	Triangle:= (b*h)/2 
	return Triangle
}
func main(){
	fmt.Printf("Area of the rectangle is : %d", area(12, 10))
	fmt.Printf("\nArea of a solid rectangle is : %d", area1(23, 21,23))
	fmt.Printf("\nArea of a triangle is : %d", areaTri(7, 8))
}	