package main

import(
	"fmt"
	"math"
)

func main(){
	// SayGreetings("Mahmoud")
	// SayHaveNiceDay("Mahmoud")

	names := []string{"Raghd", "Mahmoud", "Ahmed", "Rezan"}
	CycleNames(names, SayGreetings)
}


func SayGreetings(name string){
	fmt.Printf("Good morning %v\n", name)
}

func SayHaveNiceDay(name string){
	fmt.Printf("Have a nice day %v\n", name)
}

func circleArea(raidus float64) float{
	return math.Pi * math.Pow(raidus)
}

func CycleNames(names []string, functionToCall func(string)){

	for _, v := range names {
		functionToCall(v)
	}

}