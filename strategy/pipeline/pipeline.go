package pipeline

import (
	"fmt"

	"github.com/HaoxuanXu/TriArbBot/internal/broker"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/HaoxuanXu/TriArbBot/strategy/transaction"
)

func ExecuteTransactionPipeline(baseCoinSymbol, pairedCoinSymbol string, brokerage *broker.AlpacaBroker, model *datamodel.Model) float64 {

	var entryCashAmount float64
	var exitCashAmount float64

	firstOrder := transaction.BuyBaseCoin(baseCoinSymbol, brokerage, model)
	entryCashAmount = firstOrder.FilledQty.InexactFloat64() * firstOrder.FilledAvgPrice.InexactFloat64()

	secondOrder := transaction.TradeForPairedCoin(pairedCoinSymbol, firstOrder, brokerage)

	finalOrder := transaction.SellPairedCoin(secondOrder, brokerage)
	exitCashAmount = finalOrder.FilledQty.InexactFloat64() * finalOrder.FilledAvgPrice.InexactFloat64()

	fmt.Printf("Earned $%.2f in this transaction.\n", exitCashAmount-entryCashAmount)

	return finalOrder.FilledQty.InexactFloat64() * finalOrder.FilledAvgPrice.InexactFloat64()
}
