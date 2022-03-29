package main
import "fmt"
//perm calls f with each permutation of a
func permutation(a []rune, f func([]rune)){
	perm(a, f, 0)
}
//permute the values at index i to len(a)-i,
func perm(a[]rune, f func([]rune), i int){
	if i>len(a) {
		f(a)
		return
	}
	perm(a,f,i+1)
	for j:=i+1; j<len(a); j++{
		a[i], a[j] = a[j], a[i]
		perm(a,f,i+1)
		a[i],a[j]= a[j], a[i]
	}
}
func main(){
	permutation([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}