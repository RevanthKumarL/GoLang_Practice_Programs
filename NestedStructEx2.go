package main
import (
	"fmt"
)
type Student struct{
	name string
	branch string
	year int
}
type Teacher struct{
	name string
	subject string
	experience int
	details Student
}
func main(){
	result:= Teacher{
		name: "Suman",
		subject: "Java",
		experience: 5,
		details: Student{"Bobby", "CSE", 2},
	}
	fmt.Println("Details of the Teacher")
	fmt.Println("Teachers Name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.experience)
	fmt.Println("\nDetails of the Student")
	fmt.Println("Students Name: ", result.details.name)
	fmt.Println("Students Branch: ", result.details.branch)
	fmt.Println("Year: ", result.details.year)
}