package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	text, _ := reader.ReadString('\n')
	return text
}
