package main

import (
	"github.com/lovababu/golang-examples/cardprocessing-di/model"
	"time"
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/lovababu/golang-examples/cardprocessing-di/service"
	"os"
)

func main() {

	var ig  inject.Graph

	//Init Credit card.
	ccard := model.NewCreditCard()
	ccard.Number = "452634259080076"
	ccard.CVV2 = 786
	ccard.ExpiryDate = time.Now().AddDate(21, 0, 0)
	ccard.Limit = 120000.00

	//Init Debit card.
	dcard := model.NewDebitCard()
	dcard.Number = "452634259080076"
	dcard.CVV2 = 786
	dcard.ExpiryDate = time.Now().AddDate(21, 0, 0)
	dcard.Balance = 100000.00

	var ccService service.CreditCardService
	var dcService service.DebitCardService

	//Providing objects to injector.
	err := ig.Provide(
		&inject.Object{Value: ccard},
		&inject.Object{Value: dcard},
		&inject.Object{Value: &ccService, Name: "CreditCard"},
		&inject.Object{Value: &dcService, Name: "DebitCard"},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = ig.Populate()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//Credit card validation and post transaction.
	isValid, err := ccard.Validate()

	if err != nil {
		fmt.Println("Card is not validated, so exiting the process.")
	} else if isValid {

		fmt.Println(ccard)
		flag, err := service.PostTransaction(ccard.Number, 100.00, time.Now(), ccard.CardService)
		fmt.Println("Is post transaction success: ", flag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Card is not validated, so exiting the process.")
	}

	//Debit card validation and post transaction.
	isValid, err = dcard.Validate()

	if err != nil {
		fmt.Println("Card is not validated, so exiting the process.")
	} else if isValid {

		flag, err := service.PostTransaction(dcard.Number, 100.00, time.Now(), dcard.CardService)
		fmt.Println("Is post transaction success: ", flag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Card is not validated, so exiting the process.")
	}
}
