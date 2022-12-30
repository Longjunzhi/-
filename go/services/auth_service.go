package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"pxj/CloudTravelShopApi/go/models"
	"pxj/CloudTravelShopApi/go/utils"
	"strconv"
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
	CurrentAuthority []string `json:"currentAuthority"`
	Data             struct {
		Name string `json:"name"`
	} `json:"data"`
}

func AdminUserLoginByAccount(req *AdminUserLoginByAccountRequest) (resp *AdminUserLoginByAccountResponse, code int, err error) {
	resp = &AdminUserLoginByAccountResponse{}
	adminUser := models.NewAdminUser()
	err = models.Db.First(&adminUser, "name = ?", req.UserName).Error
	fmt.Printf("adminUser: %+v", adminUser)
	err = bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(req.Password))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	token, err := utils.GenerateJwtToken(strconv.Itoa(int(adminUser.ID)))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	resp.Type = "account"
	resp.Token = token
	resp.Status = "ok"
	resp.Data.Name = adminUser.Name
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
