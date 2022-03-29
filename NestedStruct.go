package main
import "fmt"
//creating struct
type Author struct{
	name string
	branch string
	year int
	spouse_name string
	marital_status string
}
//creating nested struct
type HR struct{
	//struct as a field
	details Author
}
type Husband struct{
	name_hubby string
	salary int
}
type HIS struct{
	details Husband
}
func main() {
	//initializing the fields
	//of the structure
	result:= HR{
		details: Author{"sona", "CSE", 2019, "Rajat", "yes"},
	}
	result1:= HIS{
		details: Husband{"Rajat", 160000},
	}
	//display the values
	fmt.Println("Details of the author")
	fmt.Println(result)
	fmt.Println("Details of the Husband")
	fmt.Println(result1)
}