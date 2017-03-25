package service

import (
	"github.com/lovababu/golang-examples/cardprocessing/repository"
	"time"
	"fmt"
)

type CardService struct {
	//vars if needed.
	cRepo 	*repository.CardRepository
	cNum 	string
}

//Creates new CardService struct, by satisfying its dependency.
func NewCardService() *CardService {
	cService := new(CardService)
	cService.cRepo = repository.NewCardRepository()
	return cService
}

//Posts transaction for the specified card.
func (cs *CardService) PostTransaction(amt float32, txnTime time.Time) (bool, error)  {
	//TODO: post transaction logic, validate if needed.
	fmt.Printf("CardService: Positing transaction for card: %.2f \n", amt)
	return cs.cRepo.PostTransaction(cs.cNum, amt, txnTime)
}

//Fetches un-billed transactions.
func (cs *CardService) ViewUnBilledTransactions() []interface{} {
	return cs.cRepo.ViewUnbilledTransactions(cs.cNum)
}











