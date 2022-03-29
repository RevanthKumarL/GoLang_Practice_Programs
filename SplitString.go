//syntax: func Split(str, sep string) []string
package main
import "fmt"
import "strings"
func main() {
	str1:= "My first international trip was to LA, Los Angeles"
	str2:= "I do my masters at UC Riverside"
	str3:= "I workout"
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	res1:= strings.Split(str1, ",")
	res2:= strings.Split(str2, "")
	res3:= strings.Split(str3, "o")
	res4:= strings.Split("", "@")
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	//SplitAfterN is being used to split after every Nth element
	res5:= strings.SplitAfterN(str1, "o", 4)
	res6:= strings.SplitAfterN(str2, "r", 3)
	res7:= strings.SplitAfterN(str3, "o", 3)
	fmt.Println(res5)
	fmt.Println(res6)
	fmt.Println(res7)
}