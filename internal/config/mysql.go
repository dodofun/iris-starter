package config

type (
	Mysql struct {
		Connmaxlifetime uint64
		Maxidleconn     uint64
		Maxopenconn     uint64
		Read            MysqlRead  `ini:"Mysql.Read"`
		Write           MysqlWrite `ini:"Mysql.Write"`
	}

	MysqlRead struct {
		Address  string
		Port     string
		User     string
		Password string
		Database string
	}

	MysqlWrite struct {
		Address  string
		Port     string
		User     string
		Password string
		Database string
	}
)
