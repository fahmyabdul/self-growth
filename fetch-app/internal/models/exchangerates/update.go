package exchangerates

import (
	"fmt"

	"github.com/fahmyabdul/efishery-task/fetch-app/app"
)

func (p *ExchangeRates) Update(from, to string, rates float64, limit int, last_fetch int64) (int, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn

	lastInsertId := 0
	query := fmt.Sprintf(`
		UPDATE %s
		SET 
			'from' = '%s', 
			'to' = '%s',
			'rates' = %f, 
			'limit' = %d, 
			'last_fetch' = %d
		WHERE %s.'from' = '%s' AND %s.'to' = '%s'
		RETURNING id;
	`, p.TableName(), from, to, rates, limit, last_fetch, p.TableName(), from, p.TableName(), to)
	if err := sqliteConn.QueryRow(query).Scan(&lastInsertId); err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
