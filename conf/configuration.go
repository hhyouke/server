package conf

// GlobalConfiguration global configuration
type GlobalConfiguration struct {
	API     APIConfiguration
	DB      DBConfiguration
	Logging LoggingConfiguration
}

// APIConfiguration api config
type APIConfiguration struct {
	Port      string
	Host      string
	MachineID string `mapstructure:"machine_id"`
}

//LoggingConfiguration logger config
type LoggingConfiguration struct {
	Level string
}

// DBConfiguration main db config
type DBConfiguration struct {
	URL         string
	Dialect     string
	LogMode     bool `mapstructure:"log_mode"`
	MaxOpenConn int  `mapstructure:"max_open_conn"`
	MaxIdelConn int  `mapstructure:"max_idle_conn"`
	AutoMigrate bool `mapstructure:"auto_migrate"`
}
