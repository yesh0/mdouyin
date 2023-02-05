package db

import (
	"common/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RelationDO struct {
	Followee  int64 `gorm:"<-:create;primaryKey;autoIncrement=false"`
	Follower  int64 `gorm:"<-:create;primaryKey;autoIncrement=false;index"`
	Mutual    bool
	CreatedAt time.Time
}

// TODO: Cache
func Follow(follower, followee int64) utils.ErrorCode {
	if follower == followee {
		return utils.ErrorUnanticipated
	}

	return utils.From(db.Transaction(func(tx *gorm.DB) error {
		previous := RelationDO{
			Followee: follower,
			Follower: followee,
		}
		mutual := tx.Limit(1).Find(&previous).RowsAffected == 1
		if previous.Mutual {
			return nil
		}

		tx.Error = nil
		if err := tx.Create(&RelationDO{
			Followee: followee,
			Follower: follower,
			Mutual:   mutual,
		}).Error; err != nil {
			return utils.ErrorRepeatedRequests
		}

		if mutual {
			other := &RelationDO{
				Followee: follower,
				Follower: followee,
			}
			if err := tx.Model(other).Select("mutual").
				Updates(RelationDO{Mutual: mutual}).Error; err != nil {
				return utils.ErrorDatabaseError
			}
		}

		return nil
	}))
}

func Unfollow(follower, followee int64) utils.ErrorCode {
	if follower == followee {
		return utils.ErrorUnanticipated
	}

	return utils.From(db.Transaction(func(tx *gorm.DB) error {
		relation := &RelationDO{
			Followee: followee,
			Follower: follower,
		}
		if err := tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "mutual"}}}).
			Delete(relation).Error; err != nil {
			return utils.ErrorUnanticipated
		}

		if relation.Mutual {
			other := &RelationDO{
				Followee: follower,
				Follower: followee,
			}
			if err := tx.Model(other).Select("mutual").
				Updates(RelationDO{Mutual: false}).Error; err != nil {
				return utils.ErrorDatabaseError
			}
		}

		return nil
	}))
}

func FollowerList(user int64) ([]int64, utils.ErrorCode) {
	var followers []int64
	if err := db.Model(&RelationDO{}).
		Select("follower").Where("followee = ?", user).
		Limit(300).Find(&followers).Error; err != nil {
		db.Error = nil
		return nil, utils.ErrorNoSuchUser
	}
	return followers, utils.ErrorOk
}

func FolloweeList(user int64) ([]int64, utils.ErrorCode) {
	var followees []int64
	if err := db.Model(&RelationDO{}).
		Select("followee").Where("follower = ?", user).
		Limit(300).Find(&followees).Error; err != nil {
		db.Error = nil
		return nil, utils.ErrorNoSuchUser
	}
	return followees, utils.ErrorOk
}

func FriendList(user int64) ([]int64, utils.ErrorCode) {
	var friends []int64
	if err := db.Model(&RelationDO{}).Select("followee").
		Where("follower = ?", user).Where("mutual = ?", true).
		Limit(300).Find(&friends).Error; err != nil {
		db.Error = nil
		return nil, utils.ErrorNoSuchUser
	}
	return friends, utils.ErrorOk
}

func FilterFollowees(user int64, ids []int64) (followees []int64, err utils.ErrorCode) {
	if e := db.Model(&RelationDO{}).Select("followee").
		Where("follower = ? AND followee IN ?", user, ids).
		Find(&followees).Error; e != nil {
		err = utils.ErrorDatabaseError
		followees = nil
	}
	return
}
