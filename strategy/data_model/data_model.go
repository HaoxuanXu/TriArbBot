package datamodel

// condition map records the coin pair and their corresponding spreads
type ConditionMap struct {
	Mapper map[string]float64
}

type CoinPairs struct {
	Pairs []string
}

type CoinDependency struct {
	Dependency map[string][]string
}

func LoadCoinPairs() {

}
