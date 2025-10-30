package duplex

import (
	"github.com/gordonklaus/portaudio"
	"time"
)

type DuplexStreamProperties struct {
	InputDevice   *portaudio.DeviceInfo
	InputChannels int
	InputLatency  time.Duration

	OutputDevice   *portaudio.DeviceInfo
	OutputChannels int
	OutputLatency  time.Duration

	// Shared stream properties
	SampleRate      float64
	FramesPerBuffer int
	Flags           portaudio.StreamFlags
}

func DefaultDuplexStreamProperties() (DuplexStreamProperties, error) {
	inputDevice, err := portaudio.DefaultInputDevice()
	if err != nil {
		return DuplexStreamProperties{}, err
	}

	outputDevice, err := portaudio.DefaultOutputDevice()
	if err != nil {
		return DuplexStreamProperties{}, err
	}

	return DuplexStreamProperties{
		InputDevice:   inputDevice,
		InputChannels: 1, // mono input (microphone)
		InputLatency:  100 * time.Millisecond,

		OutputDevice:   outputDevice,
		OutputChannels: 2, // stereo output (speakers)
		OutputLatency:  100 * time.Millisecond,

		SampleRate:      48000.0,
		FramesPerBuffer: 256,
		Flags:           portaudio.NoFlag,
	}, nil
}
