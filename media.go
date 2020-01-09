package cup

import (
	"github.com/glvd/cup/config"
	"github.com/glvd/go-fftool"
)

// Fragment ...
type Fragment struct {
	Scale     config.Scale
	Output    string
	Input     string
	Sharpness string
}

// IsMedia ...
func IsMedia(format *fftool.StreamFormat) bool {
	video := format.Video()
	audio := format.Audio()
	if audio == nil || video == nil {
		return false
	}
	return true
}
