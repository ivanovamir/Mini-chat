package configs

type PostgresDBConfig struct {
	Host            string `yml:"host"`
	Port            string `yml:"port"`
	User            string `yml:"user"`
	Password        string `env:"PG_PASSWORD"`
	Database        string `yml:"database"`
	SslMode         string `yml:"sslmode"`
	MaxConnAttempts int    `yml:"max_conn_attempts"`
	MaxConnDelay    int    `yml:"max_conn_delay"`
}
