package logic

import (
	"bluebell0002/dao/mysql"
	"bluebell0002/dao/redis"
	"bluebell0002/models"
	"bluebell0002/pkg/snowFlake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowFlake.GenID()
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}
func GetPostById(id int64) (data *models.ApiPostDetail, err error) {
	//查询并拼接组合我们接口想用的数据
	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostById(ID) failed",
			zap.Int64("pid", id),
			zap.Error(err))
		return
	}
	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return

	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	//接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Name,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
