package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetSaleSummary(t *testing.T) {
	// Init mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := NewSaleSummaryRepository(sqlxDB)
	ctx := context.Background()

	query := `SELECT * FROM sales_summary ORDER BY date DESC LIMIT 5`

	// Test case 1: Success
	t.Run("Success", func(t *testing.T) {
		// mock data
		date1, _ := time.Parse(time.RFC3339, "2023-03-18T22:32:25Z")
		date2, _ := time.Parse(time.RFC3339, "2023-09-03T13:50:20Z")
		changePct1 := decimal.NewFromFloat(61.51)
		changePct2 := decimal.NewFromFloat(-2.28)
		rows := sqlmock.NewRows([]string{"sales_summary_id", "total_value", "change_percentage", "date"}).
			AddRow("9234a776-e6ac-46e2-bc24-c959ce216751", 4754106.83, 61.51, date1).
			AddRow("e5648831-7d0e-4ef5-8e04-f6e6a0eaafb1", 1512948.97, -2.28, date2)

		// Set expectation
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		// Call function
		result, err := repo.GetSaleSummary(ctx)

		// Check result
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "9234a776-e6ac-46e2-bc24-c959ce216751", result[0].SalesSummaryID)
		assert.True(t, decimal.NewFromFloat(4754106.83).Equal(result[0].TotalValue))
		assert.True(t, changePct1.Equal(*result[0].ChangePercentage))
		assert.Equal(t, date1, result[0].Date)
		assert.Equal(t, "e5648831-7d0e-4ef5-8e04-f6e6a0eaafb1", result[1].SalesSummaryID)
		assert.True(t, decimal.NewFromFloat(1512948.97).Equal(result[1].TotalValue))
		assert.True(t, changePct2.Equal(*result[1].ChangePercentage))
		assert.Equal(t, date2, result[1].Date)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Test case 2: Query error
	t.Run("QueryError", func(t *testing.T) {
		// Set expectation
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("database connection error"))

		// Call function
		result, err := repo.GetSaleSummary(ctx)

		// Check result
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "query SalesSummary err: database connection error")
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Test case 3: Scan data error
	t.Run("ScanError", func(t *testing.T) {
		// Mock data with invalid data type
		date, _ := time.Parse(time.RFC3339, "2025-05-01T10:00:00Z")
		rows := sqlmock.NewRows([]string{"sales_summary_id", "total_value", "change_percentage", "date"}).
			AddRow("SS001", "invalid_decimal", 5.25, date) // totalValue không phải số

		// set expectation
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		// Call function
		result, err := repo.GetSaleSummary(ctx)

		// Check result
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "StructScan SalesSummary err")
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Test case 4: Not have data
	t.Run("NoData", func(t *testing.T) {
		// Empty mock data
		rows := sqlmock.NewRows([]string{"sales_summary_id", "total_value", "change_percentage", "date"})

		// Set expectation
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		result, err := repo.GetSaleSummary(ctx)

		assert.NoError(t, err)
		assert.Empty(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Test case 5: ChangePercentage is null
	t.Run("NullChangePercentage", func(t *testing.T) {
		// Mock data with changePercentage is null
		date, _ := time.Parse(time.RFC3339, "2025-05-01T10:00:00Z")
		rows := sqlmock.NewRows([]string{"sales_summary_id", "total_value", "change_percentage", "date"}).
			AddRow("SS001", 1000.50, nil, date)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		result, err := repo.GetSaleSummary(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "SS001", result[0].SalesSummaryID)
		assert.True(t, decimal.NewFromFloat(1000.50).Equal(result[0].TotalValue))
		assert.Nil(t, result[0].ChangePercentage)
		assert.Equal(t, date, result[0].Date)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
