package capture

import (
	"github.com/gordonklaus/portaudio"
	"time"
)

type StreamPropertiesInput struct {
	Channels         int
	LatencyOfReading time.Duration
	SampleRate       float64
	FramesPerBuffer  int
	Flags            portaudio.StreamFlags
}

func DefaultStreamPropertiesInput() StreamPropertiesInput {
	return StreamPropertiesInput{
		Channels:         1, // mono by default
		LatencyOfReading: 100 * time.Millisecond,
		SampleRate:       48000.0,
		FramesPerBuffer:  256,
		Flags:            portaudio.NoFlag,
	}
}
