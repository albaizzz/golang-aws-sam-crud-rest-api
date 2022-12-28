package shared

import "fmt"

type Logger struct {
}

func (logger Logger) Debug(message string) {
	fmt.Print(message)
}

func (logger Logger) Error(error error) {
	fmt.Println(error.Error())
}
