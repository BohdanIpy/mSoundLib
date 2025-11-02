package stream

import (
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/capture"
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/duplex"
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/playback"
	"github.com/gordonklaus/portaudio"
)

type StreamCallback func(in []float32, out []float32, timeInfo portaudio.StreamCallbackTimeInfo, flags portaudio.StreamCallbackFlags)

func NewOutputStream(properties playback.StreamPropertiesOutput, callback StreamCallback) (*AudioStream, error) {
	outputParameters, err := playback.GetOutputStreamParameters(properties)
	if err != nil {
		return nil, err
	}
	stream, err := portaudio.OpenStream(outputParameters, callback)
	if err != nil {
		return nil, err
	}
	return &AudioStream{Stream: stream, StreamMode: OUTPUT_AUDIO_STREAM}, nil
}

func NewInputStream(props capture.StreamPropertiesInput, callback StreamCallback) (*AudioStream, error) {
	params, err := capture.GetInputStreamParameters(props)
	if err != nil {
		return nil, err
	}

	stream, err := portaudio.OpenStream(params, callback)
	if err != nil {
		return nil, err
	}

	return &AudioStream{Stream: stream, StreamMode: INPUT_AUDIO_STREAM}, nil
}

func NewDuplexStream(props duplex.DuplexStreamProperties, callback StreamCallback) (*AudioStream, error) {
	params, err := duplex.GetDuplexStreamParameters(props)
	if err != nil {
		return nil, err
	}

	stream, err := portaudio.OpenStream(params, callback)
	if err != nil {
		return nil, err
	}

	return &AudioStream{Stream: stream, StreamMode: DUPLEX_AUDIO_STREAM}, nil
}
