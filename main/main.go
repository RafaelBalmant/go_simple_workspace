package main

import (
	"fmt"
	error "learn/error"
)

func main() {
	res, err := error.HandlerError(false)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
