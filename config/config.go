package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
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
	open, err := os.Open(filepath.Join(DefaultConfigPath, DefaultConfigName))
	if err != nil {
		return err
	}
	err = viper.MergeConfig(open)
	if err != nil {
		return err
	}
	log.Println("host", viper.Get("host"))
	_config.Host = viper.GetString("host")
	return nil
}
