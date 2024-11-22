package handler

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &HandlerImpl{
		DB: db,
	}

	testProductName := "celana jeans"
	testPrice := 700000.00
	testStock := 20

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO Products (ProductName, Price, Stock) VALUES ($1, $2, $3)")).
		WithArgs(testProductName, testPrice, testStock).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = handler.AddProduct(testProductName, testPrice, testStock)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddProduct_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &HandlerImpl{
		DB: db,
	}

	testProductName := "celana jeans"
	testPrice := 700000.00
	testStock := 20

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO Products (ProductName, Price, Stock) VALUES ($1, $2, $3)")).
		WithArgs(testProductName, testPrice, testStock).
		WillReturnError(fmt.Errorf("database connection lost"))

	err = handler.AddProduct(testProductName, testPrice, testStock)

	if err == nil {
		t.Error("expected an error but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &HandlerImpl{
		DB: db,
	}

	testProductId := 10
	testProductName := "celana jeans"
	testPrice := 700000.00
	testStock := 20

	mock.ExpectQuery(regexp.QuoteMeta("UPDATE Products SET ProductName = COALESCE(NULLIF($1, ''), ProductName), Price = COALESCE(NULLIF($2, 0), Price), Stock = COALESCE(NULLIF($3, 0), Stock) WHERE Id = $4 RETURNING Id;")).
		WithArgs(testProductName, testPrice, testStock, testProductId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testProductId))

	err = handler.UpdateProduct(testProductId, testProductName, testPrice, testStock)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateProduct_ProductIdNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &HandlerImpl{
		DB: db,
	}

	testProductId := 999
	testProductName := "celana jeans"
	testPrice := 700000.00
	testStock := 20

	mock.ExpectQuery(regexp.QuoteMeta("UPDATE Products SET ProductName = COALESCE(NULLIF($1, ''), ProductName), Price = COALESCE(NULLIF($2, 0), Price), Stock = COALESCE(NULLIF($3, 0), Stock) WHERE Id = $4 RETURNING Id;")).
		WithArgs(testProductName, testPrice, testStock, testProductId).
		WillReturnError(sql.ErrNoRows)

	err = handler.UpdateProduct(testProductId, testProductName, testPrice, testStock)

	expectedErr := fmt.Sprintf("product with id %d not found", testProductId)
	if err.Error() != expectedErr {
		t.Errorf("expected error '%s', got '%s'", expectedErr, err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
