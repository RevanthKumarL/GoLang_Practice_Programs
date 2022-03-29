package main

import "fmt"

func main(){
	//create and init strings
	//using var keyword
	str1:= "Hello "

	var str2 string
	str2 = "Meta"

	fmt.Println("New string1:",str1+str2)

	str3:= "Welcome to"
	str4:= "Menlo Park"

	result:= str3+" to "+str4

	fmt.Println("New string2:", result)
}