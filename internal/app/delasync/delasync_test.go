package delasync

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	t.Run("start_async_function_for_remove_with_auto_flush", func(t *testing.T) {
		ch := make(chan DelPair, 5)
		result := make([]DelPair, 0)

		testDelFunc := func(ctx context.Context, list []DelPair) error {
			result = append(result, list...)
			return nil
		}
		assert.Equal(t, 0, len(result))

		InitWorker(ch, testDelFunc)

		ch <- DelPair{UserID: "123", Item: "qq"}
		ch <- DelPair{UserID: "456", Item: "pp"}
		ch <- DelPair{UserID: "789", Item: "mm"}

		time.Sleep(2 * time.Second) // пододем, когда обработает

		assert.Equal(t, 3, len(result))
	})

	t.Run("start_async_function_for_remove_with_flush_by_time", func(t *testing.T) {
		ch := make(chan DelPair, 5)
		result := make([]DelPair, 0)

		testDelFunc := func(ctx context.Context, list []DelPair) error {
			result = append(result, list...)
			return nil
		}
		assert.Equal(t, 0, len(result))

		InitWorker(ch, testDelFunc)

		ch <- DelPair{UserID: "123", Item: "qq"}

		time.Sleep(6 * time.Second) // пододем, когда обработает

		assert.Equal(t, 1, len(result))
	})
}
