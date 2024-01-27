package main

import (
	"fmt"
	"net/http"
)

func main() {
	var name string
	name = "Gladys"
	name2 := "Samantha"
	fmt.Println(name)
	fmt.Println(name2)

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%+v", res.Body)

}
