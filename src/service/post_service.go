package service

import (
	"srbbs/src/dao/postgresql"
	"srbbs/src/dao/redis"
	"srbbs/src/model"
	"srbbs/src/util/lib/algo/snowflake"
)

func CreatePost(userid string, post *model.Post) (err error) {
	postId, err := snowflake.GetID()
	if err != nil {
		return
	}
	post.PostID = int64(postId)
	if err = postgresql.CreatePost(post); err != nil {
		return
	}

	// postgres
	if err = postgresql.CreatePost(post); err != nil {
		return
	}
	//验证社区存在
	communiey, err := postgresql.GetCommunityById(post.CommunityID)
	if err != nil {
		return
	}
	// redis
	err = redis.CreatePost(post.PostID, post.AuthorID, int64(communiey.ID), post.Title, post.Content)
	return
}
