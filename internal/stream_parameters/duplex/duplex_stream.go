package duplex

import (
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/capture"
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/playback"
	"github.com/gordonklaus/portaudio"
)

func GetDuplexStreamParameters(props DuplexStreamProperties) (portaudio.StreamParameters, error) {
	inputParams := capture.InitializeStreamDeviceParameterInput(props.InputDevice, props.InputChannels, props.InputLatency)
	outputParams := playback.InitializeStreamDeviceParameterOutput(props.OutputDevice, props.OutputChannels, props.OutputLatency)

	return portaudio.StreamParameters{
		Input:           inputParams,
		Output:          outputParams,
		SampleRate:      props.SampleRate,
		FramesPerBuffer: props.FramesPerBuffer,
		Flags:           props.Flags,
	}, nil
}
