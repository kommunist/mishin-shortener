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
	t.Run("Start_POST_to_create_record_in_db", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		stor.EXPECT().DeleteByUserID(
			ctx,
			"userId",
			[]string{"first", "second"},
		)

		c := config.MakeConfig()
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

		assert.Equal(t, http.StatusAccepted, res.StatusCode)
	})

}