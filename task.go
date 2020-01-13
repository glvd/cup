package cup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/glvd/go-fftool"
	"github.com/glvd/go-fftool/factory"
	"time"
)

func runInit() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	cfg := config.Get()

	factory.Initialize(func(option *factory.Option) {
		if cfg.FFMpegName != "" {
			option.MpegName = cfg.FFMpegName
		}
		if cfg.FFProbeName != "" {
			option.ProbeName = cfg.FFProbeName
		}

		if cfg.CommandPath != "" {
			option.CommandPath = cfg.CommandPath
		}
	})

}

// DummySlice ...
func DummySlice(ctx context.Context, s []byte) (f string, e error) {
	fmt.Println("received:", string(s))
	b, e := json.MarshalIndent(&Fragment{
		Scale:     fftool.Scale720P,
		Output:    "dummy",
		Input:     "dummy",
		Sharpness: "720P",
	}, "", " ")
	if e != nil {
		return "", e
	}
	time.Sleep(5 * time.Second)
	return string(b), nil
}

// TaskSlice ...
func TaskSlice(ctx context.Context, s string) (f string, e error) {
	cfg := config.DefaultSliceConfig()
	e = cfg.Parse(s)
	if e != nil {
		return "", e
	}
	frag, e := Slice(ctx, cfg)
	if e != nil {
		return "", e
	}
	return frag.JSON()
}

// Slice ...
func Slice(ctx context.Context, s *config.SliceConfig) (f *Fragment, e error) {
	probe := factory.Probe()
	format, e := probe.StreamFormat(s.Filepath)
	if e != nil {
		return nil, fmt.Errorf("probe error:%w", e)
	}
	if !IsMedia(format) {
		return nil, errors.New("file is not a video/audio")
	}
	cfg := fftool.DefaultConfig()
	cfg.Slice = true
	cfg.OutputPath = s.OutputPath
	cfg.Scale = s.Scale
	if s.Crypto != nil {
		cfg.SetCrypt(*s.Crypto)
	}

	sharpness := fmt.Sprintf("%dP", fftool.ScaleValue(cfg.Scale))
	mpeg := factory.Mpeg()
	if err := fftool.OptimizeWithFormat(cfg, format); err != nil {
		return nil, err
	}
	e = mpeg.Run(ctx, s.Filepath, func(cfg *fftool.Config) *fftool.Config {
		return cfg
	})
	if e != nil {
		return nil, fmt.Errorf("run error:%w", e)
	}

	f = &Fragment{
		Scale:     cfg.Scale,
		Output:    cfg.ProcessPath(),
		Input:     s.Filepath,
		Sharpness: sharpness,
	}
	return f, nil
}
