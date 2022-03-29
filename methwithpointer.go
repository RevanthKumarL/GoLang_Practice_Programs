package main
import (
	"fmt"
)
/*func (p *type) method_name(...type) type{ //code }
author struct */
type author struct{
	name string
	branch string
	articles int
}
//method with a receiver of author type
func (a *author) show(abranch string) {
	(*a).branch = abranch	
}
func (a *author) show_1(abranch string) {
	(*a).branch = abranch
}
func (a *author) show_2(abranch string) {
	(*a).branch = abranch
	
	a.name = "Gourav"	
}
//main function
func main(){
	res:= author{
		name: "Sona",
		branch: "CSE",
	}
	fmt.Println("Authors name: ", res.name)
	fmt.Println("Branch name(Before): ", res.branch)
	//creating a pointer
	p:= &res
	//calling the show method
	p.show("ECE")
	fmt.Println("Authors name: ", res.name)
	fmt.Println("Branch name(After): ", res.branch)
	res.show_1("EEE")
	fmt.Println("Branch name(After): ", res.branch)
	(&res).show_2("Gourav")
	fmt.Println("Authors name(After): ", res.name)
}