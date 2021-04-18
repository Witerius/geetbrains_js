package main

import (
	"fmt"
)

var workers = make(chan int, 1)

func work() {
	for i := 0; i <= 1000; i++ {
		workers <- 1

		go func(job int) {
			defer func() {
				<-workers
			}()

			fmt.Println(job)
		}(i)
	}

}

func _main() {
	work()
}
