package player

import (
	"context"
	"sync"
)

type DuplexState int

const (
	DuplexStopped DuplexState = iota
	DuplexActive
	DuplexPaused
)

type Duplex struct {
	ctx        context.Context
	cancel     context.CancelFunc
	player     *Player   // handles playback
	recorder   *Recorder // handles input capture
	state      DuplexState
	events     *StreamEventHooks
	stateMutex sync.Mutex
}

func (d *Duplex) Start() error {
	d.stateMutex.Lock()
	defer d.stateMutex.Unlock()

	if d.state == DuplexActive {
		return nil
	}

	if d.events != nil && d.events.OnStart != nil {
		d.events.OnStart()
	}

	d.player.Play()

	if err := d.recorder.Record(); err != nil {
		d.player.Stop()
		return err
	}

	d.state = DuplexActive
	return nil
}

func (d *Duplex) Pause() {
	d.stateMutex.Lock()
	defer d.stateMutex.Unlock()

	if d.state != DuplexActive {
		return
	}

	d.player.Pause()
	d.recorder.Pause()

	if d.events != nil && d.events.OnPause != nil {
		d.events.OnPause()
	}

	d.state = DuplexPaused
}

func (d *Duplex) Stop() {
	d.stateMutex.Lock()
	defer d.stateMutex.Unlock()

	if d.state == DuplexStopped {
		return
	}

	d.player.Stop()
	d.recorder.Stop()

	if d.events != nil && d.events.OnStop != nil {
		d.events.OnStop()
	}

	d.cancel() // stops all callbacks
	d.state = DuplexStopped
}

func (d *Duplex) Close() {
	d.player.Close()
	d.recorder.Close()
}

func NewDuplexSession(parentCtx context.Context, decoder DecoderInterface, encoder EncoderInterface) (*Duplex, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	playerObj, err := NewDefaultPlayer(ctx, decoder)
	if err != nil {
		cancel()
		return nil, err
	}

	recorderObj, err := NewDefaultRecorder(ctx, encoder)
	if err != nil {
		playerObj.Close()
		cancel()
		return nil, err
	}

	return &Duplex{
		ctx:      ctx,
		cancel:   cancel,
		player:   playerObj,
		recorder: recorderObj,
		state:    DuplexStopped,
		events:   &StreamEventHooks{},
	}, nil
}
