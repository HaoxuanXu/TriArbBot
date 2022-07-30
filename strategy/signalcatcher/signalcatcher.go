package signalcatcher

import datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"

func GetSignal(condition map[string]*datamodel.ConditionEntry) (string, float64) {

	maxSpread := 0.0
	winningPair := ""
	for coinPair, spreadAndEntry := range condition {
		if spreadAndEntry.SpreadPercent > maxSpread {
			maxSpread = spreadAndEntry.SpreadPercent
			winningPair = coinPair
		}
	}

	if maxSpread >= 0.01 {
		return winningPair, maxSpread
	}

	return "", maxSpread
}
