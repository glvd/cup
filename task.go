package cup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/glvd/go-fftool"
)

// InitFFTool ...
func preinit(s config.SliceConfig) {
	if s.CommandPath != "" {
		fftool.DefaultCommandPath = s.CommandPath
	}
}

// TaskSlice ...
func TaskSlice(ctx context.Context, s string) (f string, e error) {
	cfg := config.SliceConfig{}
	e = json.Unmarshal([]byte(s), &cfg)
	if e != nil {
		return "", e
	}
	sliced, e := Slice(ctx, cfg)
	if e != nil {
		return "", e
	}
	indent, e := json.MarshalIndent(sliced, "", " ")
	if e != nil {
		return "", e
	}
	return string(indent), nil
}

// Slice ...
func Slice(ctx context.Context, s config.SliceConfig) (f *Fragment, e error) {
	preinit(s)
	ffprobe := fftool.NewFFProbe()
	if s.FFProbeName != "" {
		ffprobe.Name = s.FFProbeName
	}
	format, e := ffprobe.StreamFormat(s.Filepath)
	if e != nil {
		return nil, fmt.Errorf("ffprobe error:%w", e)
	}
	if !IsMedia(format) {
		return nil, errors.New("file is not a video/audio")
	}
	cfg := fftool.DefaultConfig()
	cfg.SetSlice(true)
	cfg.OutputPath = s.OutputPath
	cfg.Scale = s.Scale
	if s.Crypto != nil {
		cfg.SetCrypt(*s.Crypto)
	}

	sharpness := fmt.Sprintf("%dP", fftool.ScaleValue(cfg.Scale))
	ff := fftool.NewFFMpeg(cfg)
	if s.FFMpegName != "" {
		ff.Name = s.FFMpegName
	}
	if err := fftool.OptimizeWithFormat(cfg, format); err != nil {
		return nil, err
	}
	e = ff.Run(ctx, s.Filepath)
	if e != nil {
		return nil, fmt.Errorf("run error:%w", e)
	}

	return &Fragment{
		Scale:     cfg.Scale,
		Output:    cfg.ProcessPath(),
		Input:     s.Filepath,
		Sharpness: sharpness,
	}, nil
}
