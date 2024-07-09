package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/umerDev/defi-balance/src/balance"
)

func main() {
	address := strings.Join(os.Args[1:], "")
	fmt.Println(address)
	alephiumBalance := balance.GetBalance(address)
	fmt.Println("alephieum balance: ", alephiumBalance)
}
