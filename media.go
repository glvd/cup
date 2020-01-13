package cup

import (
	"encoding/json"
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

// JSON ...
func (f *Fragment) JSON() (string, error) {
	marshal, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
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
