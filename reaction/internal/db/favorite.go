package db

import (
	"common/snowy"
	"common/utils"
)

type FavoriteDO struct {
	Id     int64 `gorm:"<-:create;primaryKey;autoIncrement=false;index:idx_time,priority:2"`
	Author int64 `gorm:"uniqueIndex:uk_video;index:idx_time,priority:1"`
	Video  int64 `gorm:"uniqueIndex:uk_video"`
}

func ListFavorites(user int64, limit int) (favorites []int64, err utils.ErrorCode) {
	result := db.Model(&FavoriteDO{}).Select("video").
		Where("author = ?", user).Limit(limit).Find(&favorites)
	if result.Error != nil {
		err = utils.ErrorDatabaseError
		return
	}
	return
}

func Favorite(user int64, video int64) utils.ErrorCode {
	favorite := FavoriteDO{
		Id:     snowy.ID(),
		Author: user,
		Video:  video,
	}
	if err := db.Create(&favorite).Error; err != nil {
		return utils.ErrorRepeatedRequests
	}
	return utils.ErrorOk
}

func Unfavorite(user int64, video int64) utils.ErrorCode {
	result := db.Where("author = ?", user).Where("video = ?", video).Delete(&FavoriteDO{})
	if result.Error != nil {
		return utils.ErrorDatabaseError
	}
	if result.RowsAffected == 0 {
		return utils.ErrorRepeatedRequests
	}
	return utils.ErrorOk
}

func IsFavorite(user int64, videos []int64) (favorites []int8, err utils.ErrorCode) {
	ids := make([]int64, 0)
	result := db.Model(&FavoriteDO{}).Select("video").
		Where("author = ?", user).Where("video IN ?", videos).Find(&ids)
	if result.Error != nil {
		err = utils.ErrorDatabaseError
		return
	}

	set := make(map[int64]struct{})
	for _, id := range ids {
		set[id] = struct{}{}
	}

	favorites = make([]int8, 0, len(videos))
	for _, v := range videos {
		if _, ok := set[v]; ok {
			favorites = append(favorites, 1)
		} else {
			favorites = append(favorites, 0)
		}
	}
	return
}
