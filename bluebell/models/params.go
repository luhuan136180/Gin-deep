package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	//binding:gin框架自带的tag,用于参数校验，required：不能为空，为空返回错误
	Username   string `json:"username"binding:"required"`
	Password   string `json:"password"binding:"required"`
	RePassword string `json:"re_password"binding:"required,eqfield=Password"`
}

//登录参数
type ParamLogin struct {
	Username string `json:"username"binding:"required"`
	Password string `json:"password"binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id"binding:"required"`
	Direction int8   `json:"direction,string"binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       //页码
	Size        int64  `json:"size" form:"size" example:"10"`      //每页数据量
	Order       string `json:"order" form:"order" example:"score"` //排序依据
}

type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
