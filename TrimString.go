//syntax: func Trim(str string, cutstr string) string

package main
import "fmt"
import "strings"
func main(){
	str1:= "!!I need more loyal, hearty friends!!"
	str2:= "@@Not a resolution, a serious decision@@"
	fmt.Println("Strings before trimming:")
	fmt.Println("String1: ",str1)
	fmt.Println("String2: ",str2)
	//result:= strings.Trim(str_name, "Eliminating symbol")
	res1:= strings.Trim(str1,"!")
	res2:= strings.Trim(str2, "@")
	fmt.Println("After Trimming:")
	fmt.Println("Result1: ",res1)
	fmt.Println("Result2: ",res2)
	//TrimLeft 
	res3:= strings.TrimLeft(str1,"!")
	res4:= strings.TrimLeft(str2,"@")
	fmt.Println("")
	fmt.Println(res3)
	fmt.Println(res4)
	//TrimRight
	res5:= strings.TrimRight(str1,"!")
	res6:= strings.TrimRight(str2,"@")
	fmt.Println("")
	fmt.Println(res5)
	fmt.Println(res6)
	/*Trimming the spaces of strings - space b/w the quotes
	like eg. "  I Love You  "        						*/ 
	res7:= strings.TrimSpace(str1)
	res8:= strings.TrimSpace(str2)
	fmt.Println("")
	fmt.Println(res7)
	fmt.Println(res8)
	//Trimming the suffix
	str3:= "Loyal,but honest"
	str4:= "Harsh with words,but mushy with feelings"
	res9:= strings.TrimSuffix(str3, "but honest")
	res10:= strings.TrimSuffix(str4, "but mushy with feelings")
	fmt.Println(res9)
	fmt.Println(res10)
	res11:= strings.TrimPrefix(str3, "Loyal,")
	res12:= strings.TrimPrefix(str4, "Harsh with words,")
	fmt.Println(res11)
	fmt.Println(res12)



}