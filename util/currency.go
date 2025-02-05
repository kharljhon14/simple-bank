package util

const (
	USD = "USD"
	PHP = "PHP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, PHP:
		return true
	default:
		return false
	}
}
