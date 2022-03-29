package main
import "fmt"

type employee struct{

	name string
	company string
	empid int
}

func main(){
	emp:= employee{"RevanthKumar", "Meta", 141298}

	pts:= &emp
	fmt.Println(pts)

	fmt.Println(pts.company)
	
	pts.empid = 10071972 //only = is allowed, := redeclares

	fmt.Println(pts.empid)
}