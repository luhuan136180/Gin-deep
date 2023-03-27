package models

import "time"

type Community struct {
	ID   int64  `json:"id,string"db:"community_id"`
	Name string `json:"name"db:"community_name"`
}

type CommunityDetail struct {
	ID   int64  `json:"id,string"db:"community_id"`
	Name string `json:"name"db:"community_name"`
	//详情
	Introduction string `json:"introduction,omitempty"db:"introduction"`
	//创建时间
	CreateTime time.Time `json:"create_time"db:"create_time"`
	//

}
