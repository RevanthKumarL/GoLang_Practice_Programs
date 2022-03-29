package main

import "fmt"

func main() {
	// Create and initialize a
	// variable with a string
	// shorthand declaration
	my_val1 := "Well done soldier"
	// Using var keyword
	var my_val2 string
	my_val2 = "Mission Accomplished"
	// Displaying strings
	fmt.Println("String 1: ", my_val1)
	fmt.Println("String 2: ", my_val2)
	my_val3:= "My Life\nMy Rules"
	my_val4:= `None of your business`
	my_val5:= `My Life\nMy Rules` //different quotes for 
								//printing as it is: ``
	fmt.Println(my_val3)
	fmt.Println(my_val4)
	fmt.Println(my_val5)
	mystr:= "Chal Dengai"
	fmt.Println("String: ",mystr)
	fmt.Println(mystr)
	for index, s:= range "Chal Dengai"{
		fmt.Printf("\nThe index number of %c is %d",s,index)
		//%c is for single letter in a string
	}
	str:= "Paisal levu ra labbe"
	for c:=0; c < len(str); c++{
	fmt.Printf("\nBytes = %v", str)
	}
	fmt.Println("\nThe length of mystr: ",len(mystr))
}
