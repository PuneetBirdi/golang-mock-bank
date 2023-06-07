package util

const (
	USD = "USD"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
		case USD, CAD:
		return true
	}
	return false
}
