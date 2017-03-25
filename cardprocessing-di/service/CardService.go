package service

import (
	"github.com/lovababu/golang-examples/cardprocessing-di/repository"
	"time"
	"fmt"
)


type CardService struct {
	//vars if needed.
	CRepo 	*repository.CardRepository `inject:""`
	CNum 	string
}

//Creates new CardService struct, by satisfying its dependency.
/*func NewCardService() *CardService {
	cService := new(CardService)
	//cService.cRepo = repository.NewCardRepository()
	return cService
}*/

//Posts transaction for the specified card.
func (cs *CardService) PostTransaction(amt float32, txnTime time.Time) (bool, error)  {
	//TODO: post transaction logic, validate if needed.
	fmt.Printf("CardService: Positing transaction for card: %.2f \n", amt)
	return cs.CRepo.PostTransaction(cs.CNum, amt, txnTime)
}

//Fetches un-billed transactions.
func (cs *CardService) ViewUnBilledTransactions() []interface{} {
	return cs.CRepo.ViewUnbilledTransactions(cs.CNum)
}