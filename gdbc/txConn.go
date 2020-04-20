package gdbc

import (
	"log"
)

// DB doesn't rollback, do nothing here
func (cdt *SqlDBTx) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (cdt *SqlDBTx) Commit() error {
	return nil
}

// DB doesnt rollback, do nothing here
func (cdt *SqlDBTx) TxEnd(txFunc func() error) error {
	return nil
}

func (sct *SqlConnTx) TxEnd(txFunc func() error) error {
	var err error
	tx := sct.DB

	defer func() {
		if p := recover(); p != nil {
			log.Println("found p and rollback:", p)
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			log.Println("found error and rollback:", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			log.Println("commit:")
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}


func (sct *SqlConnTx) Rollback() error {
	return sct.DB.Rollback()
}

func (sct *SqlConnTx) Commit() error {
	return sct.DB.Commit()
}
