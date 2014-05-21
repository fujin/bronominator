package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math"
	"sort"
)

// # [1, 2, 3] 4 => 4
// # 1 + 1 + 1 + 1 = 4
// # 1 + 1 + 2 = 4
// # 1 + 3 = 4
// # 2 + 2 = 4
// # alex.favaro@affirm.com

// Denominations is the a float64slice of denominations
type Denominations []float64

// Knowing the correct denomination is way cooler than knowing the number of combinations
var denominations = Denominations{
	1.0,
	2.0,
	3.0,
	// 1.1,
	// 1.2,
	// 1.3,
	// 1.4,
	// 1.5,
	// 1.6,
	// 1.7,
	// 1.8,
	// 1.9,
}

const amount = 4.0

var denomSum = []float64{}

// exactChange is used to to calcualte the exact change types
// this should probably be a method on Denomination type.
func exactChange(amount float64, denomination float64) {
	for i := denomination; i < float64(len(denominations)); i++ {
		if denominations[int(i)] > amount {
			// Current value is higher than what is next to calculate
			exactChange(amount, denomination+1)
		} else {
			denomSum[int(i)], _ = math.Modf(amount / denominations[int(i)])
			exactChange(math.Mod(amount, denominations[int(i)]), denomination+1)
			return
		}
	}
}

// Reverse method for denominations type
func (d *Denominations) Reverse() {
	sort.Sort(
		sort.Reverse(
			sort.Float64Slice(*d), // pointer de-ref
		),
	)
}

func main() {

	fmt.Println("calculating appropriate denominations for amount:", amount)
	denominations.Reverse()
	denomSum = make([]float64, len(denominations), len(denominations))

	// Rub some concurrency sauce on it
	exactChange(amount, 0)

	// spew
	spew.Dump(amount)
	spew.Dump(denominations)
	spew.Dump(denomSum)

	fmt.Printf("The right quantity of coins and/or bills for %.2f$:\n", amount)
	for i := 0; i < len(denomSum); i++ {
		if denomSum[i] > 0 {
			fmt.Printf("\t%1.0f x $%.2f\t= $%.2f\n", denomSum[i], denominations[i], denomSum[i]*denominations[i])
		}
	}

}
