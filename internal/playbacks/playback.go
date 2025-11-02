package playbacks

import "github.com/BohdanIpy/mSoundLib/internal/stream"

type PortaudioPlayback interface {
	AsPortAudioCallback() stream.StreamCallback
}

type DecoderInterface interface {
	AudioOutputGenerator(length int) ([]float32, error)
}

type EncoderInterface interface {
	AudioInputGenerator(samples []float32) error
}
