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
