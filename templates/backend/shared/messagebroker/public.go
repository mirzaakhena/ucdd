package messagebroker

// NewProducer is
func NewProducer(url string) Publisher {
	return &producer{
		defaultProducer: nsqProducerInit(url),
	}
}

// RunConsumer is
func RunConsumer(blocking bool, url string, consumers []ConsumerHandler) {
	nsqConsumerInit(blocking, url, consumers)
}
