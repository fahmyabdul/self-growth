package transactions

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transactions struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	TransactionHash string             `json:"transaction_hash" bson:"transaction_hash"`
	Username        string             `json:"username" bson:"username"`
	TransactionDate int                `json:"transaction_date" bson:"transaction_date"`
	PaidAmount      int                `json:"paid_amount" bson:"paid_amount"`
	PaymentMethod   int                `json:"payment_method" bson:"payment_method"`
	ErrCode         int                `json:"err_code,omitempty" bson:"err_code,omitempty" swaggerignore:"true"`
	ErrDescription  string             `json:"err_description,omitempty" bson:"err_description,omitempty" swaggerignore:"true"`
	CreatedAt       int                `json:"created_at,omitempty" bson:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt       int                `json:"updated_at,omitempty" bson:"updated_at,omitempty" swaggerignore:"true"`
}

func (p *Transactions) CollectionName() string {
	return "t_transactions"
}
