package messagebroker

// Publisher is
type Publisher interface {
	Publish(topic string, message []byte) error
}

// FunctionHandler is
type FunctionHandler func(message []byte) //

// ConsumerHandler is
type ConsumerHandler struct {
	Topic           string          //
	Channel         string          //
	FunctionHandler FunctionHandler //
}
