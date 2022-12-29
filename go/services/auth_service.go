package services

import (
	"net/http"
	"pxj/CloudTravelShopApi/go/models"
)

type AdminUserLoginByAccountRequest struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Type      string `json:"type"`
	AutoLogin bool   `json:"autoLogin"`
}

type AdminUserLoginByAccountResponse struct {
	Token            string   `json:"token"`
	Status           string   `json:"status"`
	Type             string   `json:"type"`
	Name             string   `json:"name"`
	CurrentAuthority []string `json:"currentAuthority"`
	Data             struct {
		Name string `json:"name"`
	} `json:"data"`
}

func AdminUserLoginByAccount(req *AdminUserLoginByAccountRequest) (resp *AdminUserLoginByAccountResponse, code int, err error) {
	resp = &AdminUserLoginByAccountResponse{}
	adminUser := models.NewAdminUser()
	adminUser.Name = req.UserName
	adminUser.Password = req.Password
	err = adminUser.FirstOrCreate()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	resp.Type = "account"
	resp.Token = "123456"
	resp.Status = "ok"
	resp.Name = "ok"
	resp.Data.Name = "ok"
	resp.CurrentAuthority = []string{
		"超级管理员",
	}
	return resp, http.StatusOK, nil
}

type CurrentUserRequest struct {
}

type CurrentUserResponse struct {
	Name        string   `json:"token"`
	Permissions []string `json:"status"`
	Roles       []string `json:"type"`
}

func CurrentUser(req *CurrentUserRequest) (resp *CurrentUserResponse, code int, err error) {
	resp = &CurrentUserResponse{}
	adminUser, err := models.GetAdminUserById(1)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	resp.Name = adminUser.Name
	return resp, http.StatusOK, nil
}
