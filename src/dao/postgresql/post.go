package postgresql

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"log"
	"sort"
	"srbbs/src/model"
	"srbbs/src/query"
	"srbbs/src/srlogger"
	"strconv"
	"time"
)

func CreatePost(post *model.Post) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = query.Post.WithContext(ctx).Create(post)
	if err != nil {
		srlogger.Logger().Error("Error creating post", zap.Error(err))
	}
	return
}

func GetPostList(page, size int) ([]*model.Post, error) {
	p := query.Post
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := query.Post.WithContext(ctx).Order(p.CreatedAt).Offset((page - 1) * size).Limit(size).Find()
	if err != nil {
		srlogger.Logger().Error("Error getting posts", zap.Error(err))
		return nil, err
	}
	return results, err
}

func GetPostTotalCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return query.Post.WithContext(ctx).Count()
}

func GetPostListByStrIDs(ids []string) ([]*model.Post, error) {
	idsInt := make([]int64, len(ids))
	for i, v := range ids {
		intV, _ := strconv.Atoi(v)
		idsInt[i] = int64(intV)
	}
	return GetPostListByIDs(idsInt)
}

func GetPostByID(id int64) (*model.Post, error) {
	p := query.Post
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	first, err := query.Post.WithContext(ctx).Where(p.PostID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return first, err
}

func GetPostListByIDs(ids []int64) ([]*model.Post, error) {
	p := query.Post
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := query.Post.WithContext(ctx).Where(p.PostID.In(ids...)).Find()

	indexOfID := func(ids []int64, id int64) int {
		for i, v := range ids {
			if v == id {
				return i
			}
		}
		return -1
	}
	sort.Slice(results, func(i, j int) bool {
		return indexOfID(ids, results[i].ID) < indexOfID(ids, results[j].PostID)
	})

	//results, err := query.Post.WithContext(ctx).Where(p.PostID.In(ids...)).Clauses(clause.OrderBy{
	//	Expression: clause.Expr{
	//		SQL:                "FIELD(post_id,?)",
	//		Vars:               []interface{}{ids},
	//		WithoutParentheses: true,
	//	},
	//}).Find()

	if err != nil {
		srlogger.Logger().Error("Error getting posts", zap.Error(err))
		return nil, err
	}
	return results, err
}

func GetPostListByKeywords(param *model.C2SGetPostParam) ([]*model.Post, error) {
	//todo 数据库模糊搜索改为Elasticsearch

	p := query.Post
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//, p.Content.Like("%" + param.Search + "%")
	res, err := query.Post.WithContext(ctx).Where(p.Title.Like("%" + param.Search + "%")).Or(p.Content.Like("%" + param.Search + "%")).Order(p.CreatedAt).Offset(int((param.Page - 1) * param.Size)).Limit(int(param.Size)).Find()
	log.Println("sql", res, err)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetPostListTotalCount(param *model.C2SGetPostParam) (int64, error) {
	p := query.Post
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := query.Post.WithContext(ctx).Where(p.Title.Like(param.Search), p.Content.Like(param.Search)).Count()
	if err != nil {
		return 0, err
	}
	return res, err
}
