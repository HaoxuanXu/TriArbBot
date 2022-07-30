package pipeline

import (
	"log"

	"github.com/HaoxuanXu/TriArbBot/internal/broker"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/HaoxuanXu/TriArbBot/strategy/transaction"
)

func ExecuteTransactionPipeline(baseCoinSymbol, pairedCoinSymbol string, brokerage *broker.AlpacaBroker, model *datamodel.Model) float64 {

	var entryCashAmount float64
	var exitCashAmount float64

	// buy base coin as the first transaction
	firstOrder := transaction.BuyBaseCoin(baseCoinSymbol, brokerage, model)
	entryCashAmount = firstOrder.FilledQty.InexactFloat64() * firstOrder.FilledAvgPrice.InexactFloat64()

	// print the result of the first transaction
	log.Printf("Bought %f %s with $%.2f.\n", firstOrder.FilledQty.InexactFloat64(), baseCoinSymbol, entryCashAmount)

	// sell the base coin for the paired coin as the second transaction
	secondOrder := transaction.TradeForPairedCoin(pairedCoinSymbol, firstOrder, brokerage)

	// print the result of the second transaction
	log.Printf("Sold %f of %s for %f of %s.\n",
		secondOrder.FilledQty.InexactFloat64(),
		baseCoinSymbol,
		secondOrder.FilledQty.InexactFloat64()*secondOrder.FilledAvgPrice.InexactFloat64(),
		pairedCoinSymbol)

	finalOrder := transaction.SellPairedCoin(secondOrder, brokerage)
	exitCashAmount = finalOrder.FilledQty.InexactFloat64() * finalOrder.FilledAvgPrice.InexactFloat64()

	log.Printf("Earned $%.2f in this transaction.\n", exitCashAmount-entryCashAmount)

	return finalOrder.FilledQty.InexactFloat64() * finalOrder.FilledAvgPrice.InexactFloat64()
}
