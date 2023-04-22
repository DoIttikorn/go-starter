package sql

import (
	"database/sql"
	"testing"
)

type MockDB struct {
	query        string
	lastInsertId int64
	rowsAffected int64
}

// สร้างเพื่อให้สามารถใช้ method Exec ได้ เพราะของที่ return เป็น sql.Result ซึ่งมี method LastInsertId
func (m *MockDB) LastInsertId() (int64, error) {
	return m.lastInsertId, nil
}

// สร้างเพื่อให้สามารถใช้ method Exec ได้ เพราะของที่ return เป็น sql.Result ซึ่งมี method RowsAffected
func (m *MockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

// เราต้องมี method Exec เพื่อ implement interface DB
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &MockDB{
		rowsAffected: 1,
	}

	r, _ := execQuery(mock, "SELECT * FROM users")

	if mock.query != "SELECT * FROM users" {
		t.Errorf("should have been call db.Exec with query but it not.")
	}

	if r != 1 {
		t.Errorf("should return row effect %d but it got %d", 1, r)
	}

}
