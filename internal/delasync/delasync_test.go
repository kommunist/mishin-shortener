package delasync

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
)

func TestWorker(t *testing.T) {
	exList := []struct {
		name        string
		pairs       []DelPair
		waitSeconds int
	}{
		{
			name: "start_async_function_for_remove_with_auto_flush",
			pairs: []DelPair{
				{UserID: "123", Item: "qq"},
				{UserID: "456", Item: "pp"},
				{UserID: "789", Item: "mm"},
			},
			waitSeconds: 2,
		},
		{
			name: "start_async_function_for_remove_with_flush_by_time",
			pairs: []DelPair{
				{UserID: "123", Item: "qq"},
			},
			waitSeconds: 6,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockRemover(ctrl)
			h := Make(stor)
			defer h.Stop()

			stor.EXPECT().DeleteByUserID(
				gomock.Any(),
				gomock.Len(len(ex.pairs)),
			)

			h.InitWorker()

			for _, pair := range ex.pairs {
				h.DelChan <- pair
			}

			time.Sleep(time.Duration(ex.waitSeconds) * time.Second)

		})

	}
}
