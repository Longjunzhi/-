package services

import (
	"net/http"
	"pxj/CloudTravelShopApi/go/database"
	"pxj/CloudTravelShopApi/go/models"
)

type AdminUserGetRequest struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LoginIp  string `json:"login_ip"`
	ParentId uint   `json:"parent_id"`
}

type AdminUserGetResponse struct {
	Page        uint `json:"page"`
	PageSize    uint `json:"page_size"`
	CurrentPage uint `json:"current_page"`
	Data        any  `json:"data"`
}

func AdminUserGet(req *CurrentUserRequest) (resp *Response, code int, err error) {
	data := &AdminUserGetResponse{}
	database.Db.Model(models.AdminUser{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	resp = (&Response{}).setData(data)
	return resp, http.StatusOK, nil

}
