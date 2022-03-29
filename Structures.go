package main
import "fmt"
//defining a struct 
type Address struct {
	Name string
	City string
	Pincode int
}
type Car struct{
	Name, Model, Color string
	WeightinKG float64
}
func main(){
	var a Address
	fmt.Println(a)
	a1:= Address{"Akshay", "Manali", 349934}
	fmt.Println("Address1: ", a1)

	a2:= Address{Name: "Samhita", City: "Godavari", Pincode: 3943434}
	fmt.Println("Address2: ", a2)

	a3:= Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)

	c:= Car{Name: "Ferrari", Model: "GTC4", Color: "Maranello Red", WeightinKG: 1920}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)
	c.Color = "Night Skin Black"
	fmt.Println("Car: ", c)
}