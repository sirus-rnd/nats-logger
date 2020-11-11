package connector

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresConfig information
type PostgresConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl_mode"`
	SSLCA    string `mapstructure:"ssl_ca"`
}

// DefaultPostgresConfig is default postgres config
var DefaultPostgresConfig = &PostgresConfig{
	User:     "postgres",
	Password: "postgres",
	Host:     "localhost",
	Port:     5432,
	Database: "nats-logger",
	SSLMode:  "disable",
	SSLCA:    "",
}

// ConnectToPostgres will create connection instance to postgre instance
// and migrate models to this connection
func ConnectToPostgres(conf *PostgresConfig, models []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		// for full parameters check https://www.postgresql.org/docs/11/libpq-connect.html#LIBPQ-PARAMKEYWORDS
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s sslrootcert=%s",
			conf.Host, conf.Port, conf.User, conf.Database, conf.Password, conf.SSLMode, conf.SSLCA,
		),
	)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}
	return db, nil
}
