package playback

import (
	"github.com/gordonklaus/portaudio"
	"time"
)

func init() {
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}
}

func GetDefaultOutputDevice() (*portaudio.DeviceInfo, error) {
	// get the hardware parts, eg your headset, earbuds, etc
	return portaudio.DefaultOutputDevice()
}

func InitializeStreamDeviceParameterOutput(outputDevice *portaudio.DeviceInfo, channels int, latency time.Duration) portaudio.StreamDeviceParameters {
	return portaudio.StreamDeviceParameters{
		Device: outputDevice,
		// 2 - stereo, 1 - mono, <2 - custom
		// how much separate outputs
		Channels: channels,
		// delay of writing
		Latency: latency,
	}
}

func GetOutputStreamParameters(properties StreamPropertiesOutput) (portaudio.StreamParameters, error) {
	outputDevice, err := GetDefaultOutputDevice()
	if err != nil {
		return portaudio.StreamParameters{}, err
	}

	streamDeviceParameter := InitializeStreamDeviceParameterOutput(outputDevice, properties.Channels, properties.LatencyOfWriting)

	return portaudio.StreamParameters{
		// where to write
		Output: streamDeviceParameter,
		// kHz of the sound
		SampleRate: properties.SampleRate,
		// how many portion send to device at once
		FramesPerBuffer: properties.FramesPerBuffer,
		// portaudio.StreamFlags
		Flags: properties.Flags,
	}, nil
}
