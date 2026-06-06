package ws

import (
	"context"
	"log/slog"
	"testing"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
	"github.com/stretchr/testify/assert"
)

// noOpValidator satisfies msgValidator without loading a schema file.
type noOpValidator struct{}

func (noOpValidator) validate(_ []byte) error { return nil }

// --- stub store implementations ---

type stubAccountStore struct{}

func (stubAccountStore) Get(_ context.Context, _ *repo.Query) ([]models.Account, error) {
	return nil, nil
}
func (stubAccountStore) GetByID(_ context.Context, _ int) (*models.Account, error) {
	return &models.Account{}, nil
}
func (stubAccountStore) Insert(_ context.Context, _ models.Account) (int, error) { return 1, nil }
func (stubAccountStore) Update(_ context.Context, _ models.Account) (int, error) { return 1, nil }

type stubCategoryStore struct{}

func (stubCategoryStore) Get(_ context.Context, _ *repo.Query) ([]models.Category, error) {
	return nil, nil
}
func (stubCategoryStore) GetByID(_ context.Context, _ int) (*models.Category, error) {
	return &models.Category{}, nil
}
func (stubCategoryStore) Insert(_ context.Context, _ models.Category) (int, error) { return 1, nil }
func (stubCategoryStore) Update(_ context.Context, _ models.Category) (int, error) { return 1, nil }
func (stubCategoryStore) Delete(_ context.Context, _ int) (int, error)             { return 1, nil }

type stubFavoritesStore struct{}

func (stubFavoritesStore) Get(_ context.Context, _ *repo.Query) ([]models.Favorite, error) {
	return nil, nil
}
func (stubFavoritesStore) Insert(_ context.Context, _ models.Favorite) error { return nil }
func (stubFavoritesStore) Update(_ context.Context, _ models.Favorite) error { return nil }

type stubLocationStore struct{}

func (stubLocationStore) Get(_ context.Context, _ *repo.Query) ([]models.Location, error) {
	return nil, nil
}
func (stubLocationStore) GetByID(_ context.Context, _ int) (*models.Location, error) {
	return &models.Location{}, nil
}
func (stubLocationStore) Insert(_ context.Context, _ models.Location) (int, error) { return 1, nil }

type stubOrderStore struct{}

func (stubOrderStore) Get(_ context.Context, _ *repo.Query) ([]models.Order, error) {
	return nil, nil
}
func (stubOrderStore) GetByID(_ context.Context, _ int) (*models.Order, error) {
	return &models.Order{}, nil
}
func (stubOrderStore) Insert(_ context.Context, _ models.Order) (int, error) { return 1, nil }

type stubProductStore struct{}

func (stubProductStore) Get(_ context.Context, _ *repo.Query) ([]models.Product, error) {
	return nil, nil
}
func (stubProductStore) GetByID(_ context.Context, _ int) (*models.Product, error) {
	return &models.Product{}, nil
}
func (stubProductStore) Insert(_ context.Context, _ models.Product) (int, error) { return 1, nil }
func (stubProductStore) Update(_ context.Context, _ models.Product) (int, error) { return 1, nil }

type stubProductGroupStore struct{}

func (stubProductGroupStore) Get(_ context.Context, _ *repo.Query) ([]models.ProductGroup, error) {
	return nil, nil
}
func (stubProductGroupStore) GetByID(_ context.Context, _ int) (*models.ProductGroup, error) {
	return &models.ProductGroup{}, nil
}
func (stubProductGroupStore) Insert(_ context.Context, _ models.ProductGroup) (int, error) {
	return 1, nil
}
func (stubProductGroupStore) Update(_ context.Context, _ models.ProductGroup) (int, error) {
	return 1, nil
}
func (stubProductGroupStore) Delete(_ context.Context, _ int) (int, error) { return 1, nil }

type stubProductOrderStore struct{}

func (stubProductOrderStore) Get(_ context.Context, _ *repo.Query) ([]models.ProductOrder, error) {
	return nil, nil
}
func (stubProductOrderStore) Insert(_ context.Context, _ models.ProductOrder) error { return nil }

type stubProductVisibilityStore struct{}

func (stubProductVisibilityStore) Get(_ context.Context, _ *repo.Query) ([]models.ProductVisibility, error) {
	return nil, nil
}
func (stubProductVisibilityStore) GetByID(_ context.Context, _ int) (*models.ProductVisibility, error) {
	return &models.ProductVisibility{}, nil
}
func (stubProductVisibilityStore) Insert(_ context.Context, _ models.ProductVisibility) (int, error) {
	return 1, nil
}
func (stubProductVisibilityStore) Delete(_ context.Context, _ int) (int, error) { return 1, nil }

type stubRightsStore struct{}

func (stubRightsStore) Get(_ context.Context, _ *repo.Query) ([]models.Rights, error) {
	return nil, nil
}
func (stubRightsStore) GetByID(_ context.Context, _ int) (*models.Rights, error) {
	return &models.Rights{}, nil
}
func (stubRightsStore) Insert(_ context.Context, _ models.Rights) (int, error) { return 1, nil }
func (stubRightsStore) Update(_ context.Context, _ models.Rights) (int, error) { return 1, nil }
func (stubRightsStore) Delete(_ context.Context, _ int) (int, error)           { return 1, nil }

type stubRoleStore struct{}

func (stubRoleStore) Get(_ context.Context, _ *repo.Query) ([]models.Role, error) {
	return nil, nil
}
func (stubRoleStore) GetByID(_ context.Context, _ int) (*models.Role, error) {
	return &models.Role{}, nil
}
func (stubRoleStore) Insert(_ context.Context, _ models.Role) (int, error) { return 1, nil }
func (stubRoleStore) Update(_ context.Context, _ models.Role) (int, error) { return 1, nil }
func (stubRoleStore) Delete(_ context.Context, _ int) (int, error)         { return 1, nil }

type stubServiceLinkStore struct{}

func (stubServiceLinkStore) Get(_ context.Context, _ *repo.Query) ([]models.ServiceLink, error) {
	return nil, nil
}
func (stubServiceLinkStore) Insert(_ context.Context, _ models.ServiceLink) error { return nil }
func (stubServiceLinkStore) Update(_ context.Context, _ models.ServiceLink) error { return nil }
func (stubServiceLinkStore) Delete(_ context.Context, _ int) error                { return nil }

type stubServiceSewobeStore struct{}

func (stubServiceSewobeStore) Get(_ context.Context, _ *repo.Query) ([]models.ServiceSewobe, error) {
	return nil, nil
}
func (stubServiceSewobeStore) GetByID(_ context.Context, _ int) (*models.ServiceSewobe, error) {
	return &models.ServiceSewobe{}, nil
}
func (stubServiceSewobeStore) Insert(_ context.Context, _ models.ServiceSewobe) (int, error) {
	return 1, nil
}
func (stubServiceSewobeStore) Update(_ context.Context, _ models.ServiceSewobe) (int, error) {
	return 1, nil
}
func (stubServiceSewobeStore) Delete(_ context.Context, _ int) (int, error) { return 1, nil }

type stubSettingsStore struct{}

func (stubSettingsStore) GetAll(_ context.Context) ([]models.Setting, error) { return nil, nil }
func (stubSettingsStore) Insert(_ context.Context, _ models.Setting) error   { return nil }
func (stubSettingsStore) Update(_ context.Context, _ models.Setting) error   { return nil }
func (stubSettingsStore) Delete(_ context.Context, _ string) error           { return nil }

type stubUnitStore struct{}

func (stubUnitStore) Get(_ context.Context, _ *repo.Query) ([]models.Unit, error) {
	return nil, nil
}
func (stubUnitStore) GetByID(_ context.Context, _ int) (*models.Unit, error) {
	return &models.Unit{}, nil
}
func (stubUnitStore) Insert(_ context.Context, _ models.Unit) (int, error) { return 1, nil }
func (stubUnitStore) Update(_ context.Context, _ models.Unit) (int, error) { return 1, nil }
func (stubUnitStore) Delete(_ context.Context, _ int) (int, error)         { return 1, nil }

type stubUserStore struct{}

func (stubUserStore) Get(_ context.Context, _ *repo.Query) ([]models.User, error) {
	return nil, nil
}
func (stubUserStore) GetByID(_ context.Context, _ int) (*models.User, error) {
	return &models.User{}, nil
}
func (stubUserStore) Insert(_ context.Context, _ models.User) (int, error) { return 1, nil }
func (stubUserStore) Update(_ context.Context, _ models.User) (int, error) { return 1, nil }
func (stubUserStore) Delete(_ context.Context, _ int) (int, error)         { return 1, nil }

type stubVatStore struct{}

func (stubVatStore) Get(_ context.Context, _ *repo.Query) ([]models.Vat, error) { return nil, nil }
func (stubVatStore) GetByID(_ context.Context, _ int) (*models.Vat, error) {
	return &models.Vat{}, nil
}
func (stubVatStore) Insert(_ context.Context, _ models.Vat) (int, error) { return 1, nil }
func (stubVatStore) Update(_ context.Context, _ models.Vat) (int, error) { return 1, nil }
func (stubVatStore) Delete(_ context.Context, _ int) (int, error)        { return 1, nil }

// --- test helper ---

func newTestService() *Service {
	return &Service{
		log:       slog.Default(),
		validator: noOpValidator{},
		stores: Stores{
			Account:           stubAccountStore{},
			Category:          stubCategoryStore{},
			Favorites:         stubFavoritesStore{},
			Location:          stubLocationStore{},
			Order:             stubOrderStore{},
			Product:           stubProductStore{},
			ProductGroup:      stubProductGroupStore{},
			ProductOrder:      stubProductOrderStore{},
			ProductVisibility: stubProductVisibilityStore{},
			Rights:            stubRightsStore{},
			Role:              stubRoleStore{},
			ServiceLink:       stubServiceLinkStore{},
			ServiceSewobe:     stubServiceSewobeStore{},
			SettingsFrontend:  stubSettingsStore{},
			SettingsPayment:   stubSettingsStore{},
			SettingsEmail:     stubSettingsStore{},
			Unit:              stubUnitStore{},
			User:              stubUserStore{},
			Vat:               stubVatStore{},
		},
	}
}

// --- tests ---

func TestProcessQueryHandlesAllTables(t *testing.T) {
	s := newTestService()
	ctx := context.Background()

	tables := []repo.TableName{
		repo.TableAccount,
		repo.TableCategory,
		repo.TableFavorites,
		repo.TableLocation,
		repo.TableOrder,
		repo.TableProduct,
		repo.TableProductGroup,
		repo.TableProductOrder,
		repo.TableProductVisibility,
		repo.TableRights,
		repo.TableRole,
		repo.TableServiceLink,
		repo.TableServiceSewobe,
		repo.TableSettingsFrontend,
		repo.TableSettingsPayment,
		repo.TableSettingsEmail,
		repo.TableUnit,
		repo.TableUser,
		repo.TableVat,
	}

	for _, table := range tables {
		t.Run(string(table), func(t *testing.T) {
			msg := Message{
				Type:    MsgTypeQuery,
				Payload: map[string]interface{}{"table": string(table)},
			}
			_, wsErr := s.processQuery(ctx, msg)
			assert.Nil(t, wsErr, "table %q returned unexpected error", table)
		})
	}
}

func TestProcessMutationHandlesAllTables(t *testing.T) {
	s := newTestService()
	ctx := context.Background()

	emptyValues := map[string]interface{}{}

	cases := []struct {
		op    repo.Operation
		table repo.TableName
		where map[string]string
	}{
		// inserts
		{repo.OpInsert, repo.TableAccount, nil},
		{repo.OpInsert, repo.TableCategory, nil},
		{repo.OpInsert, repo.TableFavorites, nil},
		{repo.OpInsert, repo.TableLocation, nil},
		{repo.OpInsert, repo.TableOrder, nil},
		{repo.OpInsert, repo.TableProduct, nil},
		{repo.OpInsert, repo.TableProductGroup, nil},
		{repo.OpInsert, repo.TableProductOrder, nil},
		{repo.OpInsert, repo.TableProductVisibility, nil},
		{repo.OpInsert, repo.TableRights, nil},
		{repo.OpInsert, repo.TableRole, nil},
		{repo.OpInsert, repo.TableServiceLink, nil},
		{repo.OpInsert, repo.TableServiceSewobe, nil},
		{repo.OpInsert, repo.TableSettingsFrontend, nil},
		{repo.OpInsert, repo.TableSettingsPayment, nil},
		{repo.OpInsert, repo.TableSettingsEmail, nil},
		{repo.OpInsert, repo.TableUnit, nil},
		{repo.OpInsert, repo.TableUser, nil},
		{repo.OpInsert, repo.TableVat, nil},
		// updates
		{repo.OpUpdate, repo.TableAccount, nil},
		{repo.OpUpdate, repo.TableCategory, nil},
		{repo.OpUpdate, repo.TableFavorites, nil},
		{repo.OpUpdate, repo.TableProduct, nil},
		{repo.OpUpdate, repo.TableProductGroup, nil},
		{repo.OpUpdate, repo.TableRights, nil},
		{repo.OpUpdate, repo.TableRole, nil},
		{repo.OpUpdate, repo.TableServiceLink, nil},
		{repo.OpUpdate, repo.TableServiceSewobe, nil},
		{repo.OpUpdate, repo.TableSettingsFrontend, nil},
		{repo.OpUpdate, repo.TableSettingsPayment, nil},
		{repo.OpUpdate, repo.TableSettingsEmail, nil},
		{repo.OpUpdate, repo.TableUnit, nil},
		{repo.OpUpdate, repo.TableUser, nil},
		{repo.OpUpdate, repo.TableVat, nil},
		// deletes
		{repo.OpDelete, repo.TableCategory, map[string]string{"category_id": "1"}},
		{repo.OpDelete, repo.TableProductGroup, map[string]string{"product_group_id": "1"}},
		{repo.OpDelete, repo.TableProductVisibility, map[string]string{"product_visibility_id": "1"}},
		{repo.OpDelete, repo.TableRights, map[string]string{"rights_id": "1"}},
		{repo.OpDelete, repo.TableRole, map[string]string{"role_id": "1"}},
		{repo.OpDelete, repo.TableServiceLink, map[string]string{"foreign_user_id": "1"}},
		{repo.OpDelete, repo.TableServiceSewobe, map[string]string{"service_sewobe_id": "1"}},
		{repo.OpDelete, repo.TableSettingsFrontend, map[string]string{"key": "k"}},
		{repo.OpDelete, repo.TableSettingsPayment, map[string]string{"key": "k"}},
		{repo.OpDelete, repo.TableSettingsEmail, map[string]string{"key": "k"}},
		{repo.OpDelete, repo.TableUnit, map[string]string{"unit_id": "1"}},
		{repo.OpDelete, repo.TableUser, map[string]string{"user_id": "1"}},
		{repo.OpDelete, repo.TableVat, map[string]string{"vat_id": "1"}},
	}

	for _, tc := range cases {
		name := string(tc.op) + "/" + string(tc.table)
		t.Run(name, func(t *testing.T) {
			mutation := &Mutation{
				Operation: tc.op,
				Table:     tc.table,
				Values:    emptyValues,
				Where:     tc.where,
			}
			_, _, wsErr := s.processMutation(ctx, mutation)
			assert.Nil(t, wsErr, "unexpected error for %s", name)
		})
	}
}

// TestProcessMutationAccountDeleteRejected documents that account deletion is intentionally unsupported.
func TestProcessMutationAccountDeleteRejected(t *testing.T) {
	s := newTestService()
	_, _, wsErr := s.processMutation(context.Background(), &Mutation{
		Operation: repo.OpDelete,
		Table:     repo.TableAccount,
	})
	if assert.NotNil(t, wsErr) {
		assert.Equal(t, WSErrorCode(WSInvalidOperation), wsErr.Code)
	}
}
