package main

import (
	"github.com/lovababu/golang-examples/cardprocessing-non-di/model"
	"time"
	"fmt"
)

func main() {

	ccard := model.NewCreditCard()
	ccard.Number = "452634259080076"
	ccard.CVV2 = 786
	ccard.ExpiryDate = time.Now().AddDate(21, 0, 0)
	ccard.Limit = 1200000.00

	isValid, err := ccard.Validate()

	if err != nil {
		fmt.Println("Card is not validated, so exiting the process.")
	} else if isValid {
		ccard.CardService.PostTransaction(100.00, time.Now())
	} else {
		fmt.Println("Card is not validated, so exiting the process.")
	}
}
