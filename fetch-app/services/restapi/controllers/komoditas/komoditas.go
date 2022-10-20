package komoditas

type CtrlKomoditas struct{}

func (p *CtrlKomoditas) KafkaTopic() string {
	return "t-komoditas"
}

func (p *CtrlKomoditas) KafkaEventName() string {
	return "komoditas"
}
