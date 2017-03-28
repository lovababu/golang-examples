package model

import (
	"time"
	"fmt"
	"strings"
	"errors"
	"github.com/lovababu/golang-examples/cardprocessing-di/service"
)


type CreditCard struct {
	Number 		string
	ExpiryDate 	time.Time
	CVV2 		int16
	ISSUER 		string
	Limit 		float32
	PaymentDueDate 	time.Time
	MinPayment 	float32
	LateFee 	float32
	CardService 	*service.CardService `inject:"CreditCard"`
}

type DebitCard struct {
	Number 		string
	ExpiryDate 	time.Time
	CVV2 		int16
	ISSUER 		string
	Balance 	float32
	CardService 	*service.CardService `inject:"DebitCard"`
}

func NewCreditCard() *CreditCard {
	ccard:= new(CreditCard)
	//card.CardService = service.NewCreditCardService()
	return ccard
}

func NewDebitCard() *DebitCard {
	dcard := new(DebitCard)
	//card.CardService = service.NewDebitCardService()
	return dcard;
}


//Validates the Credit Card. valid if start with 4526, else not.
func (cc *CreditCard)Validate() (bool, error) {
	fmt.Println("Validating Credit Card..")
	//TODO: validation specific to credit card.
	if strings.HasPrefix(cc.Number, "4526") {
		fmt.Println("Card validation passed: ", cc.Number)
		return true, nil
	} else {
		return false, errors.New("Invalid Credit card.")
	}

	return false, nil
}

//Validates the Debit Card. valid if start with 4528, else not.
func (dc *DebitCard)Validate() (bool, error) {
	fmt.Println("Validating Credit Card..")
	//TODO: validation specific to credit card.
	if strings.HasPrefix(dc.Number, "4526") {
		fmt.Println("Card validation passed: ", dc.Number)
		return true, nil
	} else {
		return false, errors.New("Invalid Debit card.")
	}

	return false, nil
}