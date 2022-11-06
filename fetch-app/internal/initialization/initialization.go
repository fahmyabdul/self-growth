package initialization

import (
	"fmt"

	"github.com/fahmyabdul/self-growth/fetch-app/internal/common"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models/endpoints_permission"
)

func InitBefore() error {
	// Generate Table Endpoints Permission
	modelPermission := endpoints_permission.EndpointsPermission{}
	err := modelPermission.GenerateTable()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Insert Default Permission to Table Endpoints Permission
	err = modelPermission.InsertDefaultData()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = common.CheckExchangeRates("InitBefore")
	if err != nil {
		return err
	}

	return nil
}

func InitAfter() error {
	return nil
}
