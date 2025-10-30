package capture

import (
	"github.com/gordonklaus/portaudio"
	"time"
)

func init() {
	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}
}

func GetDefaultInputDevice() (*portaudio.DeviceInfo, error) {
	// get the hardware parts, eg your headset, earbuds, etc
	return portaudio.DefaultInputDevice()
}

func InitializeStreamDeviceParameterInput(device *portaudio.DeviceInfo, channels int, latency time.Duration) portaudio.StreamDeviceParameters {
	return portaudio.StreamDeviceParameters{
		Device:   device,
		Channels: channels,
		Latency:  latency,
	}
}

func GetInputStreamParameters(props StreamPropertiesInput) (portaudio.StreamParameters, error) {
	inputDevice, err := GetDefaultInputDevice()
	if err != nil {
		return portaudio.StreamParameters{}, err
	}

	streamDeviceParameter := InitializeStreamDeviceParameterInput(
		inputDevice,
		props.Channels,
		props.LatencyOfReading,
	)

	return portaudio.StreamParameters{
		Input:           streamDeviceParameter,
		SampleRate:      props.SampleRate,
		FramesPerBuffer: props.FramesPerBuffer,
		Flags:           props.Flags,
	}, nil
}
