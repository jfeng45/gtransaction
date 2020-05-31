package factory

import (
	"database/sql"
	"github.com/jfeng45/gtransaction/config"
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
	"log"
)

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
	dt, err := getGdbc(db, dsc)
	if err != nil {
		return nil, err
	}
	return dt, nil

}

func getGdbc(sdb *sql.DB,dsc *config.DatabaseConfig) (gdbc.SqlGdbc, error){
	var sdt gdbc.SqlGdbc
	if dsc.Tx {
		tx, err := sdb.Begin()
		if err != nil {
			return nil, err
		}
		sdt = &gdbc.SqlConnTx{DB: tx}
		log.Println("getGdbc(), create TX:")
	} else {
		sdt = &gdbc.SqlDBTx{sdb}
		log.Println("getGdbc(), create DB:")
	}
	return sdt, nil
}