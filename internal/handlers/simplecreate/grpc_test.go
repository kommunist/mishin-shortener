package simplecreate

import (
	"context"
	"errors"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/errors/exist"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	exList := []struct {
		name      string
		ctx       context.Context
		storTimes int
		storErr   error
		respErr   error
		respBody  string
	}{
		{
			name:      "post_simple_create_happy_path",
			ctx:       context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			storTimes: 1,
			storErr:   nil,
			respErr:   nil,
			respBody:  "931691969b142b3a0f11a03e36fcc3b7",
		},
		{
			name:      "post_simple_create_when_without_user_in_context",
			ctx:       context.Background(),
			storTimes: 0,
			storErr:   nil,
			respErr:   status.Error(codes.Unknown, "Error with auth"),
			respBody:  "",
		},
		{
			name:      "post_simple_create_when_record_exist_in_db",
			ctx:       context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			storTimes: 1,
			storErr:   exist.NewExistError(nil),
			respErr:   status.Error(codes.AlreadyExists, "Error when call service"),
			respBody:  "931691969b142b3a0f11a03e36fcc3b7",
		},
		{
			name:      "post_simple_create_when_another_error",
			ctx:       context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			storTimes: 1,
			storErr:   errors.New("another"),
			respErr:   status.Error(codes.Unknown, "Error when call service"),
			respBody:  "",
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

			stor.EXPECT().Push(
				ex.ctx,
				"931691969b142b3a0f11a03e36fcc3b7",
				"biba",
				"qq",
			).Times(ex.storTimes).Return(ex.storErr)

			resp, err := h.Create(ex.ctx, &pb.CreateRequest{Original: "biba"})

			if ex.respErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, &pb.CreateResponse{Short: ex.respBody}, resp)
			} else {
				assert.EqualError(t, err, ex.respErr.Error())
			}
		})
	}
}
