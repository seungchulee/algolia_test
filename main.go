package main

import (
	"fmt"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	fmt.Println("Hello, World!")

}
