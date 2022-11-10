package transactions

import (
	"strings"

	"github.com/fahmyabdul/self-growth/fetch-app/configs"
)

type Transactions struct {
	ID              string `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	TransactionHash string `json:"transaction_hash" bson:"transaction_hash"`
	Username        string `json:"username" bson:"username"`
	TransactionDate int    `json:"transaction_date" bson:"transaction_date"`
	PaidAmount      int    `json:"paid_amount" bson:"paid_amount"`
	PaymentMethod   int    `json:"payment_method" bson:"payment_method"`
	ErrCode         int    `json:"err_code,omitempty" bson:"err_code,omitempty" swaggerignore:"true"`
	ErrDescription  string `json:"err_description,omitempty" bson:"err_description,omitempty" swaggerignore:"true"`
	CreatedAt       int    `json:"created_at,omitempty" bson:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt       int    `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggerignore:"true"`
	PaidUSD         string `json:"paid_usd"`
}

type TransactionsAggregateContent struct {
	PaymentMethod int            `json:"payment_method"`
	Paid          AggregateGroup `json:"paid"`
}

type AggregateGroup struct {
	Collection []int   `json:"collection"`
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Median     int     `json:"median"`
	Avg        float64 `json:"avg"`
}

type RestTransactionResponse struct {
	Status string         `json:"status"`
	Code   int            `json:"code"`
	Data   []Transactions `json:"data"`
}

func (p *Transactions) TableName() string {
	tablename := "<schema>.t_transactions"
	return strings.ReplaceAll(tablename, "<schema>", configs.Properties.Databases.Postgre.Schema)
}

func (p *Transactions) KeyRedis() string {
	return "data:transactions"
}
