package exchangerates

import (
	"fmt"

	"github.com/fahmyabdul/self-growth/fetch-app/app"
)

func (p *ExchangeRates) CheckTableExists() error {
	sqliteConn := app.Properties.Databases.SQLiteConn
	query := fmt.Sprintf("SELECT * FROM %s", p.TableName())
	row, err := sqliteConn.Query(query)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

func (p *ExchangeRates) GetExchangeRates(from, to string) error {
	sqliteConn := app.Properties.Databases.SQLiteConn
	query := fmt.Sprintf("SELECT * FROM %s Where %s.'from' = '%s' AND %s.'to' = '%s' LIMIT 1", p.TableName(), p.TableName(), from, p.TableName(), to)
	row, err := sqliteConn.Query(query)
	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&p.ID, &p.From, &p.To, &p.Rates, &p.Limit, &p.LastFetch); err != nil {
			return err
		}
	}

	return nil
}
