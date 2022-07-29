package broker

import (
	"sync"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
)

var lock = &sync.Mutex{}

type AlpacaBroker struct {
	client              alpaca.Client
	account             *alpaca.Account
	Clock               alpaca.Clock
	MaxPortfolioPercent float64
}

// GetBroker function creates an instance of the AlpacaBroker struct
func GetBroker(accountType, serverType string, entryPercent float64) *AlpacaBroker {
	lock.Lock()

	defer lock.Unlock()

	broker := &AlpacaBroker{}

	return broker
}

func (broker *AlpacaBroker) initialize(accountType, serverType string, entryPercent float64) {

}
