package balance

import (
	"testing"
)

func TestCreateUrl(t *testing.T) {
	address := "testAddress"
	expectedUrl := "https://wallet.mainnet.alephium.org/addresses/testAddress/balance"
	actualUrl := createUrl(address)

	if actualUrl != expectedUrl {
		t.Errorf("Expected URL %s but got %s", expectedUrl, actualUrl)
	}
}
