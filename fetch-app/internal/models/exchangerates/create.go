package exchangerates

import (
	"fmt"

	"github.com/fahmyabdul/self-growth/fetch-app/app"
)

func (p *ExchangeRates) Create(from, to string, rates float64, limit int, last_fetch int64) (int, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	lastInsertId := 0
	query := fmt.Sprintf(`
	INSERT INTO %s('from', 'to', 'rates', 'limit', 'last_fetch') 
	VALUES('%s', '%s', %f, %d, %d)
	RETURNING id
	`, p.TableName(), from, to, rates, limit, last_fetch)

	if err := sqliteConn.QueryRow(query).Scan(&lastInsertId); err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
