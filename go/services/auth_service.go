package services

import (
	"context"
	"pxj/CloudTravelShopApi/go/models"
)

type LoginByPasswordRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginByPasswordResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func LoginByPassword(ctx context.Context, req *LoginByPasswordRequest) (resp *LoginByPasswordResponse, code int, err error) {
	return resp, 200, nil
}
