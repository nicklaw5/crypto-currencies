package currencies

import "strings"

// A list of crypto coin symbols.
const (
	BCH = "BCH" // Bitcoin Cash
	BTC = "BTC" // Bitcoin
	ETH = "ETH" // Ethereum
	LTC = "LTC" // Litecoin
)

// GetAllSymbols returns a list of all currency symbols.
func GetAllSymbols() []string {
	return []string{
		BCH,
		BTC,
		ETH,
		LTC,
	}
}

// GetAllCurrencies ...
func GetAllCurrencies() map[string]string {
	return map[string]string{
		BCH: "Bitcoin Cash",
		BTC: "Bitcoin",
		ETH: "Ethereum",
		LTC: "Litecoin",
	}
}

// GetCurrencyBySymbol returns the currency for the provided symbol,
// If no currency matches a symbol, an empty string is returned.
func GetCurrencyBySymbol(symbol string) string {
	for abbr, currency := range GetAllCurrencies() {
		if abbr == strings.ToUpper(symbol) {
			return currency
		}
	}

	return ""
}

// GetSymbolByCurrency returns the symbol for the provided currency,
// If no currency matches a symbol, an empty string is returned.
func GetSymbolByCurrency(currency string) string {
	for abbr, curr := range GetAllCurrencies() {
		if strings.ToLower(curr) == strings.ToLower(currency) {
			return abbr
		}
	}

	return ""
}
