package main

import (
	"fmt"
	"strings"
)

//comparison operators: ==, !=, >=, <=, <, >

func main(){
	str1:= "RevanthKumar"
	str2:= "Rev"
	str3:= "Lavanya"
	str4:= "Sivanandam"
	str5:= "RevathKumar"
	str6:= "RevKumar"
	str7:= "Charan"

	fmt.Println(str1 == str2)
	fmt.Println(str2==str3)
	fmt.Println(str4==str5)
	fmt.Println(str5==str6)
	fmt.Println(str6==str7)

	fmt.Println(str1!=str2)
	fmt.Println(str2!=str3)
	fmt.Println(str3!=str4)
	fmt.Println(str4!=str5)
	fmt.Println(str5!=str6)
	fmt.Println(str6!=str7)

	myslice:= []string{"RevanthKumar", "MySivanandam", "MyLavanya", "MyCharan", "A Beautiful Family"}

	fmt.Println("Slice:",myslice)

	res1:= "RevanthKumar" > "MySivanandam"
	fmt.Println(res1)

	res2:= "MyLavanya">"MySivanandam"
	fmt.Println(res2)

	res3:= "RevanthKumar">="MyCharan"
	fmt.Println(res3)

	res4:= "MyCharan"<="RevanthKumar"
	fmt.Println(res4)

	res5:= "RevanthKumar" == "REvanthKumar"
	fmt.Println(res5)

	res6:= "RevanthKumar" != "REVanthKumar"
	fmt.Println(res6)

	//compare() method
	/* 
	Return 0, if str1 == str2.
	Return 1, if str1 > str2.
	Return -1, if str1 < str2.
	*/
	fmt.Println(strings.Compare("RevanthKumar", "REvanthKumarLokku"))
	fmt.Println(strings.Compare("Charan", "RevanthKumar"))
	fmt.Println(strings.Compare("Sivanandam", "Lavanya"))
	fmt.Println(strings.Compare("Charan", "Lavanya"))


}