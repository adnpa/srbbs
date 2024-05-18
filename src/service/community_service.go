package service

import (
	"srbbs/src/dao/postgresql"
	"srbbs/src/model"
)

func GetCommunityList() (res []*model.ApiCommunityDetail, err error) {
	comL, err := postgresql.GetCommunityList()
	if err != nil {
		return nil, err
	}

	res = make([]*model.ApiCommunityDetail, len(comL))
	for _, com := range comL {
		res = append(res, model.Community2Detail(com))
	}

	return

}

func GetCommunityDetailByID(id int) (*model.ApiCommunityDetail, error) {
	com, err := postgresql.GetCommunityDetailByID(int32(id))
	if err != nil {
		return nil, err
	}

	return model.Community2Detail(com), nil
}
