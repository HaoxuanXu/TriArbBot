package datamodel

import (
	"fmt"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func GetModel(assets config.Assets) *Model {
	model := Model{}
	model.LoadCoinPairs(assets)

	return &model
}

type Model struct {
	ConditionMap   *ConditionMap
	CoinPairs      *CoinPairs
	CoinDependency *CoinDependency
}

// condition map records the coin pair and their corresponding spreads
type ConditionMap struct {
	Mapper map[string][]float64 // BTC: [spreadPercent, maxEntryCashAmount]
}

type CoinPairs struct {
	Pairs  []string
	Quotes map[string]*marketdata.CryptoXBBO
}

type CoinDependency struct {
	Dependency map[string][]string
}

func (model *Model) LoadCoinPairs(assets config.Assets) {
	// populate CoinDependency
	model.CoinDependency.Dependency = assets.Coins

	// populate CoinPairs
	for baseCoin := range model.CoinDependency.Dependency {
		// populate base coin input symbols
		baseCoinUSDSymbol := fmt.Sprintf("%s/USD", baseCoin)
		model.CoinPairs.Pairs = append(model.CoinPairs.Pairs, baseCoinUSDSymbol)
	}

	for baseCoin, pairedCoinList := range model.CoinDependency.Dependency {
		for _, pairedCoin := range pairedCoinList {
			pairedCoinUSDSymbol := fmt.Sprintf("%s/USD", pairedCoin)
			pairedCoinBaseCoinSymbol := fmt.Sprintf("%s/%s", pairedCoin, baseCoin)
			model.CoinPairs.Pairs = append(model.CoinPairs.Pairs, pairedCoinUSDSymbol)
			model.CoinPairs.Pairs = append(model.CoinPairs.Pairs, pairedCoinBaseCoinSymbol)
		}
	}
}
