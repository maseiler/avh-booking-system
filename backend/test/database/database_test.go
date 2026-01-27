package test_database

import (
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
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
	require.Equal(t, accountId, daGama.Id)
	require.Equal(t, "Vasco", daGama.FirstName)
	require.Equal(t, "Cape Conqueror", daGama.Nickname)
	require.Equal(t, "da Gama", daGama.LastName)
	require.Equal(t, "indianspice@capeofgoodhope.com", daGama.Email)
	require.Equal(t, "+351 914 97 1498", daGama.Phone)
	require.Equal(t, 3355, daGama.Balance)
	require.Equal(t, 100, daGama.MaxDebt)
	require.Equal(t, 1, daGama.Category)
	require.Equal(t, true, daGama.Enabled)
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
	require.Equal(t, accountId, opt.AccountId)
	require.Equal(t, optKey, opt.Key)
	require.Equal(t, "true", opt.Value)
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
	cat, err := dbModels.category.GetById(categoryId)
	if cat == nil {
		t.Fail()
		t.Log("Could not get category")

		return
	}

	require.NoError(t, err)
	require.Equal(t, categoryId, cat.Id)
	require.Equal(t, "Sailor", cat.Name)
	require.Equal(t, true, cat.Enabled)
	require.Equal(t, "sailboat", cat.Icon)
	require.Equal(t, "account", cat.Type)
}

// --------------------------------------------------
// Product
// --------------------------------------------------

func TestInsertProduct(t *testing.T) {
	dummyProduct := models.CreateProduct("Pearl River Dynasty", 1000, 1, 1, 200, 1, 1)
	id, err := dbModels.product.Insert(dummyProduct)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductById(t *testing.T) {
	const productId = 1
	product, err := dbModels.product.GetById(productId)
	if product == nil {
		t.Fail()
		t.Log("Could not get product")

		return
	}

	require.NoError(t, err)
	require.Equal(t, product.Id, productId)
	require.Equal(t, "Rota das Especiarias", product.Name)
	require.Equal(t, 1800, product.Price)
	require.Equal(t, 1, product.VatId)
	require.Equal(t, 1, product.ProductGroupId)
	require.Equal(t, 150, product.Size)
	require.Equal(t, 1, product.UnitId)
	require.Equal(t, 2, product.CategoryId)
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
	group, err := dbModels.productGroup.GetById(groupId)
	if group == nil {
		t.Fail()
		t.Log("Could not get group")

		return
	}

	require.NoError(t, err)
	require.Equal(t, groupId, group.Id)
	require.Equal(t, "Alcohol", group.Name)
	require.Equal(t, 0, group.ParentId)
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
	unit, err := dbModels.unit.GetById(unitId)
	if unit == nil {
		t.Fail()
		t.Log("Could not get unit")

		return
	}

	require.NoError(t, err)
	require.Equal(t, unitId, unit.Id)
	require.Equal(t, "ml", unit.Name)
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
	unit, err := dbModels.productVisibility.GetById(visibilityId)
	if unit == nil {
		t.Fail()
		t.Log("Could not get product visibility")

		return
	}

	require.NoError(t, err)
	require.Equal(t, visibilityId, unit.Id)
	require.Equal(t, 1, unit.CategoryId)
	require.Equal(t, 1, unit.ProductId)
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
	location, err := dbModels.location.GetById(locationId)
	if location == nil {
		t.Fail()
		t.Log("Could not get location")

		return
	}

	require.NoError(t, err)
	require.Equal(t, locationId, location.Id)
	require.Equal(t, "Bermuda Triangle", location.Name)
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
	vat, err := dbModels.vat.GetById(vatId)
	if vat == nil {
		t.Fail()
		t.Log("Could not get vat")

		return
	}

	require.NoError(t, err)
	require.Equal(t, vatId, vat.Id)
	require.Equal(t, 19, vat.Rate)
}
