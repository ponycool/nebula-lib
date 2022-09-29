package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Trans struct {
	*sql.Tx
	DB            *DB
	TransactionID string
}

// Commit commits the transaction.
func (tx *Trans) Commit() error {

	if tx.DB.LogSQL {
		tx.DB.logger.Info(fmt.Sprintf("Executing (%s): COMMIT;", tx.TransactionID))

	}
	return tx.Tx.Commit()
}

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (tx *Trans) Exec(query string, args ...interface{}) (sql.Result, error) {

	return tx.ExecContext(context.Background(), query, args...)
}

// ExecContext executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (tx *Trans) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx.DB.LogSQL {
		tx.DB.logger.Info(fmt.Sprintf("Executing (%s):%s", tx.TransactionID, SQL{}.Format(query, args...)))
	}

	return tx.Tx.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (tx *Trans) QueryContext(ctx context.Context, query string, args ...interface{}) *RowsResult {
	if tx.DB.LogSQL {
		tx.DB.logger.Info(fmt.Sprintf("Query (%s):%s", tx.TransactionID, SQL{}.Format(query, args...)))
	}
	rs, err := tx.Tx.QueryContext(ctx, query, args...)
	return &RowsResult{rs, err}
}

// Query executes a query that returns rows, typically a SELECT.
func (tx *Trans) Query(query string, args ...interface{}) *RowsResult {
	return tx.QueryContext(context.Background(), query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (tx *Trans) QueryRow(query string, args ...interface{}) *RowResult {

	return tx.QueryRowContext(context.Background(), query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (tx *Trans) QueryRowContext(ctx context.Context, query string, args ...interface{}) *RowResult {
	if tx.DB.LogSQL {
		tx.DB.logger.Info(fmt.Sprintf("Query (%s):%s", tx.TransactionID, SQL{}.Format(query, args...)))
	}
	rows, err := tx.Tx.QueryContext(ctx, query, args...)

	return &RowResult{rows: rows, LastError: err}
}

// Rollback aborts the transaction.
func (tx *Trans) Rollback() error {
	err := tx.Tx.Rollback()
	if err != nil && err == sql.ErrTxDone {
		return err
	}
	if tx.DB.LogSQL {
		tx.DB.logger.Info(fmt.Sprintf("Executing (%s): ROLLBACK", tx.TransactionID))
	}
	return nil
}
