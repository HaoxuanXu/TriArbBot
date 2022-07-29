package quoteprocessor

import (
	"fmt"

	"github.com/HaoxuanXu/TriArbBot/internal/dataengine"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func ProcessCryptoQuotes(
	condition *datamodel.ConditionMap,
	dependency *datamodel.CoinDependency,
	pairs *datamodel.CoinPairs,
	engine *dataengine.MarketDataEngine) {
	cryptoQuotes := engine.GetLatestCryptoQuotes(pairs.Pairs)

	parseQuotes(condition, dependency, cryptoQuotes)
}

func parseQuotes(condition *datamodel.ConditionMap, dependency *datamodel.CoinDependency, quotes map[string]marketdata.CryptoQuote) {
	for baseCoin, pairedCoinList := range dependency.Dependency {
		for _, pairedCoin := range pairedCoinList {
			coinPair := fmt.Sprintf("%s/%s", pairedCoin, baseCoin)

			// get the respective quotes
			baseCoinQuote := quotes[fmt.Sprintf("%s/USD", baseCoin)]
			pairedCoinQuote := quotes[fmt.Sprintf("%s/USD", pairedCoin)]
			coinPairQuote := quotes[coinPair]

			// We assume we are committing $10000
			finalReturn := (baseCoinQuote.AskPrice * coinPairQuote.AskPrice) / pairedCoinQuote.BidPrice
			spreadPercent := (finalReturn - baseCoinQuote.AskPrice) / baseCoinQuote.AskPrice

			condition.Mapper[coinPair] = spreadPercent
		}
	}
}
