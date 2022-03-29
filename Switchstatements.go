package main
import (
	"fmt"
)
func main() {
	switch day:=5; day{
		case 1: 
		fmt.Println("Monday")
		case 2: 
		fmt.Println("Tuesday")
		case 3: 
		fmt.Println("Wednesday")
		case 4: 
		fmt.Println("Thursday")
		case 5: 
		fmt.Println("Friday")
		case 6: 
		fmt.Println("Saturday")
		case 7: 
		fmt.Println("Sunday")
		default: 
		fmt.Println("Invalid")
		}
		var value int = 2
		switch {
		case value == 1:
		fmt.Println("Hello")
		case value == 2:
		fmt.Println("Bonjour")
		case value == 3: 
		fmt.Println("Namastay") 
		default:
		fmt.Println("Invalid")
		}
		var value1 string = "Five"
		switch value1 {
		case "one": 
		fmt.Println("C#")
		case "two", "three": 
		fmt.Println("Go")
		case "four", "five", "six": 
		fmt.Println("Java")
		}
		var value2 interface{}
		switch q:= value2.(type) {
			case bool:  
			fmt.Println("Value is of boolean type")
			case float64: 
			fmt.Println("Value is of float64 type")
			case int:
			fmt.Println("Value is of int type")
			default:
			fmt.Printf("Value is of type: %T", q)
		}
	}
