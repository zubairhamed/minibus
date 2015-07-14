package minibus

import "sync"

type TopicHandler func(interface{})

func NewMiniBus() *MiniBus {
	return &MiniBus{
		subscriptions: make(map[string][]TopicHandler),
	}
}

type MiniBus struct {
	subscriptions map[string][]TopicHandler
	sync.Mutex
}

func (b *MiniBus) Sub(topic string, fn TopicHandler) {
	b.Lock()
	defer b.Unlock()
	b.subscriptions[topic] = append(b.subscriptions[topic], fn)
}

func (b *MiniBus) Pub(topic string, msg interface{}) {
	b.Lock()
	defer b.Unlock()
	topics := b.subscriptions[topic]
	if len(topics) > 0 {
		for _, t := range topics {
			go t(msg)
		}
	}
}

func (b *MiniBus) ClearTopic(topic string) {
	b.Lock()
	defer b.Unlock()
	b.subscriptions[topic] = []TopicHandler{}

}

func (b *MiniBus) ClearSubs() {
	b.Lock()
	defer b.Unlock()
	b.subscriptions = make(map[string][]TopicHandler)
}

func (b *MiniBus) GetSub(topic string) []TopicHandler {
	return b.subscriptions[topic]
}
