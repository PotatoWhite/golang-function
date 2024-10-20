package observer

type Observer func(string)

type EventEmitter struct {
	observers []Observer
}

func (s *EventEmitter) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *EventEmitter) NotifyAll(msg string) {
	for _, o := range s.observers {
		o(msg)
	}
}
