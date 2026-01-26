package test_database

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
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
	dummyAccount := models.CreateAccount("Andi", "", "Theke",
		"andiwillsaufen@bier.com", "+49 170 1234567", 9900, 10, 3)
	id, err := dbModels.account.Insert(dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	const accountId = 1
	daGama, err := dbModels.account.GetById(accountId)
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
	dummyProduct := models.CreateProduct("Pearl River Dynasty", "1000", 1, 1, 200, 1, 1)
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
	expectedPrice := new(big.Int).SetInt64(1800)
	require.Equal(t, product.Price.Int, expectedPrice)
	require.Equal(t, product.VatId, 1)
	require.Equal(t, product.ProductGroupId, 1)
	require.Equal(t, product.Size, 150)
	require.Equal(t, product.UnitId, 1)
	require.Equal(t, product.CategoryId, 2)
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
	dummyVisibility := models.CreateProductVisibility(3, 1, 3)
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

// --------------------------------------------------
// Location
// --------------------------------------------------

func TestInsertLocation(t *testing.T) {
	dummyLocation := models.CreateLocation("Nursing Home")
	id, err := dbModels.location.Insert(dummyLocation)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetLocationById(t *testing.T) {
	const locationId = 1
	location, err := dbModels.location.Get(locationId)
	if location == nil {
		t.Fail()
		t.Log("Could not get location")

		return
	}

	require.NoError(t, err)
	require.Equal(t, location.Id, locationId)
	require.Equal(t, location.Name, "Bermuda Triangle")
}

// --------------------------------------------------
// Vat
// --------------------------------------------------

func TestInsertVat(t *testing.T) {
	dummyVat := models.CreateVat(12)
	id, err := dbModels.vat.Insert(dummyVat)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetVatById(t *testing.T) {
	const vatId = 1
	vat, err := dbModels.vat.Get(vatId)
	if vat == nil {
		t.Fail()
		t.Log("Could not get vat")

		return
	}

	require.NoError(t, err)
	require.Equal(t, vat.Id, vatId)
	require.Equal(t, vat.Rate, 19)
}
