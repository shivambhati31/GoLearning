package main

import (
	"fmt"
	"net/http"
)

const myurl string = "https://api.publicapis.org/entries"

func main() {
	fmt.Println("Hello")
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
}
