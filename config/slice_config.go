package config

import "github.com/glvd/go-fftool"

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
	CommandPath     string
	FFProbeName     string
	AutoRemove      bool
	Scale           Scale
	ProcessCore     ProcessCore
	ProcessID       string
	BitRate         int64
	FrameRate       float64
	OutputPath      string //output path
	OutputName      string
	M3U8Name        string
	SegmentFileName string
	HLSTime         int
	KeyOutput       bool
}

// DefaultSliceConfig ...
func DefaultSliceConfig() *SliceConfig {
	return &SliceConfig{
		Filepath:        "",
		Crypto:          nil,
		AutoRemove:      true,
		Scale:           fftool.Scale720P,
		ProcessCore:     fftool.DefaultProcessCore,
		ProcessID:       "",
		BitRate:         0,
		FrameRate:       0,
		OutputPath:      fftool.DefaultOutputPath,
		OutputName:      fftool.DefaultOutputName,
		M3U8Name:        fftool.DefaultM3U8Name,
		SegmentFileName: fftool.DefaultSegmentFileName,
		HLSTime:         fftool.DefaultHLSTime,
		KeyOutput:       true,
	}
}
