package player

type StreamEventHooks struct {
	OnStart  func()
	OnStop   func()
	OnPause  func()
	OnResume func()
	OnEnd    func()

	OnError func(err error)

	OnBufferUnderrun func()
	OnBufferOverflow func()
	OnLatencyChange  func(latency float64)

	OnSeek func(position float64)
}
