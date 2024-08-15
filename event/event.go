package event

// Event is given to a subscription handler for processing.
type Event interface {
	// Topic return the topic of the message
	Topic() string
	// Message return the message body
	Message() *Message
}

// Message is the message entity.
type Message struct {
	key    string
	Header map[string]string
	Body   []byte
}

// SetMessageKey set a flag that represents the message
func (msg *Message) SetMessageKey(key string) {
	msg.key = key
}

// MessageKey get the flag that represents the message
func (msg Message) MessageKey() string {
	return msg.key
}

func (msg Message) Topic() string {
	return msg.MessageKey()
}

func (msg Message) Message() *Message {
	return &msg
}
