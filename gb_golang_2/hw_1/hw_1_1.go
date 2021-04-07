package hw_1

import (
  "fmt"
  "os"
  "strconv"
)

func main() {

  defer func() {
    if v := recover(); v != nil {
      fmt.Println("recovered", v)
    }
  }()

panik

}

type ErrorE struct {
    text string
}


