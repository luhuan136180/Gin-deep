package logic

import (
	"bluebell0002/dao/mysql"
	"bluebell0002/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
