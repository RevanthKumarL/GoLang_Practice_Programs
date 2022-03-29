//func(receiver_name type) method_name(parameter_lsit)(return_type){
//code
//}

package main
import "fmt"
//method with struct type receiver

//author structure
type author struct{
	name string
	branch string
	articles int
	salary int
}
//method with a receiver of author type
// func (a author) show(){ code: print statement with a.details}
func (a author) show(){
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published Articles: ", a.articles)
	fmt.Println("Salary: ", a.salary)
}
//main function
func main(){
	//intializing the values
	//of the author structure
	res1:= author{
		name: "Sonam",
		branch: "Computer Science",
		articles: 203,
		salary: 160000,
	}
	res2:= author{
		name: "Ritu",
		branch: "Electronics",
		articles: 043,
		salary: 100000,
	}
	//calling the method
	res1.show()
	res2.show()
}
