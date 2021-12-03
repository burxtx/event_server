package msg

// Message interface for messages containing payloads and headers
type Message interface {
	ID() string
	Headers() Headers
	Payload() []byte
}
