// Package txdataservice represents transaction support on data service layer
package txdataservice

// TxDataInterface represents operations needed for transaction support.
// It only needs to be implemented once for each database
type TxDataInterface interface {
	// EnableTx is called at the end of a transaction and based on whether there is an error, it commits or rollback the
	// transaction.
	// txFunc is the business function wrapped in a transaction
	EnableTx(txFunc func() error) error
}

