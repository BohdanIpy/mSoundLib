package playbacks

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/stream"
	"github.com/gordonklaus/portaudio"
)

type StreamCapture struct {
	ctx     context.Context
	encoder EncoderInterface
}

func (s *StreamCapture) AsPortAudioCallback() stream.StreamCallback {
	return func(in []float32, out []float32, timeInfo portaudio.StreamCallbackTimeInfo, flags portaudio.StreamCallbackFlags) {
		select {
		case <-s.ctx.Done():
			clear(out)
		default:
		}

		if len(in) > 0 {
			err := s.encoder.AudioInputGenerator(in)
			if err != nil {
				clear(out)
			}
		}
	}
}

func NewStreamCapture(ctx context.Context, encoder EncoderInterface) *StreamCapture {
	return &StreamCapture{
		ctx:     ctx,
		encoder: encoder,
	}
}
