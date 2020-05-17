package main

import (
	"fmt"

	"github.com/t-bonatti/stock-follower/client"
)

func main() {
	fmt.Println("Hello world")
	statusInvest := client.NewStatusInvest()

	statusInvest.Get("abev3")
}
