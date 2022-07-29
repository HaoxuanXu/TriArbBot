package dataengine

import (
	"log"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

type MarketDataEngine struct {
	client marketdata.Client
}

func GetDataEngine(accountType, serverType string) *MarketDataEngine {
	engine := &MarketDataEngine{}
	engine.initialize(accountType, serverType)
	return engine
}

func (engine *MarketDataEngine) initialize(accountType, serverType string) {
	creds := config.GetCredentials(accountType, serverType)
	engine.client = marketdata.NewClient(
		marketdata.ClientOpts{
			ApiKey:    creds.API_KEY,
			ApiSecret: creds.API_SECRET,
		},
	)
}

func (engine *MarketDataEngine) GetLatestCryptoQuotes(symbols []string) map[string]marketdata.CryptoQuote {

	quotes, err := engine.client.GetLatestCryptoQuotes(symbols, "FTXU")
	if err != nil {
		log.Printf("error occurred when getting latest quotes: %s\n", err)
	}

	return quotes
}
