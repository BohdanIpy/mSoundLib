package playback

import (
	"github.com/gordonklaus/portaudio"
	"time"
)

// StreamPropertiesOutput used for configuring the stream parameters
// you can specify the:
// channels, latency of writing, sample rate, frames per buffer and flags /*
type StreamPropertiesOutput struct {
	Channels         int
	LatencyOfWriting time.Duration
	SampleRate       float64
	FramesPerBuffer  int
	Flags            portaudio.StreamFlags
}

func DefaultStreamPropertiesOutput() StreamPropertiesOutput {
	return StreamPropertiesOutput{
		Channels:         2,
		LatencyOfWriting: 100 * time.Millisecond,
		SampleRate:       48000.0,
		FramesPerBuffer:  256,
		Flags:            portaudio.NoFlag,
	}
}
