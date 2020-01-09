package config

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/goextension/extmap"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

// Config ...
type Config struct {
	Name          string
	Task          int
	Broker        string
	QueueName     string
	ResultBackend string
	AMQP          *config.AMQPConfig
}

var _config Config

// DefaultConfigName ...
var DefaultConfigName = "config"

// DefaultConfigPath ...
var DefaultConfigPath = "."

// DefaultConfigType ...
var DefaultConfigType = "json"

// Default ...
func Default() *Config {
	return &Config{
		Name:          "cup",
		Task:          1,
		Broker:        "amqp://guest:guest@localhost:5672/",
		QueueName:     "machinery_task",
		ResultBackend: "redis://localhost:6379",
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery_exchange",
			ExchangeType: "direct",
			BindingKey:   "machinery_task",
		},
	}
}

// Get ...
func Get() *Config {
	return &_config
}

// Set ...
func Set(config *Config) {
	_config = *config
}

// SaveJSON ...
func SaveJSON() (err error) {
	fmt.Printf("Config:%+v\n", _config)
	indent, err := json.MarshalIndent(_config, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(DefaultConfigPath, DefaultConfigName), indent, 0755)
}

// SaveConfig ...
func SaveConfig() (err error) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	viper.SetConfigType("json")
	err = viper.MergeConfigMap(extmap.StructToMap(_config))
	if err != nil {
		return err
	}
	viper.SetConfigFile(filepath.Join(DefaultConfigPath, DefaultConfigName+".json"))

	return viper.WriteConfig()

}

// LoadConfig ...
func LoadConfig() (err error) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)

	err = viper.MergeInConfig()
	if err != nil {
		return err
	}
	m := extmap.ToMap(viper.AllSettings())
	return m.Struct(&_config)
}
