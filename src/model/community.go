package model

// ApiCommunityDetail 社区详情model
type ApiCommunityDetail struct {
	CommunityID   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction  string `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	CreateTime    string `json:"create_time" db:"create_time"`
}

func Community2Detail(com *Community) *ApiCommunityDetail {
	if com == nil {
		return nil
	}
	return &ApiCommunityDetail{
		CommunityID:   uint64(com.CommunityID),
		CommunityName: com.CommunityName,
		Introduction:  com.Introduction,
		CreateTime:    com.CreatedAt.String(),
	}
}
