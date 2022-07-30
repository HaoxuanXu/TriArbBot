package datamodel

import (
	"fmt"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func GetModel(assets config.Assets) *Model {
	model := Model{}
	model.ConditionMap = make(map[string]*ConditionEntry)
	model.Quotes = make(map[string]*marketdata.CryptoQuote)
	LoadCoinPairs(&model, assets)

	return &model
}

// condition map records the coin pair and their corresponding spreads
type Model struct {
	ConditionMap map[string]*ConditionEntry
	CoinPairs    []string
	Quotes       map[string]*marketdata.CryptoQuote
}

type ConditionEntry struct {
	SpreadPercent      float64
	MaxEntryCashAmount float64
}

var CoinDependency = make(map[string][]string)

func LoadCoinPairs(model *Model, assets config.Assets) {

	// populate CoinDependency
	CoinDependency = assets.Coins

	// populate CoinPairs
	for baseCoin := range CoinDependency {
		// populate base coin input symbols
		baseCoinUSDSymbol := fmt.Sprintf("%sUSD", baseCoin)
		model.CoinPairs = append(model.CoinPairs, baseCoinUSDSymbol)
	}

	for baseCoin, pairedCoinList := range CoinDependency {
		for _, pairedCoin := range pairedCoinList {
			pairedCoinUSDSymbol := fmt.Sprintf("%sUSD", pairedCoin)
			pairedCoinBaseCoinSymbol := fmt.Sprintf("%s%s", pairedCoin, baseCoin)
			model.CoinPairs = append(model.CoinPairs, pairedCoinUSDSymbol)
			model.CoinPairs = append(model.CoinPairs, pairedCoinBaseCoinSymbol)
		}
	}
}
