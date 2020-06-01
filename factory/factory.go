package factory

import (
	"database/sql"
	"github.com/jfeng45/gtransaction/config"
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
	"log"
)

// Build returns the SqlGdbc interface. This is the interface that you can use directly in your persistence layer
// If you don't need to cache sql.DB connection, you can call this function because you won't be able to get the sql.DB
// in SqlGdbc interface (if you need to do it, call BuildSqlDB()
func Build(dsc *config.DatabaseConfig) (gdbc.SqlGdbc, error) {
	db, err := sql.Open(dsc.DriverName, dsc.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt, err := buildGdbc(db, dsc)
	if err != nil {
		return nil, err
	}
	return dt, nil

}
// BuildSqlDB returns the sql.DB. The calling function need to generate corresponding gdbc.SqlGdbc struct based on
// sql.DB in order to use it in your persistence layer
// If you need to cache sql.DB connection, you need to call this function
func BuildSqlDB(dsc *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(dsc.DriverName, dsc.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return db, nil

}
func buildGdbc(sdb *sql.DB,dsc *config.DatabaseConfig) (gdbc.SqlGdbc, error){
	var sdt gdbc.SqlGdbc
	if dsc.Tx {
		tx, err := sdb.Begin()
		if err != nil {
			return nil, err
		}
		sdt = &gdbc.SqlConnTx{DB: tx}
		log.Println("buildGdbc(), create TX:")
	} else {
		sdt = &gdbc.SqlDBTx{sdb}
		log.Println("buildGdbc(), create DB:")
	}
	return sdt, nil
}