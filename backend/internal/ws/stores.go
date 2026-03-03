package ws

import "github.com/av-huette/avh-booking-system/internal/models"

type AccountStore interface {
	Get(q *models.Query) ([]models.Account, error)
	GetById(id int) (*models.Account, error)
	Insert(a models.Account) (int, error)
	Update(a models.Account) (int, error)
}

type CategoryStore interface {
	Get(q *models.Query) ([]models.Category, error)
}

type LocationStore interface {
	Get(q *models.Query) ([]models.Location, error)
}

type ProductStore interface {
	Get(q *models.Query) ([]models.Product, error)
}

type ProductGroupStore interface {
	Get(q *models.Query) ([]models.ProductGroup, error)
}

type ProductVisibilityStore interface {
	Get(q *models.Query) ([]models.ProductVisibility, error)
}

type UnitStore interface {
	Get(q *models.Query) ([]models.Unit, error)
}

type VatStore interface {
	Get(q *models.Query) ([]models.Vat, error)
}

type Stores struct {
	Account           AccountStore
	Category          CategoryStore
	Location          LocationStore
	Product           ProductStore
	ProductGroup      ProductGroupStore
	ProductVisibility ProductVisibilityStore
	Unit              UnitStore
	Vat               VatStore
}
