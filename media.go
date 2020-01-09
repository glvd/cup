package cup

import "github.com/glvd/go-fftool"

// IsMedia ...
func IsMedia(format *fftool.StreamFormat) bool {
	video := format.Video()
	audio := format.Audio()
	if audio == nil || video == nil {
		return false
	}
	return true
}
