package model

import (
	"encoding/json"
	"errors"
)

// C2SGetPostParam 获取帖子列表query 参数
type C2SGetPostParam struct {
	Search      string `json:"search" form:"search"`               // 关键字搜索
	CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

type S2CPostList struct {
	Page Page             `json:"page"`
	List []*ApiPostDetail `json:"list"`
}

type ApiPostDetail struct {
	*Post                                  // 嵌入帖子结构体
	*ApiCommunityDetail `json:"community"` // 嵌入社区信息
	AuthorName          string             `json:"author_name"`
	VoteNum             int64              `json:"vote_num"` // 投票数量
	//CommunityName string `json:"community_name"`
}

type Page struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

// UnmarshalJSON 为Post类型实现自定义的UnmarshalJSON方法
func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" db:"title"`
		Content     string `json:"content" db:"content"`
		CommunityID int64  `json:"community_id" db:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = int32(required.CommunityID)
	}
	return
}
