package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/secure"
	"mishin-shortener/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserURLs(t *testing.T) {
	t.Run("Start_POST_to_create_record_in_db", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		stor.EXPECT().GetByUserID(
			ctx,
			"userId",
		).Return(map[string]string{"short0": "long0", "short1": "long1"}, nil)

		c := config.MakeConfig()
		h := MakeShortanerHandler(c, stor)

		request :=
			httptest.NewRequest(
				http.MethodGet,
				"/api/user/urls",
				nil,
			).WithContext(ctx)

		w := httptest.NewRecorder()
		h.UserURLs(w, request)
		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		list := make([]UserURLsItem, 0)

		resBody, _ := io.ReadAll(res.Body)
		fmt.Println(string(resBody))

		json.Unmarshal(resBody, &list)

		assert.Contains(t, list, UserURLsItem{
			Short:    "http://localhost:8080/short0",
			Original: "long0",
		})
		assert.Contains(t, list, UserURLsItem{
			Short:    "http://localhost:8080/short1",
			Original: "long1",
		})
	})
}
