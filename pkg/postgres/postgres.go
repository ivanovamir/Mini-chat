package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PostgresDBConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `env:"POSTGRES_PASSWORD"`
	Database        string `yaml:"database"`
	SslMode         string `yaml:"ssl_mode"`
	MaxConn         int    `yaml:"max_conn"`
	MaxConnAttempts int    `yaml:"max_conn_attempts"`
	MaxConnDelay    int    `yaml:"max_conn_delay"`
}

func NewPostgresDB(ctx context.Context, cfg *PostgresDBConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SslMode)

	pgxCfg, parseConfigErr := pgxpool.ParseConfig(dsn)

	if parseConfigErr != nil {
		return nil, parseConfigErr
	}

	//pgxCfg.MaxConns = int32(cfg.MaxConn)
	type PostgresDBConfig struct {
		Host            string `yaml:"host"`
		Port            string `yaml:"port"`
		User            string `yaml:"user"`
		Password        string `env:"POSTGRES_PASSWORD"`
		Database        string `yaml:"database"`
		SslMode         string `yaml:"ssl_mode"`
		MaxConn         int    `yaml:"max_conn"`
		MaxConnAttempts int    `yaml:"max_conn_attempts"`
		MaxConnDelay    int    `yaml:"max_conn_delay"`
	}
	pool, parseConfigErr = pgxpool.NewWithConfig(ctx, pgxCfg)

	if parseConfigErr != nil {
		return nil, parseConfigErr
	}

	if err = DoWithAttempts(func() error {
		return pool.Ping(ctx)
	}, cfg.MaxConnAttempts, time.Duration(cfg.MaxConnDelay)*time.Second); err != nil {
		return nil, err
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		return nil
	}
	return err
}

//func MigrateDB(db *sqlx.DB, cfg config.PostgresConfig) {
//	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
//	if err != nil {
//		log.Fatalf("couldn't get database instance for running migrations; %s", err.Error())
//	}
//
//	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:%s", cfg.Migrations), cfg.Name, driver)
//	if err != nil {
//		log.Fatalf("couldn't create migrate instance; %s", err.Error())
//	}
//
//	if err = m.Up(); err != nil {
//		if err.Error() == "no change" {
//			log.Info("database migration is not required, no changes")
//		} else {
//			log.Fatalf("couldn't run database migrations: %s", err.Error())
//		}
//	} else {
//		log.Info("database migration was run successfully")
//	}
//}
