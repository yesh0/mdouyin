package services

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"context"
	"crypto/md5"
	"fmt"
	"gateway/biz/model/douyin/core"
	"gateway/internal/cache"
	"gateway/internal/db"
	"strings"
)

func FromUser(u *db.UserDO, followed bool) (user *core.User) {
	user = &core.User{
		Id:       int64(u.Id),
		Name:     u.Nickname,
		IsFollow: followed,
		Avatar: fmt.Sprintf("https://cravatar.cn/avatar/%x",
			md5.Sum([]byte(strings.ToLower(u.Name)))),
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

func getMappedUsers(userMap map[int64]*core.User, users []db.UserDO) map[int64]*core.User {
	for _, user := range users {
		userMap[user.Id] = FromUser(&user, false)
	}
	return userMap
}

func fillUserCounts(ctx context.Context, ids []int64, userMap map[int64]*core.User) error {
	counts, err := Counter.Fetch(ctx, &rpc.CounterGetRequest{
		Id: ids,
		Kinds: []int8{
			common.KindUserFollowerCount,
			common.KindUserFollowingCount,
			common.KindUserFavoriteCount,
			common.KindUserWorkCount,
		},
	})
	if err != nil {
		return err
	}

	for _, count := range counts.Counters {
		id := count.Id
		user := userMap[id]
		if user == nil {
			continue
		}
		fillKindCounts(count, user)
		cache.SetUser(id, user)
	}
	return nil
}

func fillKindCounts(count *rpc.Counts, user *core.User) {
	for _, kindCount := range count.KindCounts {
		i := int64(kindCount.Count)
		switch kindCount.Kind {
		case common.KindUserFollowerCount:
			user.FollowerCount = &i
		case common.KindUserFollowingCount:
			user.FollowCount = &i
		case common.KindUserFavoriteCount:
			user.FavoriteCount = &i
		case common.KindUserWorkCount:
			user.WorkCount = &i
		}
	}
}

func fillRelation(ctx context.Context, ids []int64, user int64, userMap map[int64]*core.User) error {
	var list []int64
	if list = cache.GetFollowing(user); list == nil {
		r, err := Feed.Following(ctx, &rpc.DouyinRelationFollowListRequest{
			UserId:        user,
			RequestUserId: user,
		})
		if err != nil {
			return err
		}
		list = r.UserList
		cache.SetFollowing(user, list)
	}
	for _, following := range list {
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
	userMap := make(map[int64]*core.User)
	coldIds := make([]int64, 0, len(ids)/2)
	for _, id := range ids {
		if u := cache.GetUser(id); u != nil {
			if !counts || (u.FollowCount != nil &&
				u.FollowerCount != nil &&
				u.FavoriteCount != nil &&
				u.WorkCount != nil) {
				userMap[id] = u
				continue
			}
		}
		coldIds = append(coldIds, id)
	}

	basicUsers, err := db.FindUsersByIds(coldIds)
	if err != nil {
		return nil, err
	}

	userMap = getMappedUsers(userMap, basicUsers)

	if counts {
		if err := fillUserCounts(ctx, coldIds, userMap); err != nil {
			return nil, err
		}
	} else {
		fillCacheIfMissing(coldIds, userMap)
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

func fillCacheIfMissing(ids []int64, userMap map[int64]*core.User) {
	for _, id := range ids {
		if cache.GetUser(id) == nil {
			cache.SetUser(id, userMap[id])
		}
	}
}
