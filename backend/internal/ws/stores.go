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
	GetByID(ctx context.Context, id int) (*models.Category, error)
	Insert(ctx context.Context, c models.Category) (int, error)
	Update(ctx context.Context, c models.Category) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type FavoritesStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Favorite, error)
	Insert(ctx context.Context, f models.Favorite) error
	Update(ctx context.Context, f models.Favorite) error
}

type LocationStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Location, error)
	GetByID(ctx context.Context, id int) (*models.Location, error)
	Insert(ctx context.Context, l models.Location) (int, error)
}

type OrderStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Order, error)
	GetByID(ctx context.Context, id int) (*models.Order, error)
	Insert(ctx context.Context, o models.Order) (int, error)
}

type ProductStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Product, error)
	GetByID(ctx context.Context, id int) (*models.Product, error)
	Insert(ctx context.Context, p models.Product) (int, error)
	Update(ctx context.Context, p models.Product) (int, error)
}

type ProductGroupStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ProductGroup, error)
	GetByID(ctx context.Context, id int) (*models.ProductGroup, error)
	Insert(ctx context.Context, g models.ProductGroup) (int, error)
	Update(ctx context.Context, g models.ProductGroup) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type ProductOrderStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ProductOrder, error)
	Insert(ctx context.Context, po models.ProductOrder) error
}

type ProductVisibilityStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ProductVisibility, error)
	GetByID(ctx context.Context, id int) (*models.ProductVisibility, error)
	Insert(ctx context.Context, v models.ProductVisibility) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type RightsStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Rights, error)
	GetByID(ctx context.Context, id int) (*models.Rights, error)
	Insert(ctx context.Context, r models.Rights) (int, error)
	Update(ctx context.Context, r models.Rights) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type RoleStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Role, error)
	GetByID(ctx context.Context, id int) (*models.Role, error)
	Insert(ctx context.Context, r models.Role) (int, error)
	Update(ctx context.Context, r models.Role) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type ServiceLinkStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ServiceLink, error)
	Insert(ctx context.Context, s models.ServiceLink) error
	Update(ctx context.Context, s models.ServiceLink) error
	Delete(ctx context.Context, foreignUserID int) error
}

type ServiceSewobeStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.ServiceSewobe, error)
	GetByID(ctx context.Context, id int) (*models.ServiceSewobe, error)
	Insert(ctx context.Context, s models.ServiceSewobe) (int, error)
	Update(ctx context.Context, s models.ServiceSewobe) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

// SettingsStore is satisfied by *repo.SettingsStore for any of the three settings tables.
type SettingsStore interface {
	GetAll(ctx context.Context) ([]models.Setting, error)
	Insert(ctx context.Context, s models.Setting) error
	Update(ctx context.Context, s models.Setting) error
	Delete(ctx context.Context, key string) error
}

type UnitStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Unit, error)
	GetByID(ctx context.Context, id int) (*models.Unit, error)
	Insert(ctx context.Context, u models.Unit) (int, error)
	Update(ctx context.Context, u models.Unit) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type UserStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
	Insert(ctx context.Context, u models.User) (int, error)
	Update(ctx context.Context, u models.User) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type VatStore interface {
	Get(ctx context.Context, q *repo.Query) ([]models.Vat, error)
	GetByID(ctx context.Context, id int) (*models.Vat, error)
	Insert(ctx context.Context, v models.Vat) (int, error)
	Update(ctx context.Context, v models.Vat) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type Stores struct {
	Account           AccountStore
	Category          CategoryStore
	Favorites         FavoritesStore
	Location          LocationStore
	Order             OrderStore
	Product           ProductStore
	ProductGroup      ProductGroupStore
	ProductOrder      ProductOrderStore
	ProductVisibility ProductVisibilityStore
	Rights            RightsStore
	Role              RoleStore
	ServiceLink       ServiceLinkStore
	ServiceSewobe     ServiceSewobeStore
	SettingsFrontend  SettingsStore
	SettingsPayment   SettingsStore
	SettingsEmail     SettingsStore
	Unit              UnitStore
	User              UserStore
	Vat               VatStore
}
