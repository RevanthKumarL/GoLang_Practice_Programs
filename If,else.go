package main
import "fmt"

func main() {
	// if (condition){
		// statements to execute if
		//condition is true
	//}
	var v int = 1200

	if(v < 1000) {
		fmt.Println("v is less than 1000")
	} else {

	fmt.Printf("Value of v is greater than 1000\n")
	}

	var v1 int = 400
	var v2 int = 700
	if (v1 == 400){
		if (v2 == 700){
			fmt.Println("Value of v1 is 400 and v2 is 700")

		}
	}
	var v3 int = 700
	if (v3 == 100){
		fmt.Println("Value of v3 is 100")
	} else if(v3 == 200){
		fmt.Println("Value of v3 is 200")
	} else if (v3 == 300) {
		fmt.Println("Value of v3 is 300")
	} else {
		fmt.Println("None of the values are matching")
	}

}