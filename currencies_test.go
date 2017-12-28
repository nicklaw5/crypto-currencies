package currencies

import "testing"

var coins = map[string]string{
	"BCH": "Bitcoin Cash",
	"BTC": "Bitcoin",
	"ETH": "Ethereum",
	"LTC": "Litecoin",
}

func TestGetCurrencyBySymbol(t *testing.T) {
	t.Parallel()

	for abbr, coin := range coins {
		currency := GetCurrencyBySymbol(abbr)
		if currency != coin {
			t.Errorf("Expected currency to be \"%s\", but got \"%s\"", coin, currency)
		}
	}

	// Test on failure to find symbol
	currency := GetCurrencyBySymbol("bogus")
	if currency != "" {
		t.Errorf("Expected currency to be \"%s\", but got \"%s\"", "", currency)
	}
}

func TestGetSymbolCurrency(t *testing.T) {
	t.Parallel()

	for abbr, coin := range coins {
		symbol := GetSymbolByCurrency(coin)
		if symbol != abbr {
			t.Errorf("Expected symbol to be \"%s\", but got \"%s\"", abbr, symbol)
		}
	}

	// Test on failure to find currency
	symbol := GetSymbolByCurrency("bogus")
	if symbol != "" {
		t.Errorf("Expected symbol to be \"%s\", but got \"%s\"", "", symbol)
	}
}
