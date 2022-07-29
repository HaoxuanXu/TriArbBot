package datamodel

// condition map records the coin pair and their corresponding spreads
type ConditionMap struct {
	Mapper    map[string]float64
	BaseCoins []string
}

type CoinPairs struct {
	Pairs []string
}

type CoinDependency struct {
	Dependency map[string][]string
}

func LoadCoinPairs() {

}
