package services

import (
	"common"
	"common/kitex_gen/douyin/rpc"
	"context"
	"encoding/binary"
	"gateway/biz/model/douyin/core"
	"gateway/internal/cache"
	"gateway/internal/db"
	"hash/fnv"
)

var temporary_background = "http://rollinggirls.com/img/bg_intro.jpg"
var temporary_quote = "「生存戦略、しましょうか。」"
var temporary_avatars = []string{
	"http://yurikuma.jp/images/special/008/yk_icon_kuma1.png",
	"http://yurikuma.jp/images/special/008/yk_icon_kuma2.png",
	"http://yurikuma.jp/images/special/008/yk_icon_kuma3.png",
	"http://yurikuma.jp/images/special/006/icon_sexy_bear.png",
	"http://yurikuma.jp/images/special/006/icon_cool_bear.png",
	"http://yurikuma.jp/images/special/006/icon_beauty_bear.png",
	"http://yurikuma.jp/images/special/005/icon_ginko_winning.png",
	"http://yurikuma.jp/images/special/005/icon_lulu_winning.png",
	"http://yurikuma.jp/images/special/004/icon_sexy.png",
	"http://yurikuma.jp/images/special/004/icon_cool.png",
	"http://yurikuma.jp/images/special/004/icon_beauty.png",
	"http://yurikuma.jp/images/special/003/icon_ginko_bear.png",
	"http://yurikuma.jp/images/special/003/icon_lulu_bear.png",
	"http://yurikuma.jp/images/special/001/icon_ginko.png",
	"http://yurikuma.jp/images/special/001/icon_lulu.png",
	"http://yurikuma.jp/images/special/001/icon_kureha.png",
}

func FromUser(u *db.UserDO, followed bool) (user *core.User) {
	bytes := [8]byte{}
	binary.LittleEndian.PutUint64(bytes[:], uint64(u.Id))
	hasher := fnv.New32()
	hasher.Write(bytes[:])
	hash := hasher.Sum32()
	user = &core.User{
		Id:              u.Id,
		Name:            u.Nickname,
		IsFollow:        followed,
		Avatar:          temporary_avatars[hash%uint32(len(temporary_avatars))],
		BackgroundImage: &temporary_background,
		Signature:       &temporary_quote,
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
			common.KindUserTotalFavorited,
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
		case common.KindUserTotalFavorited:
			user.TotalFavorited = &i
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
				u.TotalFavorited != nil &&
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
