package deleteurls

import (
	"bytes"
	"context"
	"encoding/json"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/delasync"
	"mishin-shortener/internal/app/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	t.Run("Start_DELETE_to_delete_record_in_db", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		c := config.MakeConfig()
		c.InitConfig()
		h := Make(make(chan delasync.DelPair, 5))

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete, "/api/user/urls", bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		h.Call(w, request)
		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, http.StatusAccepted, res.StatusCode)
	})

	t.Run("Start_DELETE_to_delete_record_in_db_without_user", func(t *testing.T) {
		ctx := context.Background()
		c := config.MakeConfig()
		c.InitConfig()
		h := Make(make(chan delasync.DelPair, 5))

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete, "/api/user/urls", bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		h.Call(w, request)
		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

}

func BenchmarkDeleteUrls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		c := config.MakeConfig()
		c.InitConfig()
		h := Make(make(chan delasync.DelPair, 5))

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete,
				"/api/user/urls",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		b.StartTimer()
		h.Call(w, request)
	}
}
