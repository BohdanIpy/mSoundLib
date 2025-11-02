package player

import (
	"context"
	"github.com/BohdanIpy/mSoundLib/internal/playbacks"
	"github.com/BohdanIpy/mSoundLib/internal/stream"
	"github.com/BohdanIpy/mSoundLib/internal/stream_parameters/capture"
	"sync"
)

type RecorderState int

const (
	RecorderStopped RecorderState = iota
	RecorderRecording
	RecorderPaused
)

type EncoderInterface interface {
	AudioInputGenerator(samples []float32) error
}

type Recorder struct {
	ctx      context.Context
	cancel   context.CancelFunc
	stream   *stream.AudioStream
	capture  *playbacks.StreamCapture
	state    RecorderState
	events   *StreamEventHooks
	stateMux sync.Mutex
}

func (r *Recorder) Record() error {
	r.stateMux.Lock()
	defer r.stateMux.Unlock()

	if r.state == RecorderRecording {
		return nil
	}

	if r.events != nil && r.events.OnStart != nil {
		r.events.OnStart()
	}

	if err := r.stream.Stream.Start(); err != nil {
		if r.events != nil && r.events.OnError != nil {
			r.events.OnError(err)
		}
		return err
	}

	r.state = RecorderRecording
	return nil
}

func (r *Recorder) Pause() {
	r.stateMux.Lock()
	defer r.stateMux.Unlock()

	if r.state != RecorderRecording {
		return
	}

	r.state = RecorderPaused
	if r.events != nil && r.events.OnPause != nil {
		r.events.OnPause()
	}
}

func (r *Recorder) Stop() {
	r.stateMux.Lock()
	defer r.stateMux.Unlock()

	if r.state == RecorderStopped {
		return
	}

	_ = r.stream.Stream.Stop()
	if r.events != nil && r.events.OnStop != nil {
		r.events.OnStop()
	}

	r.state = RecorderStopped
	r.cancel()
}

func (r *Recorder) Close() {
	_ = r.stream.Stream.Close()
}

func NewDefaultRecorder(parentCtx context.Context, encoder EncoderInterface) (*Recorder, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	streamCapture := playbacks.NewStreamCapture(ctx, encoder)
	audioStream, err := stream.NewInputStream(capture.DefaultStreamPropertiesInput(), streamCapture.AsPortAudioCallback())
	if err != nil {
		cancel()
		return nil, err
	}

	return &Recorder{
		ctx:     ctx,
		cancel:  cancel,
		capture: streamCapture,
		stream:  audioStream,
		state:   RecorderStopped,
		events:  &StreamEventHooks{},
	}, nil
}
