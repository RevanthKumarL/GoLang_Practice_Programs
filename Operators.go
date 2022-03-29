package main
import "fmt"

func main() {
	p:= 34
	q:= 20
	result1:= p + q
	fmt.Printf("Result of p+q = %d\n", result1)
	result2:= p - q
	fmt.Printf("\nResult of p-q = %d\n", result2)
	result3:= p*q
	fmt.Printf("\nResult of p*q = %d\n", result3)
	result4:= p/q
	fmt.Printf("\nResult of p/q = %d\n", result4)
	//modulus means reminder from the division b/w dividend and divisor
	result5:= p%q
	fmt.Printf("\nResult of p%%q = %d\n", result5)
	res1:= p==q
	fmt.Println(res1)
	res2:= p!=q
	fmt.Println(res2)
	res3:= p<q
	fmt.Println(res3)
	res4:= p>q
	fmt.Println(res4)
	res5:= p>=q
	fmt.Println(res5)
	res6:= p<=q
	fmt.Println(res6)
	var r int = 23
	var s int = 60
	if (r!=s && r<=s){
		fmt.Println("True")
	}
	if (r!=s || r<=s){
		fmt.Println("True")
	}
	if (!(r==s)){
		fmt.Println("True")
	}
	and:= p & q
	fmt.Println(and)
	or:= p | q
	fmt.Println(or)
	xor:= p ^ q
	fmt.Println(xor)
	leftshift:= p << 1
	fmt.Println(leftshift)
	rightshift:= p >> 1
	fmt.Println(rightshift)
	andnot:= p &^ q
	fmt.Println(andnot)
	var a int = 45
	var b int = 50
	a= b 
	fmt.Println(a)
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)
	c:= 4
	d:= &c //using the address of an operator and 
	//pointer indirection operator
	fmt.Println(*d)
	*d= 7
	fmt.Println(c)
	*d = 10
	fmt.Println(c)



	
}
