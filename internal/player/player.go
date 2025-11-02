package player

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/playbacks"
	"github.com/BohdanIpy/mSoundLib/internal/stream"
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/playback"
	"sync"
)

type PlayerState int

const (
	PlayerStopped PlayerState = iota
	PlayerPlaying
	PlayerPaused
)

type DecoderInterface interface {
	AudioOutputGenerator(length int) ([]float32, error)
}

type Player struct {
	ctx        context.Context
	cancel     context.CancelFunc
	stream     *stream.AudioStream
	playback   *playbacks.StreamPlayback
	state      PlayerState
	events     *StreamEventHooks
	stateMutex sync.Mutex
}

func (p *Player) Play() {
	p.stateMutex.Lock()
	defer p.stateMutex.Unlock()

	if p.state == PlayerPlaying {
		return
	}

	if p.events != nil && p.events.OnStart != nil {
		p.events.OnStart()
	}

	_ = p.stream.Stream.Start()
	p.state = PlayerPlaying
}

func (p *Player) Pause() {
	p.stateMutex.Lock()
	defer p.stateMutex.Unlock()

	if p.state != PlayerPlaying {
		return
	}

	_ = p.stream.Stream.Stop()
	p.state = PlayerPaused

	if p.events != nil && p.events.OnPause != nil {
		p.events.OnPause()
	}
}

func (p *Player) Stop() {
	p.stateMutex.Lock()
	defer p.stateMutex.Unlock()

	if p.state == PlayerStopped {
		return
	}

	_ = p.stream.Stream.Stop()
	if p.events != nil && p.events.OnStop != nil {
		p.events.OnStop()
	}

	p.state = PlayerStopped
	p.cancel()
}

func (p *Player) Close() {
	_ = p.stream.Stream.Close()
}

func NewDefaultPlayer(parentCtx context.Context, decoder DecoderInterface) (*Player, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	playbackObj := playbacks.NewStreamPlayback(ctx, decoder)
	audioStream, err := stream.NewOutputStream(playback.DefaultStreamPropertiesOutput(), playbackObj.AsPortAudioCallback())
	if err != nil {
		cancel()
		return nil, err
	}

	return &Player{
		ctx:      ctx,
		cancel:   cancel,
		playback: playbackObj,
		stream:   audioStream,
		state:    PlayerStopped,
		events:   &StreamEventHooks{},
	}, nil
}
