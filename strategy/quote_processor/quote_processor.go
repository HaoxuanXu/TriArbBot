package quoteprocessor

import (
	"fmt"

	"github.com/HaoxuanXu/TriArbBot/internal/dataEngine"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func ProcessCryptoQuotes(
	model *datamodel.Model,
	engine *dataEngine.MarketDataEngine) {
	cryptoQuotes := engine.GetLatestCryptoQuotes(model.CoinPairs.Pairs)

	parseQuotes(model, cryptoQuotes)
}

func parseQuotes(model *datamodel.Model, quotes map[string]marketdata.CryptoXBBO) {
	for baseCoin, pairedCoinList := range model.CoinDependency.Dependency {
		for _, pairedCoin := range pairedCoinList {
			coinPair := fmt.Sprintf("%s/%s", baseCoin, pairedCoin)

			// get the respective quotes
			baseCoinQuote := quotes[fmt.Sprintf("%s/USD", baseCoin)]
			pairedCoinQuote := quotes[fmt.Sprintf("%s/USD", pairedCoin)]
			coinPairQuote := quotes[coinPair]

			spreadPercent, maxEntryCash := calculateSpreadPercentAndEntryAmount(&baseCoinQuote, &pairedCoinQuote, &coinPairQuote)

			model.ConditionMap.Mapper[coinPair] = []float64{spreadPercent, maxEntryCash}

			model.CoinPairs.Quotes[baseCoin] = &baseCoinQuote
			model.CoinPairs.Quotes[pairedCoin] = &pairedCoinQuote
			model.CoinPairs.Quotes[coinPair] = &coinPairQuote

		}
	}
}

func calculateSpreadPercentAndEntryAmount(baseCoinQuote, pairedCoinQuote, coinPairQuote *marketdata.CryptoXBBO) (float64, float64) {

	// enter by buying baseCoin
	var pairedCoinQty float64
	baseCoinQty := baseCoinQuote.AskSize

	// sell baseCoin for pairedCoin
	if pairedCoinQty = baseCoinQuote.AskSize * coinPairQuote.BidPrice; pairedCoinQty > coinPairQuote.BidSize {

		// reduce entry amount so that we can buy avoid getting into level 2 for pairedCoin
		baseCoinQty *= (coinPairQuote.BidSize / pairedCoinQty)
		pairedCoinQty = coinPairQuote.BidSize
	}

	// sell pairedCoin for USD
	if pairedCoinQty > pairedCoinQuote.BidSize {
		baseCoinQty *= (pairedCoinQuote.BidSize / pairedCoinQty)
		pairedCoinQty = pairedCoinQuote.BidSize
	}

	entryCash := baseCoinQty * baseCoinQuote.AskPrice
	exitCash := pairedCoinQty * pairedCoinQuote.BidPrice
	spreadPercent := (exitCash - entryCash) / entryCash

	return spreadPercent, entryCash

}
