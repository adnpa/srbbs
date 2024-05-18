package service

import (
	"go.uber.org/zap"
	"srbbs/src/dao/postgresql"
	"srbbs/src/dao/redis"
	"srbbs/src/model"
	"srbbs/src/util/lib/algo/snowflake"
	"strconv"
)

func CreatePost(userid string, post *model.Post) (err error) {
	//1 验证community存在
	_, err = postgresql.GetCommunityById(int(post.CommunityID))
	if err != nil {
		return
	}

	//2 生成帖子id
	postId, err := snowflake.GetID()
	if err != nil {
		return
	}
	post.PostID = int64(postId)

	//4 写入postgres
	if err = postgresql.CreatePost(post); err != nil {
		return
	}

	// 5 写入redis
	err = redis.CreatePost(post.PostID, post.AuthorID, int64(post.CommunityID), post.Title, post.Content)
	return
}

func GetPostById(postId int64) (postRes *model.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	// 查询帖子信息
	post, err := postgresql.GetPostByID(postId)
	if err != nil {
		return nil, err
	}
	// 根据作者id查询作者信息
	user, err := postgresql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("postgresql.GetUserByID() failed",
			zap.Uint64("postID", uint64(post.AuthorID)),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := postgresql.GetCommunityByID(int(post.CommunityID))
	if err != nil {
		zap.L().Error("postgresql.GetCommunityByID() failed",
			zap.Uint64("community_id", uint64(post.CommunityID)),
			zap.Error(err))
		return
	}
	// 根据帖子id查询帖子的投票数
	voteNum := redis.GetPostVoteNum(int(post.PostID))

	// 接口数据拼接
	postRes = &model.ApiPostDetail{
		Post:               post,
		ApiCommunityDetail: model.Community2Detail(community),
		AuthorName:         user.Username,
		VoteNum:            int64(voteNum),
	}
	return
}

//func GetPostList(page, size int) ([]*model.ApiPostDetail, error) {
//	postList, err := postgresql.GetPostList(page, size)
//	if err != nil {
//		return nil, err
//	}
//	log.Println("postList", postList)
//	data := make([]*model.ApiPostDetail, 0, len(postList)) // data 初始化
//	for _, post := range postList {
//		// 根据作者id查询作者信息
//		user, err := postgresql.GetUserById(post.AuthorID)
//		if err != nil {
//			zap.L().Error("postgresql.GetUserByID() failed",
//				zap.Uint64("postID", uint64(post.AuthorID)),
//				zap.Error(err))
//			continue
//		}
//		// 根据社区id查询社区详细信息
//		community, err := postgresql.GetCommunityById(int(post.CommunityID))
//		if err != nil {
//			zap.L().Error("postgresql.GetCommunityByID() failed",
//				zap.Uint64("community_id", uint64(post.CommunityID)),
//				zap.Error(err))
//			continue
//		}
//		// 接口数据拼接
//		postDetail := &model.ApiPostDetail{
//			Post:               post,
//			ApiCommunityDetail: model.Community2Detail(*community),
//			AuthorName:         user.Username,
//		}
//		data = append(data, postDetail)
//	}
//	log.Println("pr", data)
//	return data, nil
//}

func GetPostListNew(p *model.C2SGetPostParam) (res *model.S2CPostList, err error) {
	if p.CommunityID == 0 {
		res, err = GetPostList2(p)
	} else {
		res, err = GetCommunityPostList(p)
	}
	return
}

// GetPostList
func GetPostList2(p *model.C2SGetPostParam) (res *model.S2CPostList, err error) {
	res = &model.S2CPostList{}
	// 从postgresql获取帖子列表总数
	total, err := postgresql.GetPostTotalCount()
	if err != nil {
		return
	}
	res.Page.Total = total
	// 1、根据参数中的排序规则去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return res, nil
	}

	// 2、提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := postgresql.GetPostListByStrIDs(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Page = p.Page
	res.Page.Size = p.Size
	res.List = make([]*model.ApiPostDetail, 0, len(posts))
	// 4、组合数据
	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, _ := postgresql.GetUserById(post.AuthorID)
		// 根据社区id查询社区详细信息
		community, _ := postgresql.GetCommunityById(int(post.CommunityID))
		apiCommunity := model.Community2Detail(community)

		// 接口数据拼接
		postDetail := &model.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			ApiCommunityDetail: apiCommunity,
			AuthorName:         user.Username,
		}
		res.List = append(res.List, postDetail)
	}
	return res, nil
}

func GetCommunityPostList(p *model.C2SGetPostParam) (res *model.S2CPostList, err error) {
	res = &model.S2CPostList{}
	// 从postgresql获取该社区下帖子列表总数
	total, err := postgresql.GetCommunityPostTotalCount(int(p.CommunityID))
	if err != nil {
		return nil, err
	}
	res.Page.Total = int64(total)
	// 1、根据参数中的排序规则去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return res, nil
	}
	// 2、提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := postgresql.GetPostListByStrIDs(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Page = p.Page
	res.Page.Size = p.Size
	res.List = make([]*model.ApiPostDetail, 0, len(posts))
	// 4、根据社区id查询社区详细信息
	community, err := postgresql.GetCommunityByID(int(p.CommunityID))
	for idx, post := range posts {
		// 过滤掉不属于该社区的帖子
		if post.CommunityID != int32(p.CommunityID) {
			continue
		}
		// 根据作者id查询作者信息
		user, _ := postgresql.GetUserById(post.AuthorID)
		// 接口数据拼接
		postDetail := &model.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			ApiCommunityDetail: model.Community2Detail(community),
			AuthorName:         user.Username,
		}
		res.List = append(res.List, postDetail)
	}
	return
}

func PostSearch(p *model.C2SGetPostParam) (res *model.S2CPostList, err error) {
	res = &model.S2CPostList{}
	// 根据搜索条件去postgresql查询符合条件的帖子列表总数
	total, err := postgresql.GetPostListTotalCount(p)
	if err != nil {
		return nil, err
	}
	res.Page.Total = total
	// 1、根据搜索条件去postgresql分页查询符合条件的帖子列表
	posts, err := postgresql.GetPostListByKeywords(p)
	if err != nil {
		return nil, err
	}
	// 查询出来的帖子总数可能为0
	if len(posts) == 0 {
		return nil, nil
	}
	// 2、查询出来的帖子id列表传入到redis接口获取帖子的投票数
	ids := make([]string, 0, len(posts))
	for _, post := range posts {
		ids = append(ids, strconv.Itoa(int(post.PostID)))
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Size = p.Size
	res.Page.Page = p.Page
	// 3、拼接数据
	res.List = make([]*model.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, _ := postgresql.GetUserById(post.AuthorID)
		// 根据社区id查询社区详细信息
		community, _ := postgresql.GetCommunityByID(int(post.CommunityID))
		// 接口数据拼接
		postDetail := &model.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			ApiCommunityDetail: model.Community2Detail(community),
			AuthorName:         user.Username,
		}
		res.List = append(res.List, postDetail)
	}
	return
}
