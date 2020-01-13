package config

import (
	"encoding/json"
	"github.com/glvd/go-fftool"
)

// Scale ...
type Scale = fftool.Scale

// ProcessCore ...
type ProcessCore = fftool.ProcessCore

// Crypto ...
type Crypto = fftool.Crypto

// SliceConfig ...
type SliceConfig struct {
	Filepath        string
	Crypto          *Crypto
	Scale           Scale
	BitRate         int64
	FrameRate       float64
	OutputPath      string //output path
	OutputName      string
	M3U8Name        string
	SegmentFileName string
	HLSTime         int
	KeyOutput       bool
	LogOutput       bool
}

// DefaultSliceConfig ...
func DefaultSliceConfig() *SliceConfig {
	return &SliceConfig{
		Filepath:        "",
		Crypto:          nil,
		Scale:           fftool.Scale720P,
		BitRate:         0,
		FrameRate:       0,
		OutputPath:      fftool.DefaultOutputPath,
		OutputName:      fftool.DefaultOutputName,
		M3U8Name:        fftool.DefaultM3U8Name,
		SegmentFileName: fftool.DefaultSegmentFileName,
		HLSTime:         fftool.DefaultHLSTime,
		KeyOutput:       true,
		LogOutput:       true,
	}
}

// String ...
func (c SliceConfig) String() string {
	m, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(m)
}

// Parse ...
func (c *SliceConfig) Parse(s []byte) error {
	return json.Unmarshal(s, c)
}

// ParseSliceConfig ...
func ParseSliceConfig(s []byte) (*SliceConfig, error) {
	var cfg SliceConfig
	err := cfg.Parse(s)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
