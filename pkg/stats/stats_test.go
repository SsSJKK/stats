package stats

import (
	"reflect"
	"testing"

	"github.com/SsSJKK/bank/v2/pkg/types"
)

func Test_1(t *testing.T) {
	var payments = []types.Payment{
		{
			Amount:   2,
			Category: "1",
			Status:   types.StatusOk,
		},
		{
			Amount:   2,
			Category: "1",
			Status:   types.StatusOk,
		},
		{
			Amount:   3,
			Category: "2",
			Status:   types.StatusOk,
		},
	}
	result := CategoriesAvg(payments)
	expected := map[types.Category]types.Money{
		"1": 2,
		"2": 3,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, ex[ected: %v, actual: %v", expected, result)
	}
}
