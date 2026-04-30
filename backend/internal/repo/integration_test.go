package repo_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

// --------------------------------------------------
// Account
// --------------------------------------------------

func TestInsertAccount(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	dummyAccount := models.CreateAccount("Edward", "Blackbeard", "Teach",
		"blackbeard@queenanne.com", "+1 910 555 1718", 9900, 10, 3)
	id, err := store.Insert(context.Background(), dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccounts(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableAccount}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, accounts, 5)
}

func TestGetAccountsFilterByDisabled(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table:  repo.TableAccount,
		Filter: []repo.Filter{{Column: "enabled", Operator: repo.Eq, Value: "false"}},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, accounts, 1)
	assert.Equal(t, "Davy", accounts[0].FirstName)
	assert.False(t, accounts[0].Enabled)
}

func TestGetAccountsWithFilter(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table:  repo.TableAccount,
		Filter: []repo.Filter{{Column: "category", Operator: repo.Eq, Value: "1"}},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, accounts, 2)
	for _, a := range accounts {
		assert.Equal(t, 1, a.Category)
	}
}

func TestGetAccountsWithMultipleFilters(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table: repo.TableAccount,
		Filter: []repo.Filter{
			{Column: "category", Operator: repo.Eq, Value: "1"},
			{Column: "balance", Operator: repo.Gt, Value: "2000"},
		},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, accounts, 1)
	assert.Equal(t, 1, accounts[0].Category)
	assert.Greater(t, accounts[0].Balance, 2000)
}

func TestGetAccountsWithFilterNoMatch(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table:  repo.TableAccount,
		Filter: []repo.Filter{{Column: "balance", Operator: repo.Eq, Value: "-1"}},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Empty(t, accounts)
}

func TestGetAccountsWithSortAsc(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table: repo.TableAccount,
		Sort:  &repo.Sorting{Column: "balance", Order: repo.Asc},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for i := 1; i < len(accounts); i++ {
		assert.LessOrEqual(t, accounts[i-1].Balance, accounts[i].Balance)
	}
}

func TestGetAccountsWithSortDesc(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table: repo.TableAccount,
		Sort:  &repo.Sorting{Column: "balance", Order: repo.Desc},
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for i := 1; i < len(accounts); i++ {
		assert.GreaterOrEqual(t, accounts[i-1].Balance, accounts[i].Balance)
	}
}

func TestGetAccountsWithLimit(t *testing.T) {
	limit := 2
	store := &repo.AccountStore{DB: beginTx(t)}
	query := repo.Query{
		Table: repo.TableAccount,
		Limit: &limit,
	}
	accounts, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, accounts, 2)
}

func TestGetAccountById(t *testing.T) {
	const accountID = 1
	store := &repo.AccountStore{DB: beginTx(t)}
	daGama, err := store.GetByID(context.Background(), accountID)

	require.NoError(t, err)
	require.NotNil(t, daGama)
	require.Equal(t, accountID, daGama.ID)
	require.Equal(t, "Vasco", daGama.FirstName)
	require.Equal(t, "Cape Conqueror", daGama.Nickname)
	require.Equal(t, "da Gama", daGama.LastName)
	require.Equal(t, "indianspice@capeofgoodhope.com", daGama.Email)
	require.Equal(t, "+351 914 97 1498", daGama.Phone)
	require.Equal(t, 3355, daGama.Balance)
	require.Equal(t, 10000, daGama.MaxDebt)
	require.Equal(t, 1, daGama.Category)
	require.Equal(t, true, daGama.Enabled)
	require.NotZero(t, daGama.CreatedAt)
}

// --------------------------------------------------
// AccountOption
// --------------------------------------------------

func TestInsertAccountOption(t *testing.T) {
	store := &repo.AccountOptionStore{DB: beginTx(t)}
	dummyOpt := models.CreateAccountOption(1, "key", "value")
	accountID, key, err := store.Insert(context.Background(), dummyOpt)

	require.NoError(t, err)
	assert.NotZero(t, accountID)
	assert.NotZero(t, key)
}

func TestGetAccountOptionByAccountAndKey(t *testing.T) {
	const accountID = 1
	const optKey = "deceased"
	store := &repo.AccountOptionStore{DB: beginTx(t)}
	opt, err := store.Get(context.Background(), accountID, optKey)

	require.NoError(t, err)
	require.NotNil(t, opt)
	require.Equal(t, accountID, opt.AccountID)
	require.Equal(t, optKey, opt.Key)
	require.Equal(t, "true", opt.Value)
}

func TestInsertAccountOptionAndReadBack(t *testing.T) {
	store := &repo.AccountOptionStore{DB: beginTx(t)}
	opt := models.CreateAccountOption(1, "favourite_sea", "Caribbean")

	accountID, key, err := store.Insert(context.Background(), opt)
	require.NoError(t, err)
	require.Equal(t, 1, accountID)
	require.Equal(t, "favourite_sea", key)

	saved, err := store.Get(context.Background(), accountID, key)
	require.NoError(t, err)
	require.NotNil(t, saved)
	assert.Equal(t, 1, saved.AccountID)
	assert.Equal(t, "favourite_sea", saved.Key)
	assert.Equal(t, "Caribbean", saved.Value)
}

// --------------------------------------------------
// Category
// --------------------------------------------------

func TestInsertCategory(t *testing.T) {
	store := &repo.CategoryStore{DB: beginTx(t)}
	dummyCategory := models.CreateCategory("Guest", "user-friends", "account")
	id, err := store.Insert(context.Background(), dummyCategory)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetCategories(t *testing.T) {
	store := &repo.CategoryStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableCategory}
	categories, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, categories, 4)
}

func TestGetCategoryById(t *testing.T) {
	const categoryID = 1
	store := &repo.CategoryStore{DB: beginTx(t)}
	cat, err := store.GetByID(context.Background(), categoryID)

	require.NoError(t, err)
	require.NotNil(t, cat)
	require.Equal(t, categoryID, cat.ID)
	require.Equal(t, "Sailor", cat.Name)
	require.Equal(t, true, cat.Enabled)
	require.Equal(t, "sailboat", cat.Icon)
	require.Equal(t, "account", cat.Type)
}

func TestGetCategoryByIdNotFound(t *testing.T) {
	store := &repo.CategoryStore{DB: beginTx(t)}
	cat, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, cat)
}

// --------------------------------------------------
// Product
// --------------------------------------------------

func TestInsertProduct(t *testing.T) {
	store := &repo.ProductStore{DB: beginTx(t)}
	dummyProduct := models.CreateProduct("Pearl River Dynasty", 1000, 1, 1, 200, 1, 1)
	id, err := store.Insert(context.Background(), dummyProduct)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProducts(t *testing.T) {
	store := &repo.ProductStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableProduct}
	products, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, products, 3)
}

func TestGetProductById(t *testing.T) {
	const productID = 1
	store := &repo.ProductStore{DB: beginTx(t)}
	product, err := store.GetByID(context.Background(), productID)

	require.NoError(t, err)
	require.NotNil(t, product)
	require.Equal(t, product.ID, productID)
	require.Equal(t, "Rota das Especiarias", product.Name)
	require.Equal(t, 1800, product.Price)
	require.Equal(t, 1, product.VatID)
	require.Equal(t, 1, product.ProductGroupID)
	require.Equal(t, 150, product.Size)
	require.Equal(t, 1, product.UnitID)
	require.Equal(t, 2, product.CategoryID)
}

func TestGetProductByIdNotFound(t *testing.T) {
	store := &repo.ProductStore{DB: beginTx(t)}
	product, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, product)
}

// --------------------------------------------------
// ProductGroup
// --------------------------------------------------

func TestInsertProductGroup(t *testing.T) {
	parentID := 1
	store := &repo.ProductGroupStore{DB: beginTx(t)}
	dummyProductGroup := models.CreateProductGroup("Beer", &parentID)
	id, err := store.Insert(context.Background(), dummyProductGroup)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductGroups(t *testing.T) {
	store := &repo.ProductGroupStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableProductGroup}
	groups, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, groups, 5)
}

func TestGetProductGroupById(t *testing.T) {
	const groupID = 1
	store := &repo.ProductGroupStore{DB: beginTx(t)}
	group, err := store.GetByID(context.Background(), groupID)

	require.NoError(t, err)
	require.NotNil(t, group)
	require.Equal(t, groupID, group.ID)
	require.Equal(t, "Alcohol", group.Name)
	require.Nil(t, group.ParentID)
}

func TestGetProductGroupByIdNotFound(t *testing.T) {
	store := &repo.ProductGroupStore{DB: beginTx(t)}
	group, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, group)
}

func TestGetProductGroupWithParent(t *testing.T) {
	// group 3 is "Port Wine" with parent 1 ("Alcohol")
	const groupID = 3
	store := &repo.ProductGroupStore{DB: beginTx(t)}
	group, err := store.GetByID(context.Background(), groupID)

	require.NoError(t, err)
	require.NotNil(t, group)
	assert.Equal(t, groupID, group.ID)
	assert.Equal(t, "Port Wine", group.Name)
	require.NotNil(t, group.ParentID)
	assert.Equal(t, 1, *group.ParentID)
}

// --------------------------------------------------
// Unit
// --------------------------------------------------

func TestInsertUnit(t *testing.T) {
	store := &repo.UnitStore{DB: beginTx(t)}
	dummyUnit := models.CreateUnit("oz")
	id, err := store.Insert(context.Background(), dummyUnit)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetUnits(t *testing.T) {
	store := &repo.UnitStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableUnit}
	units, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, units, 2)
}

func TestGetUnitById(t *testing.T) {
	const unitID = 1
	store := &repo.UnitStore{DB: beginTx(t)}
	unit, err := store.GetByID(context.Background(), unitID)

	require.NoError(t, err)
	require.NotNil(t, unit)
	require.Equal(t, unitID, unit.ID)
	require.Equal(t, "ml", unit.Name)
}

func TestGetUnitByIdNotFound(t *testing.T) {
	store := &repo.UnitStore{DB: beginTx(t)}
	unit, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, unit)
}

// --------------------------------------------------
// ProductVisibility
// --------------------------------------------------

func TestInsertProductVisibility(t *testing.T) {
	store := &repo.ProductVisibilityStore{DB: beginTx(t)}
	dummyVisibility := models.CreateProductVisibility(3, 1, 3)
	id, err := store.Insert(context.Background(), dummyVisibility)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductVisibilities(t *testing.T) {
	store := &repo.ProductVisibilityStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableProductVisibility}
	visibilities, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, visibilities, 7)
}

func TestProductVisibilityById(t *testing.T) {
	const visibilityID = 1
	store := &repo.ProductVisibilityStore{DB: beginTx(t)}
	vis, err := store.GetByID(context.Background(), visibilityID)

	require.NoError(t, err)
	require.NotNil(t, vis)
	require.Equal(t, visibilityID, vis.ID)
	require.Equal(t, 1, vis.CategoryID)
	require.Equal(t, 1, vis.ProductID)
}

func TestProductVisibilityByIdNotFound(t *testing.T) {
	store := &repo.ProductVisibilityStore{DB: beginTx(t)}
	vis, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, vis)
}

// --------------------------------------------------
// Location
// --------------------------------------------------

func TestInsertLocation(t *testing.T) {
	store := &repo.LocationStore{DB: beginTx(t)}
	dummyLocation := models.CreateLocation("Nursing Home")
	id, err := store.Insert(context.Background(), dummyLocation)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetLocations(t *testing.T) {
	store := &repo.LocationStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableLocation}
	locations, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, locations, 2)
}

func TestGetLocationById(t *testing.T) {
	const locationID = 1
	store := &repo.LocationStore{DB: beginTx(t)}
	location, err := store.GetByID(context.Background(), locationID)

	require.NoError(t, err)
	require.NotNil(t, location)
	require.Equal(t, locationID, location.ID)
	require.Equal(t, "Bermuda Triangle", location.Name)
}

func TestGetLocationByIdNotFound(t *testing.T) {
	store := &repo.LocationStore{DB: beginTx(t)}
	location, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, location)
}

// --------------------------------------------------
// Vat
// --------------------------------------------------

func TestInsertVat(t *testing.T) {
	store := &repo.VatStore{DB: beginTx(t)}
	dummyVat := models.CreateVat(12)
	id, err := store.Insert(context.Background(), dummyVat)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetVats(t *testing.T) {
	store := &repo.VatStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableVat}
	vats, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, vats, 1)
}

func TestGetVatById(t *testing.T) {
	const vatID = 1
	store := &repo.VatStore{DB: beginTx(t)}
	vat, err := store.GetByID(context.Background(), vatID)

	require.NoError(t, err)
	require.NotNil(t, vat)
	require.Equal(t, vatID, vat.ID)
	require.Equal(t, 19, vat.Rate)
}

func TestGetVatByIdNotFound(t *testing.T) {
	store := &repo.VatStore{DB: beginTx(t)}
	vat, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, vat)
}

func TestGetAccountByIdNotFound(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	account, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, account)
}

func TestGetAccountOptionNotFound(t *testing.T) {
	store := &repo.AccountOptionStore{DB: beginTx(t)}
	opt, err := store.Get(context.Background(), 1, "nonexistent_key")

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, opt)
}

// --------------------------------------------------
// Favorites
// --------------------------------------------------

func TestInsertFavorite(t *testing.T) {
	store := &repo.FavoritesStore{DB: beginTx(t)}
	fav := models.CreateFavorite(2, 1, 1)
	err := store.Insert(context.Background(), fav)

	require.NoError(t, err)
}

func TestGetFavorites(t *testing.T) {
	store := &repo.FavoritesStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableFavorites}
	favorites, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, favorites, 3)
}

func TestGetFavoritesByAccount(t *testing.T) {
	store := &repo.FavoritesStore{DB: beginTx(t)}
	query := repo.Query{
		Table:  repo.TableFavorites,
		Filter: []repo.Filter{{Column: "account", Operator: repo.Eq, Value: "1"}},
	}
	favorites, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, favorites, 2)
	for _, fav := range favorites {
		assert.Equal(t, 1, fav.AccountID)
	}
}

func TestUpdateFavorite(t *testing.T) {
	store := &repo.FavoritesStore{DB: beginTx(t)}

	err := store.Update(context.Background(), models.CreateFavorite(1, 1, 10))
	require.NoError(t, err)

	query := repo.Query{
		Table: repo.TableFavorites,
		Filter: []repo.Filter{
			{Column: "product", Operator: repo.Eq, Value: "1"},
			{Column: "account", Operator: repo.Eq, Value: "1"},
		},
	}
	favorites, err := store.Get(context.Background(), &query)
	require.NoError(t, err)
	require.Len(t, favorites, 1)
	assert.Equal(t, 10, favorites[0].Count)
}

func TestDeleteFavorite(t *testing.T) {
	store := &repo.FavoritesStore{DB: beginTx(t)}
	err := store.Delete(context.Background(), 1, 1)

	require.NoError(t, err)

	query := repo.Query{
		Table: repo.TableFavorites,
		Filter: []repo.Filter{
			{Column: "product", Operator: repo.Eq, Value: "1"},
			{Column: "account", Operator: repo.Eq, Value: "1"},
		},
	}
	favorites, err := store.Get(context.Background(), &query)
	require.NoError(t, err)
	assert.Empty(t, favorites)
}

// --------------------------------------------------
// Order
// --------------------------------------------------

func TestInsertOrder(t *testing.T) {
	store := &repo.OrderStore{DB: beginTx(t)}
	order := models.CreateOrder(1)
	id, err := store.Insert(context.Background(), order)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetOrders(t *testing.T) {
	store := &repo.OrderStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableOrder}
	orders, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, orders, 2)
}

func TestGetOrderById(t *testing.T) {
	const orderID = 1
	store := &repo.OrderStore{DB: beginTx(t)}
	order, err := store.GetByID(context.Background(), orderID)

	require.NoError(t, err)
	require.NotNil(t, order)
	assert.Equal(t, orderID, order.ID)
	assert.Equal(t, 1, order.AccountID)
	assert.NotZero(t, order.CreatedAt)
}

func TestGetOrderByIdNotFound(t *testing.T) {
	store := &repo.OrderStore{DB: beginTx(t)}
	order, err := store.GetByID(context.Background(), 99999)

	require.ErrorIs(t, err, database.ErrNoRecord)
	assert.Nil(t, order)
}

func TestInsertOrderAndReadBack(t *testing.T) {
	store := &repo.OrderStore{DB: beginTx(t)}
	order := models.CreateOrder(3)

	id, err := store.Insert(context.Background(), order)
	require.NoError(t, err)
	require.NotZero(t, id)

	saved, err := store.GetByID(context.Background(), id)
	require.NoError(t, err)
	require.NotNil(t, saved)
	assert.Equal(t, 3, saved.AccountID)
	assert.NotZero(t, saved.CreatedAt)
}

// --------------------------------------------------
// ProductOrder
// --------------------------------------------------

func TestInsertProductOrder(t *testing.T) {
	store := &repo.ProductOrderStore{DB: beginTx(t)}
	item := models.CreateProductOrder(1, 2, 1, 600)
	err := store.Insert(context.Background(), item)

	require.NoError(t, err)
}

func TestGetProductOrders(t *testing.T) {
	store := &repo.ProductOrderStore{DB: beginTx(t)}
	query := repo.Query{Table: repo.TableProductOrder}
	items, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	assert.Len(t, items, 3)
}

func TestGetProductOrdersByOrderID(t *testing.T) {
	store := &repo.ProductOrderStore{DB: beginTx(t)}
	query := repo.Query{
		Table:  repo.TableProductOrder,
		Filter: []repo.Filter{{Column: "order_id", Operator: repo.Eq, Value: "1"}},
	}
	items, err := store.Get(context.Background(), &query)

	require.NoError(t, err)
	require.Len(t, items, 2)
	for _, item := range items {
		assert.Equal(t, 1, item.OrderID)
	}
}

func TestDeleteProductOrder(t *testing.T) {
	store := &repo.ProductOrderStore{DB: beginTx(t)}
	err := store.Delete(context.Background(), 1, 1)

	require.NoError(t, err)

	query := repo.Query{
		Table:  repo.TableProductOrder,
		Filter: []repo.Filter{{Column: "order_id", Operator: repo.Eq, Value: "1"}},
	}
	items, err := store.Get(context.Background(), &query)
	require.NoError(t, err)
	assert.Len(t, items, 1)
}

// --------------------------------------------------
// WithTx
// --------------------------------------------------

func TestWithTxCommitsOnSuccess(t *testing.T) {
	var insertedID int
	err := dbPool.WithTx(context.Background(), func(tx pgx.Tx) error {
		store := &repo.CategoryStore{DB: tx}
		id, err := store.Insert(context.Background(), models.CreateCategory("Temporary", "trash", "account"))
		if err != nil {
			return err
		}
		insertedID = id
		return nil
	})
	require.NoError(t, err)

	// beginTx rolls back at cleanup — use dbPool directly to delete the committed row
	t.Cleanup(func() {
		_, _ = dbPool.Exec(context.Background(), "DELETE FROM category WHERE category_id = $1", insertedID)
	})

	// Verify the row is visible outside the transaction
	store := &repo.CategoryStore{DB: beginTx(t)}
	cat, err := store.GetByID(context.Background(), insertedID)
	require.NoError(t, err)
	assert.Equal(t, "Temporary", cat.Name)
}

func TestWithTxRollsBackOnError(t *testing.T) {
	sentinel := errors.New("abort")
	var insertedID int
	err := dbPool.WithTx(context.Background(), func(tx pgx.Tx) error {
		store := &repo.CategoryStore{DB: tx}
		id, err := store.Insert(context.Background(), models.CreateCategory("ShouldNotExist", "trash", "account"))
		if err != nil {
			return err
		}
		insertedID = id
		return sentinel
	})
	require.ErrorIs(t, err, sentinel)

	// Verify the row was rolled back and is not visible
	store := &repo.CategoryStore{DB: beginTx(t)}
	_, err = store.GetByID(context.Background(), insertedID)
	require.ErrorIs(t, err, database.ErrNoRecord)
}

// --------------------------------------------------
// Insert + read-back
// --------------------------------------------------

func TestInsertAccountAndReadBack(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}
	account := models.CreateAccount("Edward", "Blackbeard", "Teach",
		"blackbeard@queenanne.com", "+1 910 555 1718", 9900, 500, 3)

	id, err := store.Insert(context.Background(), account)
	require.NoError(t, err)
	require.NotZero(t, id)

	saved, err := store.GetByID(context.Background(), id)
	require.NoError(t, err)
	require.NotNil(t, saved)
	assert.Equal(t, "Edward", saved.FirstName)
	assert.Equal(t, "Blackbeard", saved.Nickname)
	assert.Equal(t, "Teach", saved.LastName)
	assert.Equal(t, "blackbeard@queenanne.com", saved.Email)
	assert.Equal(t, "+1 910 555 1718", saved.Phone)
	assert.Equal(t, 9900, saved.Balance)
	assert.Equal(t, 500, saved.MaxDebt)
	assert.Equal(t, 3, saved.Category)
	assert.True(t, saved.Enabled)
}

// --------------------------------------------------
// Update
// --------------------------------------------------

func TestUpdateCategory(t *testing.T) {
	store := &repo.CategoryStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.Name = "Navigator"
	original.Enabled = false
	original.Icon = "compass"

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "Navigator", saved.Name)
	assert.False(t, saved.Enabled)
	assert.Equal(t, "compass", saved.Icon)
}

func TestUpdateUnit(t *testing.T) {
	store := &repo.UnitStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.Name = "cl"

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "cl", saved.Name)
}

func TestUpdateVat(t *testing.T) {
	store := &repo.VatStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.Rate = 7

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 7, saved.Rate)
}

func TestUpdateProductGroup(t *testing.T) {
	store := &repo.ProductGroupStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.Name = "Spirits"

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "Spirits", saved.Name)
	assert.Nil(t, saved.ParentID)
}

func TestUpdateProduct(t *testing.T) {
	store := &repo.ProductStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.Name = "Vinho do Porto"
	original.Price = 2500

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "Vinho do Porto", saved.Name)
	assert.Equal(t, 2500, saved.Price)
}

func TestUpdateAccountOption(t *testing.T) {
	store := &repo.AccountOptionStore{DB: beginTx(t)}

	opt := models.CreateAccountOption(1, "deceased", "false")
	err := store.Update(context.Background(), opt)
	require.NoError(t, err)

	saved, err := store.Get(context.Background(), 1, "deceased")
	require.NoError(t, err)
	assert.Equal(t, "false", saved.Value)
}

func TestUpdateAccount(t *testing.T) {
	store := &repo.AccountStore{DB: beginTx(t)}

	original, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, original)

	original.FirstName = "Vasconius"
	original.MaxDebt = 99999
	original.Enabled = false

	updatedID, err := store.Update(context.Background(), *original)
	require.NoError(t, err)
	assert.Equal(t, 1, updatedID)

	saved, err := store.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, saved)
	assert.Equal(t, "Vasconius", saved.FirstName)
	assert.Equal(t, 99999, saved.MaxDebt)
	assert.False(t, saved.Enabled)
	assert.Equal(t, original.Balance, saved.Balance) // Balance must not be updated
}

// --------------------------------------------------
// Delete
// --------------------------------------------------

func TestDeleteCategory(t *testing.T) {
	store := &repo.CategoryStore{DB: beginTx(t)}

	// insert a category with no dependents so we can safely delete it
	id, err := store.Insert(context.Background(), models.CreateCategory("Temporary", "trash", "account"))
	require.NoError(t, err)

	deletedID, err := store.Delete(context.Background(), id)
	require.NoError(t, err)
	assert.Equal(t, id, deletedID)

	_, err = store.GetByID(context.Background(), id)
	require.ErrorIs(t, err, database.ErrNoRecord)
}

func TestDeleteUnit(t *testing.T) {
	store := &repo.UnitStore{DB: beginTx(t)}

	// unit_id 2 (pcs) has no FK dependents in seed data
	deletedID, err := store.Delete(context.Background(), 2)
	require.NoError(t, err)
	assert.Equal(t, 2, deletedID)

	_, err = store.GetByID(context.Background(), 2)
	require.ErrorIs(t, err, database.ErrNoRecord)
}

func TestDeleteVat(t *testing.T) {
	store := &repo.VatStore{DB: beginTx(t)}

	// insert a vat first so we can delete it without FK conflicts
	id, err := store.Insert(context.Background(), models.CreateVat(5))
	require.NoError(t, err)

	deletedID, err := store.Delete(context.Background(), id)
	require.NoError(t, err)
	assert.Equal(t, id, deletedID)

	_, err = store.GetByID(context.Background(), id)
	require.ErrorIs(t, err, database.ErrNoRecord)
}

func TestDeleteProductGroup(t *testing.T) {
	store := &repo.ProductGroupStore{DB: beginTx(t)}

	// insert a leaf group so we can delete it without FK conflicts
	parentID := 1
	id, err := store.Insert(context.Background(), models.CreateProductGroup("Temporary", &parentID))
	require.NoError(t, err)

	deletedID, err := store.Delete(context.Background(), id)
	require.NoError(t, err)
	assert.Equal(t, id, deletedID)

	_, err = store.GetByID(context.Background(), id)
	require.ErrorIs(t, err, database.ErrNoRecord)
}

func TestDeleteProductVisibility(t *testing.T) {
	store := &repo.ProductVisibilityStore{DB: beginTx(t)}

	deletedID, err := store.Delete(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 1, deletedID)

	_, err = store.GetByID(context.Background(), 1)
	require.ErrorIs(t, err, database.ErrNoRecord)
}
