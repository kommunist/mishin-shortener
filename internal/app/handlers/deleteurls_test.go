package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/secure"
	"mishin-shortener/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteURLs(t *testing.T) {
	t.Run("Start_DELETE_to_delete_record_in_db", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		// stor.EXPECT().DeleteByUserID(
		// 	ctx,
		// 	"userId",
		// 	[]string{"first", "second"},
		// )

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete,
				"/api/user/urls",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		h.DeleteURLs(w, request)
		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, http.StatusAccepted, res.StatusCode)
	})

	t.Run("Start_DELETE_to_delete_record_in_db_without_user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()

		// stor.EXPECT().DeleteByUserID(
		// 	ctx,
		// 	"userId",
		// 	[]string{"first", "second"},
		// ).Times(0)

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete,
				"/api/user/urls",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		h.DeleteURLs(w, request)
		res := w.Result()

		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

}

func BenchmarkDeleteUrls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ctrl := gomock.NewController(b)
		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		inputJSON, _ := json.Marshal([]string{"first", "second"})

		request :=
			httptest.NewRequest(
				http.MethodDelete,
				"/api/user/urls",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		w := httptest.NewRecorder()
		b.StartTimer()
		h.DeleteURLs(w, request)
	}
}
