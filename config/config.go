package config

type DatabaseConfig struct {
	// driver name for database
	DriverName string `yaml:"driverName"`
	// datasource name
	DataSourceName string `yaml:"dataSourceName"`
	// To indicate whether support transaction or not. "true" means supporting transaction
	Tx bool `yaml:"tx"`
}