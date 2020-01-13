package config

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/glvd/go-fftool"
	"github.com/goextension/extmap"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

// Config ...
type Config struct {
	Name          string
	CommandPath   string
	FFProbeName   string
	FFMpegName    string
	AutoRemove    bool
	ProcessCore   ProcessCore
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

var typeExt = map[string]string{
	"json": ".json",
	"toml": ".toml",
	"yaml": ".yml",
}

// Default ...
func Default() *Config {
	return &Config{
		Name:          "cup",
		CommandPath:   "",
		FFProbeName:   "",
		FFMpegName:    "",
		AutoRemove:    true,
		ProcessCore:   fftool.ProcessCUDA,
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

// Save ...
func Save() (err error) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	viper.SetConfigType(DefaultConfigType)
	err = viper.MergeConfigMap(extmap.StructToMap(_config))
	if err != nil {
		return err
	}
	viper.SetConfigFile(filepath.Join(DefaultConfigPath, DefaultConfigName+Ext()))

	return viper.WriteConfig()

}

// Load ...
func Load() (err error) {
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)

	err = viper.MergeInConfig()
	if err != nil {
		return err
	}
	m := extmap.ToMap(viper.AllSettings())
	return m.Struct(&_config)
}

// Ext ...
func Ext() string {
	if v, b := typeExt[DefaultConfigType]; b {
		return v
	}
	return ".json"
}
