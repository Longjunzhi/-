package services

import (
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

func LoginByPassword(req *LoginByPasswordRequest) (resp *LoginByPasswordResponse, code int, err error) {
	adminUser := models.NewAdminUser()
	adminUser.Name = req.UserName
	adminUser.Password = req.Password
	err = adminUser.FirstOrCreate()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return resp, http.StatusOK, nil
}
