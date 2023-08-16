package realnotification

type Publisher interface {
	Start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

type ChPublisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan struct{}
}

func NewPublisher() *ChPublisher {
	return &ChPublisher{
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}
}

func (cp *ChPublisher) Start() {
	for {
		select {
		case msg := <-cp.in:
			for _, sub := range cp.subscribers {
				sub.Notify(msg)
			}
		case sub := <-cp.addSubCh:
			cp.subscribers = append(cp.subscribers, sub)
		case sub := <-cp.removeSubCh:
			for i, candidate := range cp.subscribers {
				if candidate == sub {
					cp.subscribers = append(cp.subscribers[:i],
						cp.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-cp.stop:
			for _, sub := range cp.subscribers {
				sub.Close()
			}
			close(cp.addSubCh)
			close(cp.in)
			close(cp.removeSubCh)
			return
		}
	}
}

func (cp *ChPublisher) AddSubscriberCh() chan<- Subscriber {
	return cp.addSubCh
}

func (cp *ChPublisher) RemoveSubscriberCh() chan<- Subscriber {
	return cp.removeSubCh
}

func (cp *ChPublisher) PublishingCh() chan<- interface{} {
	return cp.in
}

func (cp *ChPublisher) Stop() {
	close(cp.stop)
}
