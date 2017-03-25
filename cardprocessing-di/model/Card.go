package model

import (
	"time"
	"fmt"
	"strings"
	"errors"
	"github.com/lovababu/golang-examples/cardprocessing-di/service"
)

type Card struct {
	Number 		string
	ExpiryDate 	time.Time
	CVV2 		int16
	ISSUER 		string
	Limit 		float32
	PaymentDueDate 	time.Time
	MinPayment 	float32
	LateFee 	float32
	Transactions[] 	Statement

	CardService 	*service.CardService `inject:""`
}

func NewCreditCard() *Card {
	card:= new(Card)
	//card.CardService = service.NewCardService()
	return card
}


type Statement struct {
	TransactionDate time.Time
	Amount 		float32
	Type 		string
	Remarks 	string
	TransactionId 	string
}

//Validates the Credit Card. valid if start with 4526, else not.
func (cc *Card) Validate() (bool, error) {
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