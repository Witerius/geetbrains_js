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

  for i := 0; i < 1000000; i++ {
    createFile(int64(i))
  }

}

func createFile(number int64) {
  file, err := os.Create("hello" + strconv.FormatInt(number, 10) + ".txt")
  if err != nil {
    fmt.Println("Unable to create file:", err)
    os.Exit(1)
  }
  defer file.Close()
  fmt.Println("created", number)
}