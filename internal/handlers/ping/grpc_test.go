package ping

import (
	"context"
	"errors"
	pb "mishin-shortener/proto"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCallGRPC(t *testing.T) {
	exList := []struct {
		name         string
		storReturned error
		status       int
		respError    error
	}{
		{
			name:         "ping_storage_when_all_success",
			storReturned: nil,
			status:       http.StatusOK,
			respError:    nil,
		},
		{
			name:         "ping_storage_when_ping_return_error",
			storReturned: errors.New("ququ"),
			status:       http.StatusInternalServerError,
			respError:    status.Error(codes.Unknown, "Error when call service"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			stor := NewMockPinger(ctrl)
			h := Make(stor)

			stor.EXPECT().Ping(ctx).Return(ex.storReturned)

			resp, err := h.CallGRPC(ctx)
			if ex.respError == nil {
				assert.NoError(t, err)
				assert.Equal(t, &pb.PingResponse{}, resp)
			} else {
				assert.EqualError(t, err, ex.respError.Error())
			}
		})

	}
}
