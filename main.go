package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the numbers: ")
	numbers, _ := reader.ReadString('\n')
	stringArr := strings.Fields(numbers)
	var numbersArr = []int{}
	for _,i := range stringArr{
		j,err := strconv.Atoi(i)
		if err != nil{
			panic(err)
		}else{
			numbersArr = append(numbersArr, j)
		}
	}

	for _,j := range numbersArr{
		fmt.Printf("You entered %d but current is %d \n",j,j+1)
	}

	

}