package test_database

import (
	"context"
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
	"os"
	"strings"
	"testing"
	"time"
)

var dbPool *database.DB

type modelStructs struct {
	account           *models.AccountModel
	accountOption     *models.AccountOptionModel
	category          *models.CategoryModel
	product           *models.ProductModel
	productGroup      *models.ProductGroupModel
	unit              *models.UnitModel
	productVisibility *models.ProductVisibilityModel
	location          *models.LocationModel
}

// run sets up members and the database before executing tests and tearing them down after execution.
func run(m *testing.M, dbModels *modelStructs) (code int, err error) {
	currentWorkDirectory, _ := os.Getwd()
	dbConf := config.LoadDbConfigFromFileEnv(string(currentWorkDirectory) + `/.env`)

	dbPool, err = database.New(dbConf.DbUser, dbConf.DbPassword, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)
	if err != nil {
		panic(err)
	}

	setUp()

	dbModels.account = &models.AccountModel{DB: dbPool}
	dbModels.accountOption = &models.AccountOptionModel{DB: dbPool}
	dbModels.category = &models.CategoryModel{DB: dbPool}
	dbModels.product = &models.ProductModel{DB: dbPool}
	dbModels.productGroup = &models.ProductGroupModel{DB: dbPool}
	dbModels.unit = &models.UnitModel{DB: dbPool}
	dbModels.productVisibility = &models.ProductVisibilityModel{DB: dbPool}
	dbModels.location = &models.LocationModel{DB: dbPool}

	defer tearDown()

	return m.Run(), nil
}

// getQueriesFromFile reads the file content specified in `filePath` into a string. It expects the content
// to be SQL queries that are terminated by `;` and returns a list of queries.
func getQueriesFromFile(filePath string) []string {
	// open file and read into string
	buf, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	s := string(buf)

	// separate queries
	s = strings.ReplaceAll(s, "\n", "")
	queries := strings.Split(s, ";")
	if queries[len(queries)-1] == "" {
		queries = queries[:len(queries)-1]
	}
	return queries
}

// batchExecQueries executes a list of queries.
func batchExecQueries(queries []string) pgx.BatchResults {
	// batch queries
	batch := &pgx.Batch{}
	for _, query := range queries {
		batch.Queue(query)
	}

	// execute queries
	ctx := context.Background()
	br := dbPool.SendBatch(ctx, batch)
	_, err := br.Exec()
	if err != nil {
		panic(err)
	}

	return br
}

// setUp set-ups the database by creating the tables and inserting test data.
func setUp() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/test_data/create_tables.sql`
	queries := getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Created tables")

	filePath = currentWorkDirectory + `/test_data/insert_test_data.sql`
	queries = getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Inserted test data")
}

// tearDown drops all tables from the database and closes the database connection pool.
func tearDown() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/test_data/drop_tables.sql`
	queries := getQueriesFromFile(filePath)

	br := batchExecQueries(queries)

	err := br.Close()
	if err != nil {
		panic(err)
	}

	logTearDown("Dropped tables")

	// TODO: this is a workaround to close the dbPool as it is stuck in an endless loop
	done := make(chan struct{})
	go func() {
		dbPool.Close()
		close(done)
	}()

	select {
	case <-done:
		logTearDown("Database closed")
	case <-time.After(5 * time.Second):
		logTearDown("Database close timed out")
	}
}

// logSetup prints a message to stdout. It is intended to be called during test setup.
func logSetup(msg string) {
	fmt.Printf("== SETUP: %s\n", msg)
}

// logTearDown prints a message to stdout. It is intended to be called during test teardown.
func logTearDown(msg string) {
	fmt.Printf("== TEARDOWN: %s\n", msg)
}
