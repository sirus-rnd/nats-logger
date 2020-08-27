package cmd

import (
	"encoding/json"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"go.sirus.dev/nats-logger/pkg/connector"

	"github.com/fatih/structs"
)

// Config define service configuration structure
type Config struct {
	LogLevel string                    `mapstructure:"log_level"`
	Postgres *connector.PostgresConfig `mapstructure:"postgres"`
	NatsURL  string                    `mapstructure:"nats_url"`
}

// DefaultConfig is default configuration
var DefaultConfig = Config{
	LogLevel: "info",
	Postgres: connector.DefaultPostgresConfig,
	NatsURL:  nats.DefaultOptions.Url,
}

// String implement string interface
func (c *Config) String() string {
	val, _ := json.Marshal(c)
	return string(val)
}

var conf = viper.New()
var keys []string

// getEnvKeys will read environment keys
func getEnvKeys(fields []*structs.Field, prefix string) {
	for _, field := range fields {
		key := field.Tag("mapstructure")
		if prefix != "" {
			keys = append(keys, prefix+"."+key)
		} else {
			keys = append(keys, key)
		}
		if field.Kind().String() == "ptr" {
			if len(prefix) > 0 {
				key = prefix + "." + key
			}
			getEnvKeys(structs.Fields(field.Value()), key)
		}
	}
}

// LoadConfig will load configurations
func LoadConfig() (*Config, error) {
	// initiate config
	config := DefaultConfig
	// get all configurations keys
	fields := structs.Fields(config)
	getEnvKeys(fields, "")
	// read from config file
	conf.SetConfigName("config")
	conf.AddConfigPath("/etc/nats-logger/")
	conf.AddConfigPath("$HOME/.nats-logger")
	conf.AddConfigPath(".")
	conf.SetConfigType("yaml")
	_ = conf.ReadInConfig()
	// replace configurations using environment
	conf.SetEnvPrefix("nats-logger")
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// reflect default configs to get all keys
	for _, key := range keys {
		_ = conf.BindEnv(key)
		val := conf.Get(key)
		conf.Set(key, val)
	}
	// unmarshal configuration
	err := conf.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
