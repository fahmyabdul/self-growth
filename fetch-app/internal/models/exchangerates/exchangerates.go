package exchangerates

import (
	"fmt"

	"github.com/fahmyabdul/efishery-task/fetch-app/app"
)

type ExchangeRates struct {
	ID        string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Rates     float64 `json:"rates"`
	Limit     int     `json:"limit"`
	LastFetch int64   `json:"last_fetch"`
}

// Used as Exchange Rates data cache
// Better implementation will be to use Redis or any nosql storage db as cache
var GlobExchangeRatesData ExchangeRates

func (p *ExchangeRates) TableName() string {
	return "t_exchangerates"
}

func (p *ExchangeRates) GenerateTable() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY AUTOINCREMENT, 'from' TEXT, 'to' TEXT, 'rates' REAL, 'limit' INTEGER, 'last_fetch' INTEGER)
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	return nil
}
