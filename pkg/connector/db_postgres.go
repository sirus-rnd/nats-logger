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
}

// DefaultPostgresConfig is default postgres config
var DefaultPostgresConfig = &PostgresConfig{
	User:     "postgres",
	Password: "postgres",
	Host:     "localhost",
	Port:     5432,
	Database: "nats-logger",
}

// ConnectToPostgres will create connection instance to postgre instance
// and migrate models to this connection
func ConnectToPostgres(conf *PostgresConfig, models []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			conf.Host, conf.Port, conf.User, conf.Database, conf.Password,
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
