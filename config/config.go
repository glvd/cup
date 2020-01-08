package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// configValue is a wrapper aroung *pflag.flag
// that implements FlagValue
type configValue struct {
	flag *pflag.Flag
}

// HasChanges returns whether the flag has changes or not.
func (v configValue) HasChanged() bool {
	return v.flag.Changed
}

// Name returns the name of the flag.
func (v configValue) Name() string {
	return v.flag.Name
}

// ValueString returns the value of the flag as a string.
func (v configValue) ValueString() string {
	return v.flag.Value.String()
}

// ValueType returns the type of the flag as a string.
func (v configValue) ValueType() string {
	return v.flag.Value.Type()
}

type Config struct {
	flags *pflag.FlagSet
}

func (c *Config) VisitAll(fn func(viper.FlagValue)) {
	c.flags.VisitAll(func(flag *pflag.Flag) {
		fn(&configValue{flag: flag})
	})
}
