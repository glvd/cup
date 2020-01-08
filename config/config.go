package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

// Config ...
type Config struct {
	Host string
}

var _config Config

// DefaultConfigName ...
var DefaultConfigName = "config.json"

// DefaultConfigPath ...
var DefaultConfigPath = "."

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

// LoadJSON ...
func LoadJSON() (err error) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	err = viper.MergeInConfig()
	if err != nil {
		return err
	}
	_config.Host = viper.GetString("host")
	return nil
}
