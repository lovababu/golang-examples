package main

import (
	"github.com/lovababu/golang-examples/cardprocessing-di/model"
	"time"
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/lovababu/golang-examples/cardprocessing-di/service"
	"os"
	"github.com/lovababu/golang-examples/cardprocessing/repository"
)

func main() {

	var ig  inject.Graph

	ccard := model.NewCreditCard()
	ccard.Number = "452634259080076"
	ccard.CVV2 = 786
	ccard.ExpiryDate = time.Now().AddDate(21, 0, 0)
	ccard.Limit = 1200000.00

	//var card model.Card
	var cService service.CardService
	var cRepository repository.CardRepository

	err := ig.Provide(
		&inject.Object{Value: ccard},
		&inject.Object{Value: &cService},
		&inject.Object{Value: &cRepository},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := ig.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	isValid, err := ccard.Validate()

	if err != nil {
		fmt.Println("Card is not validated, so exiting the process.")
	} else if isValid {
		ccard.CardService.PostTransaction(100.00, time.Now())
	} else {
		fmt.Println("Card is not validated, so exiting the process.")
	}
}
