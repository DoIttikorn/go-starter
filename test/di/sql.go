package sql

import "database/sql"

// สร้าง interface ชื่อ DB ที่มี method Exec เพื่อทำ dependency injection แทนการใช้งานตรงๆ กับ package sql
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func execQuery(db DB, query string, args ...interface{}) (int64, error) {
	res, err := db.Exec(query, args...)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
