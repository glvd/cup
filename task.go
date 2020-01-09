package cup

import (
	"context"
	"errors"
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/glvd/go-fftool"
)

// Crypto ...
type Crypto = fftool.Crypto

// Config ...
type Config = fftool.Config

var _ffprobe *fftool.FFProbe

// InitFFTool ...
func InitFFTool() {
	_ffprobe = fftool.NewFFProbe()
}

// Slice ...
func Slice(ctx context.Context, s config.SliceConfig) error {
	format, e := _ffprobe.StreamFormat(s.Filepath)
	//format, e := split.FFProbeStreamFormat(input)
	if e != nil {
		return fmt.Errorf("ffprobe error:%w", e)
	}
	if !IsMedia(format) {
		return errors.New("file is not a video/audio")
	}
	cfg := s.Config
	cfg.SetSlice(true)
	cfg.OutputPath = s.Output
	//cfg.Scale = w.WorkImpl.Scale
	//if w.Crypto != nil {
	//	cfg.SetCrypt(*w.Crypto)
	//}

	sharpness := fmt.Sprintf("%dP", fftool.ScaleValue(cfg.Scale))
	ff := fftool.NewFFMpeg(cfg)

	ff = ff.OptimizeWithFormat(format)

	e = ff.Run(ctx, input)
	//sa, e := split.FFMpegSplitToM3U8(ctx, input, split.StreamFormatOption(format), split.ScaleOption(formatScale(w.Scale)), split.OutputOption(w.Output()), split.AutoOption(true))
	if e != nil {
		return nil, Wrap(e)
	}

	return &Fragment{
		scale:     cfg.Scale,
		output:    cfg.ProcessPath(),
		skip:      w.Skip,
		input:     input,
		sharpness: sharpness,
	}, nil
}
