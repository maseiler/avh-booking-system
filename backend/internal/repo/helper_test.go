package repo_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/av-huette/avh-booking-system/internal/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/repo"
	"github.com/jackc/pgx/v5"
)

var dbPool *database.DB

type modelStructs struct {
	account           *repo.AccountModel
	accountOption     *repo.AccountOptionModel
	category          *repo.CategoryModel
	product           *repo.ProductModel
	productGroup      *repo.ProductGroupModel
	unit              *repo.UnitModel
	productVisibility *repo.ProductVisibilityModel
	location          *repo.LocationModel
	vat               *repo.VatModel
}

// run sets up members and the database before executing tests and tearing them down after execution.
func run(m *testing.M, dbModels *modelStructs) (code int, err error) {
	currentWorkDirectory, _ := os.Getwd()
	if err := config.LoadEnvFromFile(currentWorkDirectory + `/.env`); err != nil {
		panic(err)
	}
	dbConf, err := config.LoadDBConfig()
	if err != nil {
		panic(err)
	}

	dbPool, err = database.New(dbConf.DBUser, dbConf.DBPassword, dbConf.DBHost, dbConf.DBPort, dbConf.DBName)
	if err != nil {
		panic(err)
	}

	setUp()

	dbModels.account = &repo.AccountModel{DB: dbPool}
	dbModels.accountOption = &repo.AccountOptionModel{DB: dbPool}
	dbModels.category = &repo.CategoryModel{DB: dbPool}
	dbModels.product = &repo.ProductModel{DB: dbPool}
	dbModels.productGroup = &repo.ProductGroupModel{DB: dbPool}
	dbModels.unit = &repo.UnitModel{DB: dbPool}
	dbModels.productVisibility = &repo.ProductVisibilityModel{DB: dbPool}
	dbModels.location = &repo.LocationModel{DB: dbPool}
	dbModels.vat = &repo.VatModel{DB: dbPool}

	defer tearDown()

	return m.Run(), nil
}

// getQueriesFromFile reads the file content specified in `filePath` into a string. It expects the content
// to be SQL queries that are terminated by `;` and returns a list of queries.
func getQueriesFromFile(filePath string) []string {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	s := string(buf)

	s = strings.ReplaceAll(s, "\n", "")
	queries := strings.Split(s, ";")
	if queries[len(queries)-1] == "" {
		queries = queries[:len(queries)-1]
	}
	return queries
}

// batchExecQueries executes a list of queries.
func batchExecQueries(queries []string) pgx.BatchResults {
	batch := &pgx.Batch{}
	for _, query := range queries {
		batch.Queue(query)
	}

	ctx := context.Background()
	br := dbPool.SendBatch(ctx, batch)
	_, err := br.Exec()
	if err != nil {
		panic(err)
	}

	return br
}

// setUp creates tables and inserts test data.
func setUp() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/testdata/create_tables.sql`
	queries := getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Created tables")

	filePath = currentWorkDirectory + `/testdata/insert_test_data.sql`
	queries = getQueriesFromFile(filePath)
	batchExecQueries(queries)
	logSetup("Inserted test data")
}

// tearDown drops all tables and closes the database connection pool.
func tearDown() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/testdata/drop_tables.sql`
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

func logSetup(msg string) {
	fmt.Printf("== SETUP: %s\n", msg)
}

func logTearDown(msg string) {
	fmt.Printf("== TEARDOWN: %s\n", msg)
}
