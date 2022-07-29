package signalcatcher

import datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"

func GetSignal(condition *datamodel.ConditionMap) string {

	maxSpread := 0.0
	winningPair := ""
	for coinPair, spread := range condition.Mapper {
		if spread > maxSpread {
			maxSpread = spread
			winningPair = coinPair
		}
	}

	if maxSpread >= 0.01 {
		return winningPair
	}

	return ""
}
