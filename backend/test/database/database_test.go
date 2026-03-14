package test_database

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	id, err := dbModels.account.Insert(context.Background(), dummyAccount)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAccountById(t *testing.T) {
	const accountID = 1
	daGama, err := dbModels.account.GetByID(context.Background(), accountID)
	if daGama == nil {
		t.Fail()
		t.Log("Could not get account")

		return
	}

	require.NoError(t, err)
	require.Equal(t, accountID, daGama.ID)
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
	accountID, key, err := dbModels.accountOption.Insert(context.Background(), dummyOpt)

	require.NoError(t, err)
	assert.NotZero(t, accountID)
	assert.NotZero(t, key)
}

func TestGetAccountOptionByAccountAndKey(t *testing.T) {
	const accountID = 1
	const optKey = "deceased"
	opt, err := dbModels.accountOption.Get(context.Background(), accountID, optKey)
	if opt == nil {
		t.Fail()
		t.Log("Could not get account option")

		return
	}

	require.NoError(t, err)
	require.Equal(t, accountID, opt.AccountID)
	require.Equal(t, optKey, opt.Key)
	require.Equal(t, "true", opt.Value)
}

// --------------------------------------------------
// Category
// --------------------------------------------------

func TestInsertCategory(t *testing.T) {
	dummyCategory := models.CreateCategory("Guest", "user-friends", "account")
	id, err := dbModels.category.Insert(context.Background(), dummyCategory)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetCategoryById(t *testing.T) {
	const categoryID = 1
	cat, err := dbModels.category.GetByID(context.Background(), categoryID)
	if cat == nil {
		t.Fail()
		t.Log("Could not get category")

		return
	}

	require.NoError(t, err)
	require.Equal(t, categoryID, cat.ID)
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
	id, err := dbModels.product.Insert(context.Background(), dummyProduct)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductById(t *testing.T) {
	const productID = 1
	product, err := dbModels.product.GetByID(context.Background(), productID)
	if product == nil {
		t.Fail()
		t.Log("Could not get product")

		return
	}

	require.NoError(t, err)
	require.Equal(t, product.ID, productID)
	require.Equal(t, "Rota das Especiarias", product.Name)
	require.Equal(t, 1800, product.Price)
	require.Equal(t, 1, product.VatID)
	require.Equal(t, 1, product.ProductGroupID)
	require.Equal(t, 150, product.Size)
	require.Equal(t, 1, product.UnitID)
	require.Equal(t, 2, product.CategoryID)
}

// --------------------------------------------------
// ProductGroup
// --------------------------------------------------

func TestInsertProductGroup(t *testing.T) {
	dummyProductGroup := models.CreateProductGroup("Beer", 1)
	id, err := dbModels.productGroup.Insert(context.Background(), dummyProductGroup)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetProductGroupById(t *testing.T) {
	const groupID = 1
	group, err := dbModels.productGroup.GetByID(context.Background(), groupID)
	if group == nil {
		t.Fail()
		t.Log("Could not get group")

		return
	}

	require.NoError(t, err)
	require.Equal(t, groupID, group.ID)
	require.Equal(t, "Alcohol", group.Name)
	require.Equal(t, 0, group.ParentID)
}

// --------------------------------------------------
// Unit
// --------------------------------------------------

func TestInsertUnit(t *testing.T) {
	dummyUnit := models.CreateUnit("oz")
	id, err := dbModels.unit.Insert(context.Background(), dummyUnit)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetUnitById(t *testing.T) {
	const unitID = 1
	unit, err := dbModels.unit.GetByID(context.Background(), unitID)
	if unit == nil {
		t.Fail()
		t.Log("Could not get unit")

		return
	}

	require.NoError(t, err)
	require.Equal(t, unitID, unit.ID)
	require.Equal(t, "ml", unit.Name)
}

// --------------------------------------------------
// ProductVisibility
// --------------------------------------------------

func TestInsertProductVisibility(t *testing.T) {
	dummyVisibility := models.CreateProductVisibility(3, 1, 3)
	id, err := dbModels.productVisibility.Insert(context.Background(), dummyVisibility)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestProductVisibilityById(t *testing.T) {
	const visibilityID = 1
	unit, err := dbModels.productVisibility.GetByID(context.Background(), visibilityID)
	if unit == nil {
		t.Fail()
		t.Log("Could not get product visibility")

		return
	}

	require.NoError(t, err)
	require.Equal(t, visibilityID, unit.ID)
	require.Equal(t, 1, unit.CategoryID)
	require.Equal(t, 1, unit.ProductID)
}

// --------------------------------------------------
// Location
// --------------------------------------------------

func TestInsertLocation(t *testing.T) {
	dummyLocation := models.CreateLocation("Nursing Home")
	id, err := dbModels.location.Insert(context.Background(), dummyLocation)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetLocationById(t *testing.T) {
	const locationID = 1
	location, err := dbModels.location.GetByID(context.Background(), locationID)
	if location == nil {
		t.Fail()
		t.Log("Could not get location")

		return
	}

	require.NoError(t, err)
	require.Equal(t, locationID, location.ID)
	require.Equal(t, "Bermuda Triangle", location.Name)
}

// --------------------------------------------------
// Vat
// --------------------------------------------------

func TestInsertVat(t *testing.T) {
	dummyVat := models.CreateVat(12)
	id, err := dbModels.vat.Insert(context.Background(), dummyVat)

	require.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetVatById(t *testing.T) {
	const vatID = 1
	vat, err := dbModels.vat.GetByID(context.Background(), vatID)
	if vat == nil {
		t.Fail()
		t.Log("Could not get vat")

		return
	}

	require.NoError(t, err)
	require.Equal(t, vatID, vat.ID)
	require.Equal(t, 19, vat.Rate)
}
