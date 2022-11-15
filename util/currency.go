package util

const (
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TRY:
		return true
	}
	return false
}

// not the best way but it was practical due to time constraints :/
func GetRate(FromCurrency string, ToCurrency string) float64 {
	rates := make(map[string]float64)

	rates[USD+TRY] = float64(18.61)
	rates[USD+EUR] = float64(0.96)
	rates[EUR+TRY] = float64(19.34)

	rates[TRY+USD] = float64(0.054)
	rates[EUR+USD] = float64(1.04)
	rates[TRY+EUR] = float64(0.052)

	return rates[FromCurrency+ToCurrency]
}

func Markup() float64 {
	return float64(0.05) // 5% cutoff
}
