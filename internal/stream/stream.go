package stream

import "github.com/gordonklaus/portaudio"

type StreamMode int

const (
	INPUT_AUDIO_STREAM = iota
	OUTPUT_AUDIO_STREAM
	DUPLEX_AUDIO_STREAM
)

type AudioStream struct {
	Stream     *portaudio.Stream
	StreamMode StreamMode
}

/*
func (s *AudioStream) Start() error {
	return s.stream.Start()
}

func (s *AudioStream) Stop() error {
	return s.stream.Stop()
}

func (s *AudioStream) Close() error {
	return s.stream.Close()
}
*/
