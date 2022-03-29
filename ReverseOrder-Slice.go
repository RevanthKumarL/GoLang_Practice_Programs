package main
import "fmt"
func reverse(sw []int){
	for a,b:= 0, len(sw)-1;a<b;a,b = a+1,b-1{
		sw[a], sw[b] = sw[b],sw[a]
	}
}
func reversestring(rs []string){
	for c,d:= 0, len(rs)-1; c<d; c,d=c+1,d-1{
		rs[c],rs[d]=rs[d],rs[c]
	}
}
func main(){
	x:= []int{9,8,7,6,5,4,3,2,1}
	y:= []string{"Look", "Who's", "Here"}
	reversestring(y)
	reverse(x)
	fmt.Println(x)
	fmt.Println(y)
}