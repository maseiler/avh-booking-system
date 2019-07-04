package data

// ItemFromTo is needed to unmarshal an item and two strings from API
type ItemFromTo struct {
	Item Item
	From string
	To   string
}
