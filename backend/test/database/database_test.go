package test_database

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var dbModels *modelStructs

func TestMain(m *testing.M) {
	dbModels = &modelStructs{}
	code, err := run(m, dbModels)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

// --------------------------------------------------
// Account
// --------------------------------------------------

func TestInsertAccount(t *testing.T) {
	dummyAccount := models.CreateAccount("一嫂", "Lady Hurricane", "鄭",
		"zheng.yi.sao@redfleet.cn", "+86 20 1807-2010", 1234, 100, 1)
	id, err := dbModels.account.Insert(dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	const accountId = 1
	daGama, err := dbModels.account.Get(accountId)
	if daGama == nil {
		t.Fail()
		t.Log("Could not get account")

		return
	}

	require.NoError(t, err)
	require.Equal(t, daGama.Id, accountId)
	require.Equal(t, daGama.FirstName, "Vasco")
	require.Equal(t, daGama.Nickname, "Cape Conqueror")
	require.Equal(t, daGama.LastName, "da Gama")
	require.Equal(t, daGama.Email, "indianspice@capeofgoodhope.com")
	require.Equal(t, daGama.Phone, "+351 914 97 1498")
	require.Equal(t, daGama.Balance, 3355)
	require.Equal(t, daGama.MaxDebt, 100)
	require.Equal(t, daGama.Category, 1)
	require.Equal(t, daGama.Enabled, true)
	require.NotZero(t, daGama.CreatedAt)
}

// --------------------------------------------------
// AccountOption
// --------------------------------------------------

func TestInsertAccountOption(t *testing.T) {
	dummyOpt := models.CreateAccountOption(1, "key", "value")
	accountId, key, err := dbModels.accountOption.Insert(dummyOpt)

	require.NoError(t, err)
	assert.NotZero(t, accountId)
	assert.NotZero(t, key)
}

func TestGetAccountOptionByAccountAndKey(t *testing.T) {
	const accountId = 1
	const optKey = "deceased"
	opt, err := dbModels.accountOption.Get(accountId, optKey)
	if opt == nil {
		t.Fail()
		t.Log("Could not get account option")

		return
	}

	require.NoError(t, err)
	require.Equal(t, opt.AccountId, accountId)
	require.Equal(t, opt.Key, optKey)
	require.Equal(t, opt.Value, "true")
}

// --------------------------------------------------
// Category
// --------------------------------------------------

func TestInsertCategory(t *testing.T) {
	dummyCategory := models.CreateCategory("Guest", "user-friends", "account")
	id, err := dbModels.category.Insert(dummyCategory)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetCategoryById(t *testing.T) {
	const categoryId = 1
	cat, err := dbModels.category.Get(categoryId)
	if cat == nil {
		t.Fail()
		t.Log("Could not get category")

		return
	}

	require.NoError(t, err)
	require.Equal(t, cat.Id, categoryId)
	require.Equal(t, cat.Name, "Sailor")
	require.Equal(t, cat.Enabled, true)
	require.Equal(t, cat.Icon, "sailboat")
	require.Equal(t, cat.Type, "account")
}

// --------------------------------------------------
// Product
// --------------------------------------------------

func TestInsertProduct(t *testing.T) {
	dummyProduct := models.CreateProduct("Pearl River Dynasty", "10", 1, 200, 1, "0.19", 1)
	id, err := dbModels.product.Insert(dummyProduct)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductById(t *testing.T) {
	const productId = 1
	product, err := dbModels.product.Get(productId)
	if product == nil {
		t.Fail()
		t.Log("Could not get product")

		return
	}

	require.NoError(t, err)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, "Rota das Especiarias")
	require.Equal(t, product.Price, 1800)
	require.Equal(t, product.ProductGroupId, 1)
	require.Equal(t, product.Size, 150)
	require.Equal(t, product.UnitId, 1)
	expectedTax := 19
	require.Equal(t, product.Tax, expectedTax)
	require.Equal(t, product.CategoryId, 1)
}

// --------------------------------------------------
// ProductGroup
// --------------------------------------------------

func TestInsertProductGroup(t *testing.T) {
	dummyProductGroup := models.CreateProductGroup("Beer", 1)
	id, err := dbModels.productGroup.Insert(dummyProductGroup)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductGroupById(t *testing.T) {
	const groupId = 1
	group, err := dbModels.productGroup.Get(groupId)
	if group == nil {
		t.Fail()
		t.Log("Could not get group")

		return
	}

	require.NoError(t, err)
	require.Equal(t, group.Id, groupId)
	require.Equal(t, group.Name, "Alcohol")
	require.Equal(t, group.ParentId, 0)
}

// --------------------------------------------------
// Unit
// --------------------------------------------------

func TestInsertUnit(t *testing.T) {
	dummyUnit := models.CreateUnit("oz")
	id, err := dbModels.unit.Insert(dummyUnit)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetUnitById(t *testing.T) {
	const unitId = 1
	unit, err := dbModels.unit.Get(unitId)
	if unit == nil {
		t.Fail()
		t.Log("Could not get unit")

		return
	}

	require.NoError(t, err)
	require.Equal(t, unit.Id, unitId)
	require.Equal(t, unit.Name, "ml")
}

// --------------------------------------------------
// ProductVisibility
// --------------------------------------------------

func TestInsertProductVisibility(t *testing.T) {
	dummyVisibility := models.CreateProductVisibility(1, 2)
	id, err := dbModels.productVisibility.Insert(dummyVisibility)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestProductVisibilityById(t *testing.T) {
	const visibilityId = 1
	unit, err := dbModels.productVisibility.Get(visibilityId)
	if unit == nil {
		t.Fail()
		t.Log("Could not get product visibility")

		return
	}

	require.NoError(t, err)
	require.Equal(t, unit.Id, visibilityId)
	require.Equal(t, unit.CategoryId, 1)
	require.Equal(t, unit.ProductId, 1)
}
