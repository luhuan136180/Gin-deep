package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err = Db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	//初始化
	community = new(models.CommunityDetail)
	//初始化sql语句
	//查询指定列需要参数，此处用id搜索
	sqlStr := "select " +
		"community_id,community_name,introduction,create_time " +
		"from community " +
		"where community_id=?"
	if err = Db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID //id有误
		}
	}
	return
}
