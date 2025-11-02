package playbacks

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/stream"
	"github.com/gordonklaus/portaudio"
)

type StreamDuplex struct {
	ctx     context.Context
	encoder EncoderInterface
	decoder DecoderInterface
}

func (s *StreamDuplex) AsPortAudioCallback() stream.StreamCallback {
	return func(in, out []float32, timeInfo portaudio.StreamCallbackTimeInfo, flags portaudio.StreamCallbackFlags) {
		select {
		case <-s.ctx.Done():
			clear(out)
		default:
		}

		if len(in) > 0 {
			_ = s.encoder.AudioInputGenerator(in)
		}

		dataArr, err := s.decoder.AudioOutputGenerator(len(out))
		if err != nil {
			clear(out)
		}

		for i := range out {
			if i < len(dataArr) {
				out[i] = dataArr[i]
			} else {
				out[i] = 0
			}
		}
	}
}
