package stats

import (
	"github.com/SsSJKK/bank/v2/pkg/types"
)

//Avg func
func Avg(pyments []types.Payment) types.Money {
	var avgSumm types.Money
	var countPy types.Money
	for _, pyments := range pyments {
		if pyments.Status != types.StatusFail {
			avgSumm += pyments.Amount
			countPy++
		}
	}
	avgSumm /= countPy
	return avgSumm
}

//TotalInCategory func
func TotalInCategory(pyments []types.Payment, category types.Category) types.Money {
	var totalSumm types.Money
	for _, pyments := range pyments {
		if pyments.Category == category {
			if pyments.Status != types.StatusFail {
				totalSumm += pyments.Amount
			}
		}
	}
	return totalSumm
}

//CategoriesAvg func
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesAvg := make(map[types.Category]types.Money)
	countPayments := make(map[types.Category]types.Money)
	for _, payments := range payments {
		if payments.Status != types.StatusFail {
			countPayments[payments.Category]++
		}
	}
	for _, payments := range payments {
		if payments.Status != types.StatusFail {
			categoriesAvg[payments.Category] += payments.Amount / countPayments[payments.Category]
		}
	}

	return categoriesAvg
}
