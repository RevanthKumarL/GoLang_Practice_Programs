package main
import (
	"fmt"
	"sort"
	"strings"
	"time"
)
func main() {
	s:= []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	//sort.ints (s); ints here sorts (s) in an increasing order
	fmt.Println("Sorted slice: ", s)

	//finding the index
	fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))

	//finding the time
	fmt.Println("Time: ", time.Now())
	fmt.Println("Time in unix: ", time.Now().Unix())
}