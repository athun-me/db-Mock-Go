package testingg_test

import (
	"testing"

	"athunlal/testingg"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	// Start mocking the database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected insert query and result
	expectedInsert := "INSERT INTO todo VALUES\\('test', '60'\\)"

	mock.ExpectExec(expectedInsert).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function being tested
	testingg.CreateTask(db, "test", "50")

	// Verify that the expected queries were executed
	assert.NoError(t, mock.ExpectationsWereMet())
}
