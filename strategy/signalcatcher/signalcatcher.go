package signalcatcher

import datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"

func GetSignal(condition *datamodel.ConditionMap) string {

	maxSpread := 0.0
	winningPair := ""
	for coinPair, spreadAndEntry := range condition.Mapper {
		if spreadAndEntry[0] > maxSpread {
			maxSpread = spreadAndEntry[0]
			winningPair = coinPair
		}
	}

	if maxSpread >= 0.01 {
		return winningPair
	}

	return ""
}
