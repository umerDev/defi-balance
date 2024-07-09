package balance

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type alephiumBody struct {
	Balance           string `json:"balance"`
	BalanceHint       string `json:"balanceHint"`
	LockedBalance     string `json:"lockedBalance"`
	LockedBalanceHint string `json:"lockedBalanceHint"`
	UtxoNum           int    `json:"utxoNum"`
}

func createUrl(address string) string {
	url := fmt.Sprintf("https://wallet.mainnet.alephium.org/addresses/%s/balance", address)
	return url
}

func GetBalance(address string) string {
	url := createUrl(address)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result alephiumBody
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("unable to marshal json")
	}
	return result.Balance
}
