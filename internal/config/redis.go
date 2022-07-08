package config

type (
	Redis struct {
		Maxretries   uint64
		Minidleconns uint64
		Poolsize     uint64
		Read         RedisRead  `ini:"Redis.Read"`
		Write        RedisWrite `ini:"Redis.Write"`
	}

	RedisRead struct {
		Address  string
		Password string
		Database uint
	}

	RedisWrite struct {
		Address  string
		Password string
		Database uint
	}
)
