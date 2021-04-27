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
