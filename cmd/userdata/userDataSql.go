// Package sql represents SQL database implementation of the user data persistence layer
package userdata

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	// use it to test transaction rollback
	//DELETE_USER        string = "delete from userinf where username=?"
	DELETE_USER        string = "delete from userinfo where username=?"
	QUERY_USER                = "SELECT * FROM userinfo "
	UPDATE_USER               = "update userinfo set username=?, department=?, created=? where uid=?"
	INSERT_USER               = "INSERT userinfo SET username=?,department=?,created=?"

	FORMAT_ISO8601_DATE                 = "2006-01-02"
)

// UserDataSql is the SQL implementation of UserDataInterface
type UserDataSql struct {
	DB gdbc.SqlGdbc
}

func (uds *UserDataSql) Remove(username string) (int64, error) {

	stmt, err := uds.DB.Prepare(DELETE_USER)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()

	res, err := stmt.Exec(username)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	log.Println("remove:row affected ", rowsAffected)
	return rowsAffected, nil
}

func (uds *UserDataSql) FindAll() error {

	rows, err := uds.DB.Query(QUERY_USER)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer rows.Close()
	var id int
	var name, department,dateString string
	for rows.Next() {
		err = rows.Scan(&id, &name, &department, &dateString)
		created, err := time.Parse(FORMAT_ISO8601_DATE, dateString)
		if err != nil {
			log.Printf("date err:", err)
		}
		log.Printf("id=%v,name= %v,department= %v, date:= %v", id, name, department, created)

	}
	//need to check error for rows.Next()
	if err = rows.Err(); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func (uds *UserDataSql) Update(name string, department string, created time.Time, id int64) (rowsAffected int64, e error) {

	stmt, err := uds.DB.Prepare(UPDATE_USER)

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, department, created, id)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err = res.RowsAffected()

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	log.Println("update: rows affected: ", rowsAffected)

	return rowsAffected, nil
}

func (uds *UserDataSql) Insert(name string, department string, created time.Time) (id int64, e error) {

	stmt, err := uds.DB.Prepare(INSERT_USER)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, department, created)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	newId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	log.Println("user inserted and id=:", newId)
	return newId, nil
}

func (uds *UserDataSql) EnableTx(txFunc func() error) error {
	return uds.DB.TxEnd(txFunc)
}

