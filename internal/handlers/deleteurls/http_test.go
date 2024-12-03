package deleteurls

import (
	"bytes"
	"context"
	"encoding/json"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name   string
		ctx    context.Context
		input  []byte
		status int
	}{
		{
			name: "Start_DELETE_to_delete_record_in_db",
			ctx:  context.WithValue(context.Background(), secure.UserIDKey, "userId"),
			input: func() []byte {
				json, _ := json.Marshal([]string{"first", "second"})
				return json
			}(),
			status: http.StatusAccepted,
		},
		{
			name: "Start_DELETE_to_delete_record_in_db_without_user",
			ctx:  context.Background(),
			input: func() []byte {
				json, _ := json.Marshal([]string{"first", "second"})
				return json
			}(),
			status: http.StatusInternalServerError,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {

			c := config.MakeConfig()
			c.InitConfig()
			h := Make(make(chan delasync.DelPair, 5))

			request :=
				httptest.NewRequest(
					http.MethodDelete, "/api/user/urls", bytes.NewReader(ex.input),
				).WithContext(ex.ctx)

			w := httptest.NewRecorder()
			h.Call(w, request)
			res := w.Result()

			defer res.Body.Close()

			assert.Equal(t, ex.status, res.StatusCode)
		})
	}
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
