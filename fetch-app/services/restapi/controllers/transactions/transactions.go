package transactions

type CtrlTransactions struct{}

func (p *CtrlTransactions) KafkaTopic() string {
	return "t-transactions"
}

func (p *CtrlTransactions) KafkaEventName() string {
	return "transactions"
}
