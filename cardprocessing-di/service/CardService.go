package service

import (
	"time"
)


type CardService interface {

	PostTransaction(cNum string, amt float32, txnTime time.Time) (bool, error)
	ViewStatement(cNum string) []interface{}
}

func PostTransaction(cNum string, amt float32, txnTime time.Time, cardService CardService) (bool, error) {
	return cardService.PostTransaction(cNum, amt, txnTime)
}

func ViewStatement(cNum string, cardService CardService) []interface{} {
	return cardService.ViewStatement(cNum)
}
