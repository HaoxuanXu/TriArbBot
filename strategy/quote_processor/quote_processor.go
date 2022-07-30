package quoteprocessor

import (
	"fmt"
	"log"

	"github.com/HaoxuanXu/TriArbBot/internal/dataEngine"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func ProcessCryptoQuotes(
	model *datamodel.Model,
	engine *dataEngine.MarketDataEngine,
	coinDependency map[string][]string) {
	cryptoQuotes := engine.GetLatestCryptoQuotes(model.CoinPairs)

	parseQuotes(model, cryptoQuotes, coinDependency)
}

func parseQuotes(model *datamodel.Model, quotes map[string]marketdata.CryptoQuote, coinDependency map[string][]string) {
	for baseCoin, pairedCoinList := range coinDependency {
		for _, pairedCoin := range pairedCoinList {
			coinPair := fmt.Sprintf("%s%s", baseCoin, pairedCoin)

			// get the respective quotes
			baseCoinQuote := quotes[fmt.Sprintf("%sUSD", baseCoin)]
			pairedCoinQuote := quotes[fmt.Sprintf("%sUSD", pairedCoin)]
			coinPairQuote := quotes[coinPair]

			spreadPercent, maxEntryCash := calculateSpreadPercentAndEntryAmount(&baseCoinQuote, &pairedCoinQuote, &coinPairQuote)

			if _, ok := model.ConditionMap[coinPair]; !ok {
				model.ConditionMap[coinPair] = &datamodel.ConditionEntry{}
			}
			model.ConditionMap[coinPair].SpreadPercent = spreadPercent
			model.ConditionMap[coinPair].MaxEntryCashAmount = maxEntryCash

			model.Quotes[fmt.Sprintf("%s/USD", baseCoin)] = &baseCoinQuote
			model.Quotes[fmt.Sprintf("%s/USD", pairedCoin)] = &pairedCoinQuote
			model.Quotes[fmt.Sprintf("%s/%s", baseCoin, pairedCoin)] = &coinPairQuote

		}
	}
}

func calculateSpreadPercentAndEntryAmount(baseCoinQuote, pairedCoinQuote, coinPairQuote *marketdata.CryptoQuote) (float64, float64) {

	// enter by buying baseCoin
	var pairedCoinQty float64
	baseCoinQty := baseCoinQuote.AskSize
	// log.Printf("baseCoinQty: %f\n", baseCoinQty)
	// sell baseCoin for pairedCoin
	if pairedCoinQty = (baseCoinQuote.AskSize / coinPairQuote.BidPrice); pairedCoinQty > coinPairQuote.AskSize {

		// reduce entry amount so that we can buy avoid getting into level 2 for pairedCoin
		baseCoinQty *= (coinPairQuote.AskSize / pairedCoinQty)
		pairedCoinQty = coinPairQuote.AskSize
	}
	// log.Printf("pairedCoinQty: %f\n", pairedCoinQty)
	// sell pairedCoin for USD
	if pairedCoinQty > pairedCoinQuote.BidSize {
		baseCoinQty *= (pairedCoinQuote.BidSize / pairedCoinQty)
		pairedCoinQty = pairedCoinQuote.BidSize
	}

	entryCash := baseCoinQty * baseCoinQuote.AskPrice
	exitCash := pairedCoinQty * pairedCoinQuote.BidPrice
	spreadPercent := (exitCash - entryCash) / entryCash
	log.Println(spreadPercent)

	return spreadPercent, entryCash

}
