package currencies

import "testing"

var coins = map[string]string{
	"BCH": "Bitcoin Cash",
	"BTC": "Bitcoin",
	"ETH": "Ethereum",
	"LTC": "Litecoin",
}

func TestGetAllSymbols(t *testing.T) {
	t.Parallel()

	currencies := make(map[string]string, len(coins))
	for k, v := range coins {
		currencies[k] = v
	}

	for _, abbr := range GetAllSymbols() {
		delete(currencies, abbr)
	}

	if len(currencies) != 0 {
		t.Errorf("Expected currencies to be an empty map, but got %v", currencies)
	}
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
