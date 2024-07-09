package balance

import (
	"fmt"
	"testing"
)

func Test_GetBalance(t *testing.T) {
	address := "1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH"

	alephiumBalance := GetBalance(address)

	if alephiumBalance != "0" {
		t.Error("Unexpected balance amount")
	}
	fmt.Println("balance is: ", alephiumBalance)
}
