package main

import (
	"FizzBuzz/entity"
	"fmt"
)

func main() {
	res, err := entity.FizzBuzz(6, 7, 10, "fizz", "buzz")
	fmt.Println(res)
	fmt.Println(err)
}
