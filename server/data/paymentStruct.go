package data

// Payment
type Payment struct {
	User          User
	Balance       float32
	PaymentMethod string
	IntentID      string
	CardReader    string
}
