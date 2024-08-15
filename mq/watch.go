package mq

import (
	"time"

	"github.com/victorzhou123/simplemq/consume"
)

type watcher struct {
	mq        MQ
	subscribe consume.Distributer
	period    time.Duration
}

func NewWatcher(mq MQ, sub consume.Distributer) watcher {
	return watcher{
		mq:        mq,
		subscribe: sub,
		period:    20 * time.Millisecond,
	}
}

func (w *watcher) Watch() {

	ticker := time.NewTicker(w.period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if w.mq.HasMsg() {
				msg := w.mq.Pop()
				w.subscribe.Distribute(msg)
			}
		}
	}
}
