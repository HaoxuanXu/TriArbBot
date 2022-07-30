package transaction

import (
	"fmt"
	"math"
	"strings"

	"github.com/HaoxuanXu/TriArbBot/internal/broker"
	datamodel "github.com/HaoxuanXu/TriArbBot/strategy/data_model"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
)

func BuyBaseCoin(baseCoinSymbol string, brokerage *broker.AlpacaBroker, model *datamodel.Model) *alpaca.Order {

	qty := math.Min(brokerage.ENTRY_LIMIT, model.ConditionMap[baseCoinSymbol].MaxEntryCashAmount) / model.Quotes[baseCoinSymbol].AskPrice
	inputSymbol := fmt.Sprintf("%s/USD", baseCoinSymbol)

	order := brokerage.SubmitOrder(qty, inputSymbol, broker.BUY_SIDE, broker.MARKET_ORDER, broker.GTC)

	return order
}

func TradeForPairedCoin(pairedCoinSymbol string, baseCoinOrder *alpaca.Order, brokerage *broker.AlpacaBroker) *alpaca.Order {
	baseCoin := strings.Split(baseCoinOrder.Symbol, "/")[0]
	baseCoinQty := baseCoinOrder.FilledQty.InexactFloat64()
	inputSymbol := fmt.Sprintf("%s/%s", baseCoin, pairedCoinSymbol)

	order := brokerage.SubmitOrder(baseCoinQty, inputSymbol, broker.SELL_SIDE, broker.MARKET_ORDER, broker.GTC)

	return order
}

func SellPairedCoin(pairedCoinOrder *alpaca.Order, brokerage *broker.AlpacaBroker) *alpaca.Order {
	pairedCoinSymbol := pairedCoinOrder.Symbol
	pairedCoinQty := pairedCoinOrder.FilledQty.InexactFloat64()
	inputSymbol := fmt.Sprintf("%s/USD", pairedCoinSymbol)

	order := brokerage.SubmitOrder(pairedCoinQty, inputSymbol, broker.SELL_SIDE, broker.MARKET_ORDER, broker.GTC)

	return order
}
