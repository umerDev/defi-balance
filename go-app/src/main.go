package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/umerDev/defi-balance/src/balance"
)

func main() {
	address := strings.Join(os.Args[1:], "")
	alephiumBalance, err := balance.GetBalance(address, balance.CreateUrl)

	if err != nil {
		log.Fatal("unable to get balance", err)
	}
	fmt.Println("alephieum balance: ", alephiumBalance)
}
