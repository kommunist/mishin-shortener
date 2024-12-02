package userurls

import (
	"context"
	"errors"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPC(t *testing.T) {
	exList := []struct {
		name       string
		ctx        context.Context
		storTimes  int
		storResult map[string]string
		storError  error
		respError  error
		resp       *pb.UserUrlsResponse
	}{
		{
			name:       "post_to_create_record_in_db",
			ctx:        context.WithValue(context.Background(), secure.UserIDKey, "userId"),
			storTimes:  1,
			storResult: map[string]string{"short0": "long0", "short1": "long1"},
			storError:  nil,
			resp: &pb.UserUrlsResponse{
				List: []*pb.UserUrlsResponseItem{
					{
						Short:    "http://localhost:8080/short0",
						Original: "long0",
					},
					{
						Short:    "http://localhost:8080/short1",
						Original: "long1",
					},
				},
			},
			respError: nil,
		},
		{
			name:       "post_to_create_record_in_db_without_userID_in_context",
			ctx:        context.Background(),
			storTimes:  0,
			storResult: nil,
			storError:  nil,
			resp:       nil,
			respError:  status.Error(codes.Unknown, "Error with auth"),
		},
		{
			name:       "post_to_create_record_in_db_when_error_from_database",
			ctx:        context.WithValue(context.Background(), secure.UserIDKey, "userId"),
			storTimes:  1,
			storResult: nil,
			storError:  errors.New("some error"),
			resp:       nil,
			respError:  status.Error(codes.Unknown, "Error when call service"),
		},
		{
			name:       "post_to_create_record_in_db_when_empty_result_from_db",
			ctx:        context.WithValue(context.Background(), secure.UserIDKey, "userId"),
			storTimes:  1,
			storResult: map[string]string{},
			storError:  nil,
			resp:       nil,
			respError:  status.Error(codes.NotFound, "No data"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			stor := NewMockByUserIDGetter(ctrl)

			stor.EXPECT().GetByUserID(
				ex.ctx,
				"userId",
			).Times(ex.storTimes).Return(ex.storResult, ex.storError)

			c := config.MakeConfig()
			c.InitConfig()
			h := Make(c, stor)

			resp, err := h.CallGRPC(ex.ctx, &pb.UserUrlsRequest{})

			if ex.respError == nil {
				assert.NoError(t, err)
				assert.Equal(t, ex.resp, resp)
			} else {
				assert.EqualError(t, err, ex.respError.Error())
			}

		})
	}
}
