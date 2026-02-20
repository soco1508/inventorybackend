package main

import (
	"backend/config"
	"backend/internal/db/models"
	"backend/pkg/db"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//D:\Go_Project\Inventory\backend>insert-by-table.exe -input seedData/users.json -table Users

func main() {
	ctx := context.Background()

	config, err := config.NewParsedConfig()
	if err != nil {
		log.Fatalf("cannot get config, err:\n %+v", err)
	}

	dbConfig := db.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Name:     config.Database.Name,
	}

	pgxDb, err := db.PgxInitDB(ctx, dbConfig)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer pgxDb.Close(ctx)

	inputFile := flag.String("input", "example.json", "Path to input JSON file")
	tableName := flag.String("table", "", "Table name")
	flag.Parse()

	jsonData, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %+v", err)
	}

	err = BatchInsertByTable(ctx, pgxDb, jsonData, *tableName)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fmt.Println("Data insertion completed")
}

func BatchInsertByTable(ctx context.Context, pgxDb *pgx.Conn, jsonData []byte, tableName string) error {
	switch tableName {
	case db.Users.String():
		users := []models.User{}
		if err := json.Unmarshal(jsonData, &users); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.Users.String(), users)
	case db.Products.String():
		products := []models.Product{}
		if err := json.Unmarshal(jsonData, &products); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.Products.String(), products)
	case db.Sales.String():
		sales := []models.Sale{}
		if err := json.Unmarshal(jsonData, &sales); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.Sales.String(), sales)
	case db.SalesSummary.String():
		salesSummary := []models.SaleSummary{}
		if err := json.Unmarshal(jsonData, &salesSummary); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.SalesSummary.String(), salesSummary)
	case db.Purchases.String():
		purchases := []models.Purchase{}
		if err := json.Unmarshal(jsonData, &purchases); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.Purchases.String(), purchases)
	case db.PurchaseSummary.String():
		purchaseSummary := []models.PurchaseSummary{}
		if err := json.Unmarshal(jsonData, &purchaseSummary); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.PurchaseSummary.String(), purchaseSummary)
	case db.Expenses.String():
		expenses := []models.Expense{}
		if err := json.Unmarshal(jsonData, &expenses); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.Expenses.String(), expenses)
	case db.ExpenseSummary.String():
		expenseSummary := []models.ExpenseSummary{}
		if err := json.Unmarshal(jsonData, &expenseSummary); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.ExpenseSummary.String(), expenseSummary)
	case db.ExpenseByCategory.String():
		expenseByCategory := []models.ExpenseByCategory{}
		if err := json.Unmarshal(jsonData, &expenseByCategory); err != nil {
			return fmt.Errorf("error parsing from json to struct: %+v", err)
		}
		return BulkInsertCopy(ctx, pgxDb, db.ExpenseByCategory.String(), expenseByCategory)
	default:
		return nil
	}
}

func BulkInsertCopy[T any](ctx context.Context, conn *pgx.Conn, tableName string, items []T) error {
	if len(items) == 0 {
		return nil
	}

	var columns []string //contain column names
	var rows [][]any     //contain data in each row

	t := reflect.TypeOf(items[0])

	//nếu t là pointer thì gán về struct để thao tác
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("BulkInsertCopy only supports struct slices")
	}

	//get and add column name from tag db to columns
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columns = append(columns, dbTag)
		}
	}

	//get and add data in each row to rows
	for _, item := range items {
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		var row []any
		for i := range t.NumField() {
			dbTag := t.Field(i).Tag.Get("db")
			if dbTag == "" || dbTag == "-" {
				continue
			}
			row = append(row, val.Field(i).Interface())
		}
		rows = append(rows, row)
	}

	_, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		columns,
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return fmt.Errorf("bulk insert %s err:\n %+v", tableName, err)
	}

	return nil
}
