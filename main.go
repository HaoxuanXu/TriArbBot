package main

import (
	"flag"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/HaoxuanXu/TriArbBot/strategy"
)

func main() {
	accountType := flag.String("accounttype", "paper", "This determines if we will run the job on paper or live accounts")
	serverType := flag.String("servertype", "production", "This determines if we are using the production or staging brokerage account")
	entryPercent := flag.Float64("entrypercent", 0.12, "this is the percent of portfolio value we will commit")

	flag.Parse()

	creds := config.GetCredentials(*accountType, *serverType)
	assets := config.GetAssets()

	strategy.RunTriangularArbitrage(creds, assets, *entryPercent)
}
