package service

import (
	"time"
	"fmt"
)

type CreditCardService struct {
	//vars if needed.
}

//Creates new CardService struct, by satisfying its dependency.
/*func NewCardService() *CardService {
	cService := new(CardService)
	//cService.cRepo = repository.NewCardRepository()
	return cService
}*/

//Posts transaction for the specified card.
func (cs *CreditCardService) PostTransaction(cNum string, amt float32, txnTime time.Time) (bool, error)  {
	//TODO: post transaction logic, validate if needed.
	fmt.Printf("CreditCardService: Positing transaction for card: %.2f \n", amt)
	return true, nil
}

//Fetches un-billed transactions.
func (cs *CreditCardService) ViewStatement(cNum string) []interface{} {
	return make([]interface{}, 0, 1)
}