package postgresql

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/model"
	"srbbs/src/query"
	"srbbs/src/srlogger"
	"time"
)

func GetCommunityById(id int) (*model.Community, error) {
	u := query.Community
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	first, err := query.Community.WithContext(ctx).Where(u.CommunityID.Eq(int32(id))).First()
	if err != nil {
		srlogger.Logger().Warn("get community by id err", zap.Error(err))
		return nil, err
	}
	return first, err
}

func CreateCommunity(community *model.Community) (*model.Community, error) {
	return nil, nil
}

//func CreateComm

func GetCommunityPostTotalCount(cid int) (int, error) {

	return 0, nil
}

func GetCommunityByID(cid int) (*model.Community, error) {
	return nil, nil
}

func GetCommunityList() ([]*model.Community, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := query.Community.WithContext(ctx).Find()
	if err != nil {
		srlogger.Logger().Error("Error GetCommunityList", zap.Error(err))
		return nil, err
	}
	return results, err
}

func GetCommunityDetailByID(id int32) (*model.Community, error) {
	c := query.Community
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results, err := query.Community.WithContext(ctx).Where(c.CommunityID.Eq(id)).First()
	if err != nil {
		srlogger.Logger().Error("Error GetCommunityDetailByID", zap.Error(err))
		return nil, err
	}
	return results, err
}
