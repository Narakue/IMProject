package dao

import (
	"Test/common"
	"Test/model"
)

func Login(user model.User) bool {
	db := common.GetDB()
	result := db.First(&user)
	affected := result.RowsAffected
	err := result.Error
	if err != nil {
		println(err)
		return false
	}
	if affected > 0 {
		return true
	} else {
		return false
	}
}
