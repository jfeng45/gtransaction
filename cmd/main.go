package main

import (
	"database/sql"
	"github.com/jfeng45/gtransaction/cmd/userdata"
	"github.com/jfeng45/gtransaction/config"
	"github.com/jfeng45/gtransaction/factory"
	"github.com/jfeng45/gtransaction/gdbc"
	"log"
	"time"
)

const (
	DRIVER_NAME       string = "mysql"
	DATA_SOURCE_NAME       string ="root:@tcp(localhost:4333)/service_config?charset=utf8"
	TX = true
)

func main() {
	tx :=true
	//tx := false
	dsc := config.DatabaseConfig{DRIVER_NAME, DATA_SOURCE_NAME, tx}
	//testTxSql(&dsc)
	//testSql(&dsc)
	testBuildSqlDB(&dsc)
}

func testBuildSqlDB(dsc *config.DatabaseConfig) {
	db, err :=factory.BuildSqlDB(dsc)
	if err != nil {
		log.Println("can't make connection:", err)
	}
	tx :=true
	gdbc, err :=buildGdbc(db,tx )
	uds := userdata.UserDataSql{gdbc}
	testModifyAndUnregisterWithTx(uds)
}

func testTxSql(dsc *config.DatabaseConfig) {
	g, err :=factory.Build(dsc)
	if err != nil {
		log.Println("can't make connection:", err)
	}
	uds := userdata.UserDataSql{g}
	testModifyAndUnregisterWithTx(uds)
}

func testSql(dsc *config.DatabaseConfig) {
	g, err :=factory.Build(dsc)
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
		log.Printf("date format err: %v\n", err)
	}
	name :="Brian"
	department :="marketing"

	id, err := uds.Insert(name, department,created)
	if err != nil {
		log.Printf("user registration failed: %v\n", err)
	}
	log.Printf("new user registered:id= %v", id)
}

func testUnregister(username string, uds userdata.UserDataSql) error{
	//username := "Brian"
	rowsAffected, err := uds.Remove(username)
	if err != nil {
		return err
	}
	log.Printf("testUnregister successully: rowsAffected: %v",rowsAffected )
	return nil
}

func testModifyUser(id int64, uds userdata.UserDataSql) error{

	created, err := time.Parse(userdata.FORMAT_ISO8601_DATE, "2019-12-01")
	if err != nil {
		log.Printf("date format err: %v\n", err)
	}
	name := "Aditi"
	department :="HR"
	//var id int64= 28
	rowsAffected, err := uds.Update(name, department,created, id)
	if err != nil {
		log.Printf("Modify user failed: %v\n", err)
		return err
	}
	log.Printf("user modified succeed:rowsAffected= %v", rowsAffected)
	return nil
}

func testModifyAndUnregister(uds userdata.UserDataSql) error {
	var id int64= 27
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
		log.Printf("ModifyAndUnregisterWithTx failed: %v", err)
	}
}

func buildGdbc(sdb *sql.DB,tx bool) (gdbc.SqlGdbc, error){
	var sdt gdbc.SqlGdbc
	if tx {
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

