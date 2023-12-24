package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	rate := 2.475
	switch {
	case balance < 0.0:
		rate = 3.213
	case balance < 1000.0:
		rate = .5
	case balance < 5000.0:
		rate = 1.621
	}
	return float32(rate)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	count := 0
	for balance < targetBalance {
		count++
		balance = AnnualBalanceUpdate(balance)
	}
	return count
}
