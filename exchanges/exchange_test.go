package exchange

import (
	"testing"

	"github.com/thrasher-/gocryptotrader/currency/pair"

	"github.com/thrasher-/gocryptotrader/config"
)

func TestGetName(t *testing.T) {
	GetName := Base{
		Name: "TESTNAME",
	}

	name := GetName.GetName()
	if name != "TESTNAME" {
		t.Error("Test Failed - Exchange getName() returned incorrect name")
	}
}

func TestGetEnabledCurrencies(t *testing.T) {
	enabledPairs := []string{"BTCUSD", "BTCAUD", "LTCUSD", "LTCAUD"}
	GetEnabledCurrencies := Base{
		Name:         "TESTNAME",
		EnabledPairs: enabledPairs,
	}

	enCurr := GetEnabledCurrencies.GetEnabledCurrencies()
	if enCurr[0] != "BTCUSD" {
		t.Error("Test Failed - Exchange GetEnabledCurrencies() incorrect string")
	}
}

func TestGetAvailableCurrencies(t *testing.T) {
	availablePairs := []string{"BTCUSD", "BTCAUD", "LTCUSD", "LTCAUD"}
	GetEnabledCurrencies := Base{
		Name:           "TESTNAME",
		AvailablePairs: availablePairs,
	}

	enCurr := GetEnabledCurrencies.GetAvailableCurrencies()
	if enCurr[0] != "BTCUSD" {
		t.Error("Test Failed - Exchange GetAvailableCurrencies() incorrect string")
	}
}

func TestFormatCurrency(t *testing.T) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig(config.ConfigTestFile)
	if err != nil {
		t.Fatalf("Failed to load config file. Error: %s", err)
	}

	currency := pair.NewCurrencyPair("btc", "usd")
	expected := "BTC-USD"
	actual := FormatCurrency(currency).String()
	if actual != expected {
		t.Errorf("Test failed - Exchange TestFormatCurrency %s != %s",
			actual, expected)
	}
}

func TestSetEnabled(t *testing.T) {
	SetEnabled := Base{
		Name:    "TESTNAME",
		Enabled: false,
	}

	SetEnabled.SetEnabled(true)
	if !SetEnabled.Enabled {
		t.Error("Test Failed - Exchange SetEnabled(true) did not set boolean")
	}
}

func TestIsEnabled(t *testing.T) {
	IsEnabled := Base{
		Name:    "TESTNAME",
		Enabled: false,
	}

	if IsEnabled.IsEnabled() {
		t.Error("Test Failed - Exchange IsEnabled() did not return correct boolean")
	}
}

func TestSetAPIKeys(t *testing.T) {
	SetAPIKeys := Base{
		Name:    "TESTNAME",
		Enabled: false,
	}

	SetAPIKeys.SetAPIKeys("RocketMan", "Digereedoo", "007", false)
	if SetAPIKeys.APIKey != "RocketMan" && SetAPIKeys.APISecret != "Digereedoo" && SetAPIKeys.ClientID != "007" {
		t.Error("Test Failed - Exchange SetAPIKeys() did not set correct values")
	}
	SetAPIKeys.SetAPIKeys("RocketMan", "Digereedoo", "007", true)
}

func TestUpdateAvailableCurrencies(t *testing.T) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig(config.ConfigTestFile)
	UAC := Base{Name: "ANX"}
	exchangeProducts := []string{"ltc", "btc", "usd", "aud"}

	if err != nil {
		t.Error(
			"Test Failed - Exchange UpdateAvailableCurrencies() did not set correct values",
		)
	}
	err2 := UAC.UpdateAvailableCurrencies(exchangeProducts)
	if err2 != nil {
		t.Errorf("Test Failed - Exchange UpdateAvailableCurrencies() error: %s", err2)
	}
}
