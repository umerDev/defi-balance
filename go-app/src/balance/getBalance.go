package balance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type alephiumBody struct {
	Balance           string `json:"balance"`
	BalanceHint       string `json:"balanceHint"`
	LockedBalance     string `json:"lockedBalance"`
	LockedBalanceHint string `json:"lockedBalanceHint"`
	UtxoNum           int    `json:"utxoNum"`
}

type URLGenerator func(address string) string

func CreateUrl(address string) string {
	url := fmt.Sprintf("https://wallet.mainnet.alephium.org/addresses/%s/balance", address)
	return url
}

func GetBalance(address string, generateURL URLGenerator) (string, error) {
	url := generateURL(address)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result alephiumBody
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return result.BalanceHint, nil
}
