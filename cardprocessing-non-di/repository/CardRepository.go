package repository

import (
	"fmt"
	"time"
)

/*type CardRepository interface {
	PostTransaction(cardNumber string)
	GetTransactionHistory(cardNumber string)
}*/

type CardRepository struct {
	//vars if needed.
}

//Creates new CardRepository struct.
func NewCardRepository() *CardRepository  {
	return new(CardRepository)
}

//Posts transaction for the specified card.
func (cR *CardRepository) PostTransaction(cN string, amt float32, txTime time.Time)  (bool, error) {
	fmt.Printf("CardRepository: Posting transaction for card: %s . \n", cN)
	return true, nil
}

//Fetches un-billed transactions.
func (cR *CardRepository) ViewUnbilledTransactions(cN string) ([]interface{}){
	fmt.Println("Get Transaction History in Repository invoked.")
	change :=make([]interface{}, 0, 10)
	return change
}

