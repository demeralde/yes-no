package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var responses [2]string

	responses[0] = "Yes"
	responses[1] = "No"

	selectedResponseIndex := rand.Intn(len(responses))
	selectedResponse := responses[selectedResponseIndex]

	fmt.Println(selectedResponse)
}
