package strategy

import (
	"log"
	"strings"
	"time"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/HaoxuanXu/TriArbBot/internal/broker"
	"github.com/HaoxuanXu/TriArbBot/internal/dataEngine"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/HaoxuanXu/TriArbBot/strategy/pipeline"
	quoteprocessor "github.com/HaoxuanXu/TriArbBot/strategy/quote_processor"
	"github.com/HaoxuanXu/TriArbBot/strategy/signalcatcher"
)

func RunTriangularArbitrage(creds config.Credentials, assets config.Assets, entryPercent float64) {

	log.Println("Initializing strategy ...")
	model := datamodel.GetModel(assets)
	marketDataEngine := dataEngine.GetDataEngine(creds.ACCOUNT_TYPE, creds.SERVER_TYPE)
	brokerage := broker.GetBroker(creds.ACCOUNT_TYPE, creds.SERVER_TYPE, entryPercent)

	log.Println("Begin strategy")
	for {

		// get and process all crypto quotes
		quoteprocessor.ProcessCryptoQuotes(model, marketDataEngine, assets.Coins)

		winningPair, maxSpread := signalcatcher.GetSignal(model.ConditionMap)

		log.Printf("Max Spread is %.2f%% for %s.\n", maxSpread, winningPair)
		pairs := strings.Split(winningPair, "/")

		if len(pairs) == 2 {
			baseCoin := pairs[0]
			pairedCoin := pairs[1]

			pipeline.ExecuteTransactionPipeline(baseCoin, pairedCoin, brokerage, model)
		}

		// pause for 1 second after each iteration
		time.Sleep(10 * time.Second)
	}
}
