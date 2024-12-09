package redirect

import (
	"context"
	"errors"
	"mishin-shortener/internal/errors/deleted"
	pb "mishin-shortener/proto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGet(t *testing.T) {
	exList := []struct {
		name          string
		shorted       string
		expected      string
		storError     error
		returnedError error
	}{
		{
			name:          "Start_GET_on_url_persisted_in_db",
			shorted:       "qwerty",
			expected:      "ya.ru",
			returnedError: nil,
		},
		{
			name:          "Start_GET_on_url_already_deleted",
			shorted:       "qqqq",
			expected:      "",
			storError:     deleted.NewDeletedError(nil),
			returnedError: status.Error(codes.OutOfRange, "Error when call service"),
		},
		{
			name:          "Start_GET_on_another_error",
			shorted:       "qqqq",
			expected:      "",
			returnedError: status.Error(codes.NotFound, "Error when call service"),
			storError:     errors.New("ququ"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockGetter(ctrl)
			h := Make(stor)

			ctx := context.Background()

			stor.EXPECT().Get(ctx, ex.shorted).Return(ex.expected, ex.storError)

			resp, err := h.Get(ctx, &pb.GetRequest{Short: ex.shorted})

			if ex.returnedError == nil {
				assert.NoError(t, err)
				assert.Equal(t, resp.Original, ex.expected)
			} else {
				assert.EqualError(t, err, ex.returnedError.Error())
			}
		})

	}
}
