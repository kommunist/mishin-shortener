package deleteurls

import (
	"context"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/secure"
	pb "mishin-shortener/proto"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCallGRPC(t *testing.T) {
	exList := []struct {
		name   string
		ctx    context.Context
		input  *pb.DeleteUrlsRequest
		err    error
		status int
	}{
		{
			name: "Start_DELETE_to_delete_record_in_db",
			ctx:  context.WithValue(context.Background(), secure.UserIDKey, "userId"),
			input: &pb.DeleteUrlsRequest{
				List: []*pb.DeleteUrlsRequestItem{
					{Short: "first"}, {Short: "second"},
				},
			},
			err: nil,
		},
		{
			name: "Start_DELETE_to_delete_record_in_db_without_user",
			ctx:  context.Background(),
			input: &pb.DeleteUrlsRequest{
				List: []*pb.DeleteUrlsRequestItem{
					{Short: "first"}, {Short: "second"},
				},
			},
			err: status.Error(codes.Unknown, "Error with auth"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {

			c := config.MakeConfig()
			c.InitConfig()
			h := Make(make(chan delasync.DelPair, 5))

			resp, err := h.CallGRPC(ex.ctx, ex.input)
			if ex.err != nil {
				assert.EqualError(t, err, ex.err.Error())
			} else {
				assert.Equal(t, &pb.DeleteUrlsResponse{}, resp)
				assert.NoError(t, err)
			}
		})
	}
}
