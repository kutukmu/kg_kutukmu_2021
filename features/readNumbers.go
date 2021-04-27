package features

import (
	"bufio"
	"os"
)

func ReadNumbers() *string{
	reader := bufio.NewReader(os.Stdin)
	numbers, _ := reader.ReadString('\n')
	return &numbers
}

