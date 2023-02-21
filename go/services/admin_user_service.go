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
	Page     uint   `json:"page"`
	PageSize uint   `json:"page_size"`
}

type AdminUserGetResponse struct {
	AdminUsers []models.AdminUser `json:"admin_users"`
	//Page        uint             `json:"page"`
	//PageSize    uint             `json:"page_size"`
	//CurrentPage uint             `json:"current_page"`
	//Data        any              `json:"data"`
}

func AdminUserGet(req *AdminUserGetRequest) (resp *Response, code int, err error) {
	//data := &AdminUserGetResponse{}
	//data.AdminUsers = make([]models.AdminUser, 0)
	adminUsers := make([]models.AdminUser, 0)
	database.Db.Model(models.AdminUser{}).Find(&adminUsers)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	resp = (&Response{}).setData(adminUsers)
	return resp, http.StatusOK, nil

}
