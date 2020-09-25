package stats

import (
	"github.com/SsSJKK/bank/pkg/types"
)

//Avg func
func Avg(pyments []types.Payment) types.Money {
	var avgSumm types.Money
	for _, pyments := range pyments {
		avgSumm += pyments.Amount
	}
	avgSumm /= types.Money(len(pyments))
	return avgSumm
}

//TotalInCategory func
func TotalInCategory(pyments []types.Payment, category types.Category) types.Money {
	var totalSumm types.Money
	for _,pyments := range pyments{
		if (pyments.Category == category){
			totalSumm += pyments.Amount
		}
	}
	return totalSumm
}