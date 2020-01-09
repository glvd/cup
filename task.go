package cup

import (
	"context"
	"errors"
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/glvd/go-fftool"
)

var _ffprobe *fftool.FFProbe

// InitFFTool ...
func preinit(s config.SliceConfig) {
	if s.CommandPath != "" {
		fftool.DefaultCommandPath = s.CommandPath
	}
	_ffprobe = fftool.NewFFProbe()
	if s.FFProbeName != "" {
		_ffprobe.Name = s.FFProbeName
	}
}

// Slice ...
func Slice(ctx context.Context, s config.SliceConfig) (f *Fragment, e error) {
	preinit(s)
	format, e := _ffprobe.StreamFormat(s.Filepath)
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

	ff = ff.OptimizeWithFormat(format)

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
