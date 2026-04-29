// Package repo implements the data access layer for the AVH Booking System.
//
// Each entity has a corresponding *Model type (e.g. [AccountModel]) that holds
// a database handle and exposes typed methods for querying and mutating that
// entity's table. Most models drive list queries through the generic [Query]
// struct — table name, filters, sort, and limit — which buildSelectSQL
// translates into a parameterised SQL statement. Single-row lookups delegate to
// the same path via a filter on the primary key column. Entities with compound
// primary keys (e.g. [AccountOptionModel]) use hand-written queries instead.
//
// Column identifiers supplied by callers are quoted with pgx.Identifier to
// prevent SQL injection; filter values are always passed as query parameters.
// Invalid or non-existent column names are rejected by PostgreSQL at query time.
package repo
