package createjsonbatch

import (
	"context"
	"errors"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCallGrpc(t *testing.T) {
	exList := []struct {
		name          string
		input         *pb.CreateBatchRequest
		ctx           context.Context
		userID        string
		responseItems []responseBatchItem
		storTimes     int
		err           error
		storErr       error
	}{
		{
			name: "Start_POST_to_create_record_in_storage",
			input: &pb.CreateBatchRequest{
				List: []*pb.CreateBatchRequestItem{
					{CorrelationId: "123", OriginalUrl: "biba"},
					{CorrelationId: "456", OriginalUrl: "boba"},
				},
			},
			ctx:    context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID: "qq",
			err:    nil,
			responseItems: []responseBatchItem{
				{CorrelationID: "123", ShortURL: "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7"},
				{CorrelationID: "456", ShortURL: "http://localhost:8080/2cce0ec300cfe8dd3024939db0448893"},
			},
			storTimes: 1,
			storErr:   nil,
		},
		{
			name: "Start_POST_to_create_record_in_storage_without_user_in_context",
			input: &pb.CreateBatchRequest{
				List: []*pb.CreateBatchRequestItem{
					{CorrelationId: "123", OriginalUrl: "biba"},
					{CorrelationId: "456", OriginalUrl: "boba"},
				},
			},
			ctx:           context.Background(),
			userID:        "qq",
			err:           status.Error(codes.Unknown, "Error with auth"),
			responseItems: nil,
			storTimes:     0,
			storErr:       nil,
		},
		{
			name: "Start_POST_to_create_record_in_storage_but_error_happen",
			input: &pb.CreateBatchRequest{
				List: []*pb.CreateBatchRequestItem{
					{CorrelationId: "123", OriginalUrl: "biba"},
					{CorrelationId: "456", OriginalUrl: "boba"},
				},
			},
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			err:           status.Error(codes.Unknown, "Error when call service"),
			responseItems: nil,
			storTimes:     1,
			storErr:       errors.New("qq"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockPusher(ctrl)

			c := config.MakeConfig()
			c.InitConfig()

			h := Make(c, stor)

			stor.EXPECT().PushBatch(
				ex.ctx,
				&map[string]string{
					"931691969b142b3a0f11a03e36fcc3b7": "biba", "2cce0ec300cfe8dd3024939db0448893": "boba",
				},
				ex.userID,
			).Times(ex.storTimes).Return(ex.storErr)

			resp, err := h.CallGRPC(ex.ctx, ex.input)
			if err != nil {
				assert.EqualError(t, err, ex.err.Error())

			} else {
				assert.NoError(t, err, "it_received_without_err")

				list := make([]responseBatchItem, 0, len(resp.List))
				for _, item := range resp.List {
					list = append(
						list,
						responseBatchItem{
							CorrelationID: item.CorrelationId, ShortURL: item.ShortUrl,
						},
					)
				}
				assert.Equal(t, ex.responseItems, list)

			}
		})
	}
}
