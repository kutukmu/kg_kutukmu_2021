package features

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadNumbers() *string{
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return &str
}

func ConvertStringToIntArr(str *string) []int{

	stringsArr := strings.Fields(*str)
	numbersArr := []int{}
	for _,j := range stringsArr{
		i,err := strconv.Atoi(j)
		if err != nil{
			panic(err)
		}else{
			numbersArr = append(numbersArr, i)
		}
	}

	return numbersArr

}

func ConvertToPhonetic(numbers []int)string {
	obj := map[int]string{0:"Zero",1:"One", 2:"Two", 3:"Three", 4:"Four", 5:"Five", 6:"Six", 7:"Seven", 8:"Eight", 9:"Nine"}
	arr := []string{}
	for _,j := range numbers{
		var str string = ""
		val := j
		for val > 0{
			digit := val % 10
			val = val /10
			str = obj[digit] + str
			
		}
		arr = append(arr, str)
		
	}

	return strings.Join(arr[:], ",")
}
