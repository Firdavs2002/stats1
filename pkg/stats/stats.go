package stats

import (
	"github.com/Firdavs2002/bank1/v2/pkg/types"
)

//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	// countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	indexPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
		indexPayments++
	}
	return sumPaymenys / indexPayments
}

// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys
}
