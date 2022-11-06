package endpoints_permission

import (
	"fmt"

	"github.com/fahmyabdul/self-growth/fetch-app/app"
)

type EndpointsPermission struct {
	ID         int    `json:"id"`
	Endpoint   string `json:"endpoint"`
	Permission string `json:"permission"`
}

func (p *EndpointsPermission) TableName() string {
	return "t_endpoints_permission"
}

func (p *EndpointsPermission) GenerateTable() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			endpoint TEXT,
			permission TEXT,
			UNIQUE(endpoint)
		)
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	return nil
}

func (p *EndpointsPermission) InsertDefaultData() error {
	sqliteConn := app.Properties.Databases.SQLiteConn

	query := fmt.Sprintf(`
		INSERT OR IGNORE INTO %s (endpoint, permission) 
			VALUES('/komoditas/get', '*'), ('/komoditas/get/aggregate', 'admin')
	`, p.TableName())

	if _, err := sqliteConn.Exec(query); err != nil {
		return err
	}

	return nil
}
