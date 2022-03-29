package main
import "fmt"
import "strings"

// function function_name(para1, para2...type)type{
	//code...
//}

func joinstr(element...string)string{
	return strings.Join(element, "-")

}

func main(){
	
	fmt.Println(joinstr())
	
	fmt.Println(joinstr("geeks", "freaks"))
	fmt.Println(joinstr("geeks", "for", "freaks"))
	fmt.Println(joinstr("tit", "for", "tat"))

	element:= []string{"geeks", "for", "geeks"}
	fmt.Println(joinstr(element...))

}

