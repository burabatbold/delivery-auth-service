package adminGrpcService

import (
	"context"

	protos "github.com/burabatbold/delivery-auth-service/grpc/protos/admin-service/grpc/protos"
	adminDto "github.com/burabatbold/delivery-auth-service/modules/admin/dto"
	adminUsecase "github.com/burabatbold/delivery-auth-service/modules/admin/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminService struct {
	protos.UnimplementedAdminServiceServer
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (h *AdminService) Login(ctx context.Context, req *protos.LoginRequest) (*protos.LoginResponse, error) {

	admin, err := adminUsecase.NewAdminAuthUsecase().Login(adminDto.LoginDto{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &protos.LoginResponse{
		Token: admin.Token,
		Admin: &protos.AdminInfo{
			Id:       uint64(admin.Admin.ID),
			Username: admin.Admin.Email,
			Email:    admin.Admin.Email,
		},
	}, nil
}

func (h *AdminService) VerifyToken(ctx context.Context, req *protos.VerifyTokenRequest) (*protos.VerifyTokenResponse, error) {

	token := req.Token

	err := adminUsecase.NewAdminAuthUsecase().VerifyToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &protos.VerifyTokenResponse{
		Valid: true,
	}, nil
}
