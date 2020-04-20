package main

import (
	"database/sql"
	"github.com/jfeng45/gtransaction/cmd/userdata"
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	DB_DRIVER_NAME       string = "mysql"
	DB_SOURCE_NAME       string ="root:@tcp(localhost:4333)/service_config?charset=utf8"
)

func main() {
	tx :=true
	testTxSql(tx)
	//tx= false
	//testSql(tx)
}

func testTxSql(tx bool) {
	g, err :=initGdbc(tx)
	if err != nil {
		log.Println("can't make connection:", err)
	}
	uds := userdata.UserDataSql{g}
	testModifyAndUnregisterWithTx(uds)
}

func testSql(tx bool) {
	g, err :=initGdbc(tx)
	if err != nil {
		log.Println("can't make connection:", err)
	}
	uds := userdata.UserDataSql{g}
	//testListUser(uds)
	//testRegisterUser(uds)
	err = testModifyAndUnregister(uds)
	if err != nil {
		log.Println("testModifyAndUnregister failed:", err)
	}

}

func testListUser(uds userdata.UserDataSql) {
	err := uds.FindAll()
	if err != nil {
		log.Println("FindAll failed:", err)
	}
}

func testRegisterUser(uds userdata.UserDataSql) {
	created, err := time.Parse(userdata.FORMAT_ISO8601_DATE, "2018-12-09")
	if err != nil {
		log.Printf("date format err:%+v\n", err)
	}
	name :="Brian"
	department :="marketing"

	id, err := uds.Insert(name, department,created)
	if err != nil {
		log.Printf("user registration failed:%v\n", err)
	}
	log.Printf("new user registered:id=", id)
}

func testUnregister(username string, uds userdata.UserDataSql) error{
	//username := "Brian"
	rowsAffected, err := uds.Remove(username)
	if err != nil {
		return err
	}
	log.Printf("testUnregister successully: rowsAffected:",rowsAffected )
	return nil
}

func testModifyUser(id int64, uds userdata.UserDataSql) error{

	created, err := time.Parse(userdata.FORMAT_ISO8601_DATE, "2019-12-01")
	if err != nil {
		log.Printf("date format err:%+v\n", err)
	}
	name := "Aditi"
	department :="HR"
	//var id int64= 28
	rowsAffected, err := uds.Update(name, department,created, id)
	if err != nil {
		log.Printf("Modify user failed:%+v\n", err)
		return err
	}
	log.Printf("user modified succeed:rowsAffected=", rowsAffected)
	return nil
}

func testModifyAndUnregister(uds userdata.UserDataSql) error {
	var id int64= 30
	username := "Aditi"
	err := testModifyUser(id, uds)
	if err != nil {
		return err
	}
	err = testUnregister(username,uds)
	if err != nil {
		return err
	}
	return nil
}

func testModifyAndUnregisterWithTx(uds userdata.UserDataSql) {
	err := uds.EnableTx(func() error {
		// wrap the business function inside the TxEnd function
		return testModifyAndUnregister(uds)
	})
	if err != nil {
		log.Printf("ModifyAndUnregisterWithTx failed:", err)
	}
}

func initGdbc(tx bool) (gdbc.SqlGdbc,error) {

	db, err := sql.Open(DB_DRIVER_NAME, DB_SOURCE_NAME)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	var sqlConn gdbc.SqlGdbc
	if tx {
		tx, err := db.Begin()
		if err != nil {
			return nil, err
		}
		sqlConn = &gdbc.SqlConnTx{DB: tx}
		log.Printf("create TX:")
	} else {
		sqlConn = &gdbc.SqlDBTx{db}
		log.Printf("create DB:")
	}
	return sqlConn, nil
}


