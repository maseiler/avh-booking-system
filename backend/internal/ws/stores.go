package ws

import (
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
)

type AccountStore interface {
	Get(q *repo.Query) ([]models.Account, error)
	GetByID(id int) (*models.Account, error)
	Insert(a models.Account) (int, error)
	Update(a models.Account) (int, error)
}

type CategoryStore interface {
	Get(q *repo.Query) ([]models.Category, error)
}

type LocationStore interface {
	Get(q *repo.Query) ([]models.Location, error)
}

type ProductStore interface {
	Get(q *repo.Query) ([]models.Product, error)
}

type ProductGroupStore interface {
	Get(q *repo.Query) ([]models.ProductGroup, error)
}

type ProductVisibilityStore interface {
	Get(q *repo.Query) ([]models.ProductVisibility, error)
	Insert(v models.ProductVisibility) (int, error)
	Delete(id int) (int, error)
}

type UnitStore interface {
	Get(q *repo.Query) ([]models.Unit, error)
}

type VatStore interface {
	Get(q *repo.Query) ([]models.Vat, error)
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
