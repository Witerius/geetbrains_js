package hw_1

import (
	"fmt"
	"runtime/debug"
)

func main() {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
			err := New("error!")
			fmt.Println(err)
		}
	}()

	panic("AAA!!")

}

type ErrorE struct {
	text  string
	trace string
}

func (e ErrorE) Error() string {
	panic("implement me")
}

func New(text string) error {
	return &ErrorE{
		text:  text,
		trace: string(debug.Stack()),
	}
}
