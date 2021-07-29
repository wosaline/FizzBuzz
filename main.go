package main

import (
	"FizzBuzz/entity"
	"fmt"
)

func main() {
	res, err := entity.FizzBuzz(30, 3, 5, "", "")
	fmt.Println(res)
	fmt.Println(err)
}
