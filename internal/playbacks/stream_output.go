package playbacks

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/stream"
	"github.com/gordonklaus/portaudio"
)

type StreamPlayback struct {
	ctx     context.Context
	decoder DecoderInterface
}

func (s *StreamPlayback) AsPortAudioCallback() stream.StreamCallback {
	return func(in, out []float32, timeInfo portaudio.StreamCallbackTimeInfo, flags portaudio.StreamCallbackFlags) {
		select {
		case <-s.ctx.Done():
			for i := range out {
				out[i] = 0
			}
			return
		default:
		}
		dataArr, err := s.decoder.AudioOutputGenerator(len(out))
		if err != nil {
			for i := range out {
				out[i] = 0
			}
			return
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

func NewStreamPlayback(ctx context.Context, decoder DecoderInterface) *StreamPlayback {
	return &StreamPlayback{
		ctx:     ctx,
		decoder: decoder,
	}
}
