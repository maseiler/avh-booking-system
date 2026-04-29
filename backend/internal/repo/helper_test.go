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
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

var dbPool *database.DB

// run sets up the database before executing tests and tears it down after.
func run(m *testing.M) (code int, err error) {
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
	defer tearDown()

	return m.Run(), nil
}

// beginTx starts a transaction and registers a rollback via t.Cleanup,
// leaving the database in its seeded state after each test.
func beginTx(t *testing.T) pgx.Tx {
	t.Helper()
	tx, err := dbPool.Begin(context.Background())
	require.NoError(t, err)
	t.Cleanup(func() { _ = tx.Rollback(context.Background()) })
	return tx
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

// execQueries executes a list of queries sequentially, panicking if any one fails.
func execQueries(queries []string) {
	ctx := context.Background()
	for _, query := range queries {
		if _, err := dbPool.Exec(ctx, query); err != nil {
			panic(err)
		}
	}
}

// setUp creates tables and inserts test data.
func setUp() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/testdata/create_tables.sql`
	queries := getQueriesFromFile(filePath)
	execQueries(queries)
	logSetup("Created tables")

	filePath = currentWorkDirectory + `/testdata/insert_test_data.sql`
	queries = getQueriesFromFile(filePath)
	execQueries(queries)
	logSetup("Inserted test data")
}

// tearDown drops all tables and closes the database connection pool.
func tearDown() {
	currentWorkDirectory, _ := os.Getwd()
	filePath := currentWorkDirectory + `/testdata/drop_tables.sql`
	queries := getQueriesFromFile(filePath)

	execQueries(queries)
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
