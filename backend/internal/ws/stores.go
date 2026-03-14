package ws

import (
	"context"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
)

type AccountStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Account, error)
	GetByID(ctx context.Context, id int) (*models.Account, error)
	Insert(ctx context.Context, a models.Account) (int, error)
	Update(ctx context.Context, a models.Account) (int, error)
}

type CategoryStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Category, error)
}

type LocationStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Location, error)
}

type ProductStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Product, error)
}

type ProductGroupStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ProductGroup, error)
}

type ProductVisibilityStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ProductVisibility, error)
	Insert(ctx context.Context, v models.ProductVisibility) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type UnitStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Unit, error)
}

type VatStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Vat, error)
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
