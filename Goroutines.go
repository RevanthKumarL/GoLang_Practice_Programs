package main

import (
	"fmt"
	"time"
)

func display(str string){
	for k:=0;k<7;k++ {

		time.Sleep(1*time.Second)
		fmt.Println(str)
	}
}

func main(){

	go display("Rev")
	display("Product Security Engineer")
}