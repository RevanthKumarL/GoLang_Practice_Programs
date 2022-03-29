package main
import "fmt"

func main() {
	for i:= 0; i<7; i++{
		fmt.Println("OmNamahShivay")
	}
	i:= 0
	for i<3 {
		i+=2
	}
	fmt.Println(i)

	rvariable := []string{"rev", "aish", "rosh", "v", "shiv", "sid", "paksha"}
	for a, b:= range rvariable {
		fmt.Println(a, b)
	}
	for a, b:= range "bjhabfwfewy" {
		fmt.Printf("The index number of %U is %d\n", b, a)
	}
	//for loops using maps; map declare:=map[type]type
	mmap:= map[int]string {
		22:"geeks",
		33:"freaks",
		44:"nerds",
	}
	for key, value:= range mmap {
		fmt.Println(key, value)
	}

	chnl:= make(chan int)
	go func () {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i:= range chnl {
		fmt.Println(i)
	}
}