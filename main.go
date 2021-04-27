package main

import (
	"fmt"

	"github.com/kutukmu/kg_kutukmu_2021/features"
)


func main() {

	fmt.Println("Please start entering:")

	str :=features.ReadNumbers()
	numArr := features.ConvertStringToIntArr(str)
	result := features.ConvertToPhonetic(numArr)
	fmt.Print(result)
}