package stats

import (
	"fmt"

	"github.com/SsSJKK/bank/pkg/types"
)

func ExampleAvg() {
	var pyments = []types.Payment{
		{
			Amount: 1,
		},
		{
			Amount: 2,
		},
		{
			Amount: 3,
		},
	}

	pyment := Avg(pyments)
	fmt.Println(pyment)

	//Output: 2
}

func ExampleTotalInCategory() {
	var pyments = []types.Payment{
		{
			Amount:   1,
			Category: "1",
		},
		{
			Amount:   2,
			Category: "1",
		},
		{
			Amount:   3,
			Category: "2",
		},
	}

	pyment := TotalInCategory(pyments, "1")
	fmt.Println(pyment)

	//Output: 3
}
