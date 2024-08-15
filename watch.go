package main

import "time"

type watcher struct {
	mq        MQ
	subscribe Distributer
	period    time.Duration
}

func NewWatcher(mq MQ, sub Distributer) watcher {
	return watcher{
		mq:        mq,
		subscribe: sub,
		period:    20 * time.Millisecond,
	}
}

func (w *watcher) Watch() {

	timer := time.NewTimer(w.period)

	for {
		<-timer.C

		if w.mq.HasMsg() {
			msg := w.mq.Pop()
			w.subscribe.Distribute(msg)
		}
	}
}
