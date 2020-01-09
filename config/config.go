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
	Broker        string
	QueueName     string
	ResultBackend string
	AMQP          config.AMQPConfig
}

var _config Config

// DefaultConfigName ...
var DefaultConfigName = "config"

// DefaultConfigPath ...
var DefaultConfigPath = "."

// DefaultConfigType ...
var DefaultConfigType = "json"

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
