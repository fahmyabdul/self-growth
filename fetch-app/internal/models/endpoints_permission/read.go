package endpoints_permission

import (
	"fmt"

	"github.com/fahmyabdul/self-growth/fetch-app/app"
)

func (p *EndpointsPermission) GetListPermission() (map[string]string, error) {
	sqliteConn := app.Properties.Databases.SQLiteConn
	query := fmt.Sprintf(`SELECT * FROM %s`, p.TableName())
	row, err := sqliteConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	listPermissions := make(map[string]string)
	for row.Next() {
		var current_data EndpointsPermission
		if err = row.Scan(&current_data.ID, &current_data.Endpoint, &current_data.Permission); err != nil {
			return nil, err
		}

		listPermissions[current_data.Endpoint] = current_data.Permission
	}

	return listPermissions, nil
}
