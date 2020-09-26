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

func Test_2(t *testing.T) {
	first := map[types.Category]types.Money{
		"1": 2,
		"2": 3,
		"3": 3,
	}
	second := map[types.Category]types.Money{
		"1": 2,
		"2": 3,
		"4": 2,
	}

	result := PeriodsDynamic(first, second)
	expected := map[types.Category]types.Money{
		"1": 0,
		"2": 0,
		"3": -3,
		"4": 2,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, ex[ected: %v, actual: %v", expected, result)
	}
}
