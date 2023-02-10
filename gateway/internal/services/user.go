package services

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"context"
	"crypto/md5"
	"fmt"
	"gateway/biz/model/douyin/core"
	"gateway/internal/db"
	"strings"
)

func FromUser(u *db.UserDO, counts []*rpc.Counts, followed bool) (user *core.User) {
	user = &core.User{
		Id:       int64(u.Id),
		Name:     u.Nickname,
		IsFollow: followed,
		Avatar: fmt.Sprintf("https://cravatar.cn/avatar/%x",
			md5.Sum([]byte(strings.ToLower(u.Name)))),
	}
	for _, c := range counts {
		if c.Id == u.Id {
			for _, kind := range c.KindCounts {
				switch kind.Kind {
				case common.KindUserFollowerCount:
					followers := int64(kind.Count)
					user.FollowerCount = &followers
				case common.KindUserFollowingCount:
					following := int64(kind.Count)
					user.FollowCount = &following
				}
			}
		}
	}
	return
}

func getIds(users []*rpc.User) []int64 {
	dedup := make(map[int64]struct{})
	ids := make([]int64, 0)
	for _, user := range users {
		if _, ok := dedup[user.Id]; !ok {
			dedup[user.Id] = struct{}{}
			ids = append(ids, user.Id)
		}
	}
	return ids
}

func getMappedUsers(users []db.UserDO) (userMap map[int64]*core.User) {
	userMap = make(map[int64]*core.User)
	for _, user := range users {
		userMap[user.Id] = FromUser(&user, nil, false)
	}
	return
}

func fillCounts(ctx context.Context, ids []int64, userMap map[int64]*core.User) error {
	counts, err := Counter.Fetch(ctx, &rpc.CounterGetRequest{
		Id:    ids,
		Kinds: []int8{common.KindUserFollowerCount, common.KindUserFollowingCount},
	})
	if err != nil {
		return err
	}

	for _, count := range counts.Counters {
		id := count.Id
		for _, kindCount := range count.KindCounts {
			if userMap[id] == nil {
				continue
			}
			i := int64(kindCount.Count)
			switch kindCount.Kind {
			case common.KindUserFollowerCount:
				userMap[id].FollowerCount = &i
			case common.KindUserFollowingCount:
				userMap[id].FollowCount = &i
			}
		}
	}
	return nil
}

func fillRelation(ctx context.Context, ids []int64, user int64, userMap map[int64]*core.User) error {
	r, err := Feed.Following(ctx, &rpc.DouyinRelationFollowListRequest{
		UserId:        user,
		RequestUserId: user,
	})
	if err != nil {
		return err
	}
	for _, following := range r.UserList {
		if userMap[following] == nil {
			continue
		}
		if _, ok := userMap[following]; ok {
			userMap[following].IsFollow = true
		}
	}
	return nil
}

func GatherUserInfo(ctx context.Context, user int64, users []*rpc.User,
	counts bool, follow bool) (map[int64]*core.User, error) {
	ids := getIds(users)
	return GatherUserInfoFromIds(ctx, user, ids, users, counts, follow)
}

func GatherUserInfoFromIds(ctx context.Context, user int64,
	ids []int64, users []*rpc.User, counts bool, follow bool) (map[int64]*core.User, error) {
	basicUsers, err := db.FindUsersByIds(ids)
	if err != nil {
		return nil, err
	}

	userMap := getMappedUsers(basicUsers)

	if counts {
		if err := fillCounts(ctx, ids, userMap); err != nil {
			return nil, err
		}
	}

	if users == nil || (follow && user != 0) {
		if err := fillRelation(ctx, ids, user, userMap); err != nil {
			return nil, err
		}
	} else {
		for _, user := range users {
			if user.IsFollow {
				if userMap[user.Id] == nil {
					continue
				}
				userMap[user.Id].IsFollow = true
			}
		}
	}

	return userMap, nil
}
