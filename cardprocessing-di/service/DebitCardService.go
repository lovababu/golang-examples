package service

import (
	"time"
	"fmt"
)

type DebitCardService struct {
	//vars if needed.
	CNum 	string
}

//Creates new CardService struct, by satisfying its dependency.
/*func NewCardService() *CardService {
	cService := new(CardService)
	return cService
}*/

//Posts transaction for the specified card.
func (cs *DebitCardService) PostTransaction(amt float32, txnTime time.Time) (bool, error)  {

	fmt.Printf("DebitCardService: Positing transaction for card: %.2f \n", amt)

	return true, nil
}

//Fetches Statement.
func (cs *DebitCardService) ViewStatement() []interface{} {
	fmt.Printf("DebitCardService: View statement")
	return make([]interface{}, 0, 1)
}