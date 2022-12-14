package services

import (
	"context"
	"net/http"
	"pxj/CloudTravelShopApi/go/models"
)

type LoginByPasswordRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type LoginByPasswordResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func LoginByPassword(ctx context.Context, req *LoginByPasswordRequest) (resp *LoginByPasswordResponse, code int, err error) {

	return resp, http.StatusOK, nil
}
