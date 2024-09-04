package delasync

import (
	"context"
	"time"
)

type DelPair struct {
	UserID string
	Item   string
}

func InitWorker(ch <-chan DelPair, delFunc func(context.Context, []DelPair) error) {
	go func(in <-chan DelPair) {
		var buf []DelPair // сюда будем складывать накопленные

		rf := func(in <-chan DelPair) (DelPair, bool) {
			select {
			case val := <-in:
				return val, true
			case <-time.After(5 * time.Second):
				return DelPair{}, false
			}
		}

		for {
			val, found := rf(in)
			if found {
				buf = append(buf, val)
				if len(buf) > 2 {
					delFunc(context.Background(), buf)
					buf = nil
				}
			} else {
				if len(buf) > 0 {
					delFunc(context.Background(), buf)
					buf = nil
				}
			}
		}
	}(ch)
}
