//go:build userpb

package grpcif

import (
	"context"

	"github.com/lyonnee/go-template/internal/application/queries"
	"github.com/lyonnee/go-template/internal/interfaces/grpc/userpb"
	"github.com/lyonnee/go-template/pkg/di"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 当使用 -tags userpb 构建，且 proto 代码已生成时，启用真实注册
func RegisterGRPCServices(s *grpc.Server) {
	userpb.RegisterUserServiceServer(s, &userServiceImpl{})
}

type userServiceImpl struct {
	qs *queries.UserQueryService
}

func (u *userServiceImpl) ensure() {
	if u.qs == nil {
		u.qs = di.Get[*queries.UserQueryService]()
	}
}

func (u *userServiceImpl) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	u.ensure()
	user, err := u.qs.GetUserById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user failed: %v", err)
	}
	return &userpb.GetUserResponse{
		User: &userpb.UserBasic{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
