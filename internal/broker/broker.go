package broker

import (
	"sync"
	"time"

	"github.com/HaoxuanXu/TriArbBot/config"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/shopspring/decimal"
)

var lock = &sync.Mutex{}

type AlpacaBroker struct {
	client              alpaca.Client
	account             *alpaca.Account
	Clock               alpaca.Clock
	MaxPortfolioPercent float64
	Cash                float64
	ENTRY_LIMIT         float64
}

// GetBroker function creates an instance of the AlpacaBroker struct
func GetBroker(accountType, serverType string, entryPercent float64) *AlpacaBroker {
	lock.Lock()

	defer lock.Unlock()

	broker := &AlpacaBroker{}
	broker.initialize(accountType, serverType, entryPercent)

	return broker
}

func (broker *AlpacaBroker) initialize(accountType, serverType string, entryPercent float64) {
	creds := config.GetCredentials(accountType, serverType)
	broker.client = alpaca.NewClient(
		alpaca.ClientOpts{
			ApiKey:    creds.API_KEY,
			ApiSecret: creds.API_SECRET,
			BaseURL:   creds.BASE_URL,
		},
	)
	account, _ := broker.client.GetAccount()
	clock, _ := broker.client.GetClock()
	broker.account = account
	broker.Clock = *clock
	broker.Cash = broker.account.Cash.InexactFloat64()
	broker.MaxPortfolioPercent = entryPercent
	broker.ENTRY_LIMIT = broker.Cash * broker.MaxPortfolioPercent
}

func (broker *AlpacaBroker) refreshOrderStatus(orderID string) (string, *alpaca.Order) {
	newOrder, _ := broker.client.GetOrder(orderID)
	status := newOrder.Status
	return status, newOrder
}

func (broker *AlpacaBroker) MonitorOrder(order *alpaca.Order) (*alpaca.Order, bool) {
	success := false
	orderID := order.ID
	status, updatedOrder := broker.refreshOrderStatus(orderID)
	for !success {
		switch status {
		case "new", "accepted", "partially_filled":
			time.Sleep(100 * time.Millisecond)
			status, updatedOrder = broker.refreshOrderStatus(orderID)
		case "filled":
			success = true
		case "done_for_day", "canceled", "expired", "replaced":
			success = false
		default:
			time.Sleep(100 * time.Millisecond)
			status, updatedOrder = broker.refreshOrderStatus(orderID)
		}
	}
	return updatedOrder, success
}

func (broker *AlpacaBroker) SubmitOrder(qty float64, symbol, side, orderType, timeInForce string) *alpaca.Order {
	quantity := decimal.NewFromFloat(qty)
	order, _ := broker.client.PlaceOrder(
		alpaca.PlaceOrderRequest{
			AssetKey:    &symbol,
			AccountID:   broker.account.ID,
			Qty:         &quantity,
			Side:        alpaca.Side(side),
			Type:        alpaca.OrderType(orderType),
			TimeInForce: alpaca.TimeInForce(timeInForce),
		},
	)

	finalOrder, _ := broker.MonitorOrder(order)
	return finalOrder
}

func (broker *AlpacaBroker) CloseAllPositions() {
	broker.client.CloseAllPositions()
}
