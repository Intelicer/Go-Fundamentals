package main

import (
	"08-PackageScope/package1"
	"08-PackageScope/package2"
	"fmt"
)

func main() {
	package1.SayHello("Mahmoud")
	AvgInIntSlice(package2.Numbers)
}

func AvgInIntSlice(slice []int) {
	avgNum := 0
	for _, value := range slice {
		avgNum += value
	}
	fmt.Println("avg: ", (avgNum / len(slice)))
}
