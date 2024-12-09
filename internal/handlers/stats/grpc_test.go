package stats

import (
	"context"
	"errors"
	"mishin-shortener/internal/config"
	pb "mishin-shortener/proto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCallGRPC(t *testing.T) {
	exList := []struct {
		name      string
		storTimes int
		storErr   error
		respErr   error
		storUsers int
		storUrls  int
	}{
		{
			name:      "simple_happy_path",
			storTimes: 1,
			storErr:   nil,
			respErr:   nil,
			storUsers: 1,
			storUrls:  1,
		},
		{
			name:      "when_error_from_db",
			storTimes: 1,
			storErr:   errors.New("Ququ"),
			respErr:   status.Error(codes.Unknown, "Error when call service"),
			storUsers: 0,
			storUrls:  0,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			stor := NewMockStatsGetter(ctrl)

			c := config.MakeConfig()
			c.InitConfig()

			h, err := Make(c, stor)
			assert.NoError(t, err)

			ctx := context.Background()

			stor.EXPECT().GetStats(ctx).Times(ex.storTimes).Return(
				ex.storUsers, ex.storUrls, ex.storErr,
			)

			resp, err := h.CallGRPC(ctx)

			if ex.respErr == nil {
				assert.NoError(t, err)
				assert.Equal(
					t,
					&pb.GetStatsResponse{
						CountUrls: uint32(ex.storUrls), CountUsers: uint32(ex.storUsers),
					},
					resp,
				)
			} else {
				assert.EqualError(t, err, ex.respErr.Error())
			}
		})
	}

}
