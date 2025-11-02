package pkg

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/player"
)

type DecoderInterface interface {
	AudioOutputGenerator(length int) ([]float32, error)
}

type EncoderInterface interface {
	AudioInputGenerator(samples []float32) error
}

func NewPlayer(ctx context.Context, decoder DecoderInterface) (*player.Player, error) {
	return player.NewDefaultPlayer(ctx, decoder)
}

func NewRecorder(ctx context.Context, encoder EncoderInterface) (*player.Recorder, error) {
	return player.NewDefaultRecorder(ctx, encoder)
}

func NewDuplex(ctx context.Context, decoder DecoderInterface, encoder EncoderInterface) (*player.Duplex, error) {
	return player.NewDuplexSession(ctx, decoder, encoder)
}
