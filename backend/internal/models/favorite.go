package models

// Favorite represents a product that an account has marked or frequently ordered.
// Count tracks how many times the product has been ordered by the account.
type Favorite struct {
	ProductID int `json:"productId" db:"product"`
	AccountID int `json:"accountId" db:"account"`
	Count     int `json:"count" db:"count"`
}

// CreateFavorite creates a new Favorite entry for the given account and product.
func CreateFavorite(productID, accountID, count int) Favorite {
	return Favorite{
		ProductID: productID,
		AccountID: accountID,
		Count:     count,
	}
}
